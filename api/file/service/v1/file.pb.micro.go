// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: file.proto

package v1

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for FileService service

func NewFileServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FileService service

type FileService interface {
	Query(ctx context.Context, in *QueryFile, opts ...client.CallOption) (*FileInfo, error)
	MarkChunkOK(ctx context.Context, in *MarkChunk, opts ...client.CallOption) (*FileInfo, error)
	StartSegmentUpload(ctx context.Context, in *SegmentUpload, opts ...client.CallOption) (*FileInfo, error)
}

type fileService struct {
	c    client.Client
	name string
}

func NewFileService(name string, c client.Client) FileService {
	return &fileService{
		c:    c,
		name: name,
	}
}

func (c *fileService) Query(ctx context.Context, in *QueryFile, opts ...client.CallOption) (*FileInfo, error) {
	req := c.c.NewRequest(c.name, "FileService.Query", in)
	out := new(FileInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) MarkChunkOK(ctx context.Context, in *MarkChunk, opts ...client.CallOption) (*FileInfo, error) {
	req := c.c.NewRequest(c.name, "FileService.MarkChunkOK", in)
	out := new(FileInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) StartSegmentUpload(ctx context.Context, in *SegmentUpload, opts ...client.CallOption) (*FileInfo, error) {
	req := c.c.NewRequest(c.name, "FileService.StartSegmentUpload", in)
	out := new(FileInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FileService service

type FileServiceHandler interface {
	Query(context.Context, *QueryFile, *FileInfo) error
	MarkChunkOK(context.Context, *MarkChunk, *FileInfo) error
	StartSegmentUpload(context.Context, *SegmentUpload, *FileInfo) error
}

func RegisterFileServiceHandler(s server.Server, hdlr FileServiceHandler, opts ...server.HandlerOption) error {
	type fileService interface {
		Query(ctx context.Context, in *QueryFile, out *FileInfo) error
		MarkChunkOK(ctx context.Context, in *MarkChunk, out *FileInfo) error
		StartSegmentUpload(ctx context.Context, in *SegmentUpload, out *FileInfo) error
	}
	type FileService struct {
		fileService
	}
	h := &fileServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FileService{h}, opts...))
}

type fileServiceHandler struct {
	FileServiceHandler
}

func (h *fileServiceHandler) Query(ctx context.Context, in *QueryFile, out *FileInfo) error {
	return h.FileServiceHandler.Query(ctx, in, out)
}

func (h *fileServiceHandler) MarkChunkOK(ctx context.Context, in *MarkChunk, out *FileInfo) error {
	return h.FileServiceHandler.MarkChunkOK(ctx, in, out)
}

func (h *fileServiceHandler) StartSegmentUpload(ctx context.Context, in *SegmentUpload, out *FileInfo) error {
	return h.FileServiceHandler.StartSegmentUpload(ctx, in, out)
}
