package handlers

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"mime/multipart"
	"net/http"
	fs "paper-translation/api/file/service/v1"
	"paper-translation/pkg/errutil"
	"sort"
	"strconv"
	"strings"
)

const (
	Bucket = "forwork"
)

type ReqStartUpload struct {
	Hash        string `json:"hash"`
	FileName    string `json:"fileName"`
	ChunkNums   int64  `json:"chunkNums"`
	SegmentSize int64  `json:"segmentSize"`
}

type ReqUploadChunk struct {
	Chunk      *multipart.FileHeader `form:"chunk"`
	Hash       string                `form:"hash"`
	ChunkIndex int64                 `form:"chunkIndex"`
}

type FileHandler struct {
	fileService fs.FileService
	oss         *oss.Client
}

func NewFileHandler(fileService fs.FileService, oss *oss.Client) *FileHandler {
	return &FileHandler{fileService: fileService, oss: oss}
}

func (f *FileHandler) StartUploadFile(ctx *gin.Context) {
	var req ReqStartUpload
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		errutil.ResponseError(ctx, errutil.RequestParamError, err)
		return
	}
	_, err = f.fileService.StartSegmentUpload(ctx, &fs.SegmentUpload{
		Hash:        req.Hash,
		Filename:    req.FileName,
		ChunkNums:   req.ChunkNums,
		SegmentSize: req.SegmentSize,
		Bucket:      Bucket,
		FilePath:    fmt.Sprintf("files/%s-%s", uuid.NewString(), req.FileName),
	})
	if err != nil {
		errutil.ResponseError(ctx, errutil.UnknownError, err)
		return
	}
}

func (f *FileHandler) UploadChunk(ctx *gin.Context) {

	var req ReqUploadChunk
	err := ctx.Bind(&req)
	if err != nil {
		errutil.ResponseError(ctx, errutil.RequestParamError, err)
		return
	}

	file, err := req.Chunk.Open()
	if err != nil {
		errutil.ResponseError(ctx, errutil.RequestParamError, err)
		return
	}

	bkt, err := f.oss.Bucket(Bucket)
	if err != nil {
		errutil.ResponseError(ctx, errutil.UnknownError, err)
		return
	}

	key := fmt.Sprintf("chunks/%s/%d", req.Hash, req.ChunkIndex)
	err = bkt.PutObject(key, file)
	if err != nil {
		errutil.ResponseError(ctx, errutil.UnknownError, err)
		return
	}

	fileInfo, err := f.fileService.MarkChunkOK(ctx, &fs.MarkChunk{
		Hash:       req.Hash,
		ChunkIndex: req.ChunkIndex,
	})
	if err != nil {
		errutil.ResponseError(ctx, errutil.UnknownError, err)
		return
	}

	if fileInfo.Status == fs.FileStatus_Uploaded {
		result, err := bkt.ListObjectsV2(oss.Prefix(fmt.Sprintf("chunks/%s/", req.Hash)))
		if err != nil {
			errutil.ResponseError(ctx, errutil.UnknownError, err)
			return
		}

		var keys = make([]string, 0)
		for _, object := range result.Objects {
			keys = append(keys, object.Key)
		}

		sort.SliceStable(keys, func(i, j int) bool {
			ix, _ := strconv.ParseInt(strings.Split(keys[i], "/")[2], 10, 64)
			jx, _ := strconv.ParseInt(strings.Split(keys[j], "/")[2], 10, 64)
			return ix < jx
		})

		var appendPosition = int64(0)
		for _, objectKey := range keys {
			log.Printf("append file: %s", objectKey)
			obj, err := bkt.GetObject(objectKey)
			if err != nil {
				errutil.ResponseError(ctx, errutil.UnknownError, err)
				return
			}
			nextPosition, err := bkt.AppendObject(*fileInfo.FilePath, obj, appendPosition)
			if err != nil {
				errutil.ResponseError(ctx, errutil.UnknownError, err)
				return
			}
			appendPosition = nextPosition
			obj.Close()
		}

		go func() {
			//文件备份
			_, _ = bkt.CopyObject(*fileInfo.FilePath, fmt.Sprintf("buckup/%s", *fileInfo.FilePath))
		}()
	}
}

func (f *FileHandler) QueryFile(ctx *gin.Context) {
	query, err := f.fileService.Query(ctx, &fs.QueryFile{Hash: ctx.Param("hash")})
	if err != nil {
		errutil.ResponseError(ctx, errutil.UnknownError, err)
		return
	}
	ctx.JSON(200, query)
}

func (f *FileHandler) GetFileURL(ctx *gin.Context) {
	query, err := f.fileService.Query(ctx, &fs.QueryFile{Hash: ctx.Param("hash")})
	if err != nil {
		errutil.ResponseError(ctx, errutil.UnknownError, err)
		return
	}

	if query.Status != fs.FileStatus_Uploaded {
		errutil.ResponseError(ctx, errutil.FileNotExistError)
		return
	}

	bkt, err := f.oss.Bucket(Bucket)
	if err != nil {
		errutil.ResponseError(ctx, errutil.UnknownError, err)
		return
	}

	url, err := bkt.SignURL(*query.FilePath, http.MethodGet, 10000)
	if err != nil {
		return
	}
	ctx.JSON(200, bson.M{
		"url": url,
	})
}
