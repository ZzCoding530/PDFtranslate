package ocr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"io"
	"log"
	"net/http"
	"os"
	v1 "paper-translation/api/ocr/service/v1"
	"paper-translation/pkg/pdf"
	"sync"
	"time"

	ocr "github.com/alibabacloud-go/ocr-api-20210707/client"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type OCRStatus struct {
	Text     string
	Finished bool
}

func (t *OCRStatus) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, t)
}

func (t OCRStatus) MarshalBinary() (data []byte, err error) {
	return json.Marshal(t)
}

type OCRService struct {
	ocrRepo OCRRepository

	ocr         *ocr.Client
	oss         *oss.Client
	redisClient *redis.Client
}

func NewOCRService(ocrRepo OCRRepository, ocr *ocr.Client, oss *oss.Client, redisClient *redis.Client) *OCRService {
	return &OCRService{ocrRepo: ocrRepo, ocr: ocr, oss: oss, redisClient: redisClient}
}

func (t *OCRService) OCR(ctx context.Context, param *v1.OCRParam, resp *v1.OCRTaskID) error {

	ocx, err := t.ocrRepo.Get(param.Bucket, param.ObjectKey, param.FileType)
	if err == nil {
		resp.TaskId = uuid.NewString()
		t.redisClient.Set(ctx, resp.TaskId, OCRStatus{Text: ocx.OcredText, Finished: true}, time.Hour)
		return nil
	}

	if !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}

	resp.TaskId = uuid.NewString()
	t.redisClient.Set(ctx, resp.TaskId, OCRStatus{Text: "", Finished: false}, time.Hour)
	go func() {
		err = t.StartPipeline(context.TODO(), resp.TaskId, param.Bucket, param.ObjectKey)
		if err != nil {
			log.Printf("exec ocr pipeline failed err: %+v", err)
		}
	}()
	return nil
}

func (t *OCRService) GetStatus(ctx context.Context, req *v1.OCRTaskID, resp *v1.OCRText) error {

	var status OCRStatus
	err := t.redisClient.Get(ctx, req.TaskId).Scan(&status)
	if err != nil {
		return err
	}
	resp.Text = status.Text
	resp.Finished = status.Finished
	return nil
}

func (t *OCRService) OCRLocalImage(bucket, filePath string) (string, error) {

	bkt, err := t.oss.Bucket(bucket)
	if err != nil {
		return "", err
	}

	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	objectKey := fmt.Sprintf("images/%s.jpg", uuid.NewString())
	err = bkt.PutObject(objectKey, f)
	if err != nil {
		return "", err
	}

	fileURL, err := bkt.SignURL(objectKey, http.MethodGet, 120)
	if err != nil {
		return "", err
	}

	log.Printf("start ocr for image %s", fileURL)
	req := &ocr.RecognizeEnglishRequest{Url: &fileURL}
	resp, err := t.ocr.RecognizeEnglish(req)
	if err != nil {
		return "", err
	}

	var data map[string]any
	err = json.Unmarshal([]byte(*resp.Body.Data), &data)
	if err != nil {
		return "", err
	}

	return data["content"].(string), nil
}

func (t *OCRService) ConvertLocalImages(bucket, filePath string) ([]string, func(), error) {
	bkt, err := t.oss.Bucket(bucket)
	if err != nil {
		return nil, nil, err
	}

	log.Printf("start ocr for object: %s", filePath)
	object, err := bkt.GetObject(filePath)
	if err != nil {
		log.Printf("get object %s/%s err: %+v", bucket, filePath, err)
		return nil, nil, err
	}

	defer object.Close()

	localFilePath := fmt.Sprintf("%s/%s.pdf", os.TempDir(), uuid.NewString())
	file, err := os.OpenFile(localFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, err
	}
	_, _ = io.Copy(file, object)
	_ = file.Close()
	defer func() {
		_ = os.RemoveAll(localFilePath)
	}()

	return pdf.ConvertPdfToImages(localFilePath)
}

func (t *OCRService) StartPipeline(ctx context.Context, taskID, bucket, filePath string) error {
	images, clean, err := t.ConvertLocalImages(bucket, filePath)
	if err != nil {
		return err
	}
	defer clean()

	log.Printf("convert images is %+v", images)
	var wg sync.WaitGroup
	var texts = make([]string, len(images))
	for index, imagePath := range images {
		wg.Add(1)
		go func(index int, imagePath string) {
			defer wg.Done()
			text, err := t.OCRLocalImage(bucket, imagePath)
			if err != nil {
				log.Printf("ocr err: %+v", err)
				return
			}
			texts[index] = text
		}(index, imagePath)
	}
	wg.Wait()

	var buf bytes.Buffer
	for i := range texts {
		buf.WriteString(texts[i])
	}

	t.redisClient.Set(ctx, taskID, OCRStatus{Text: buf.String(), Finished: true}, time.Hour)
	if buf.String() != "" {
		_ = t.ocrRepo.Create(&OCR{
			ID:        taskID,
			Bucket:    bucket,
			ObjectKey: filePath,
			OcredText: buf.String(),
		})
	}
	return nil
}
