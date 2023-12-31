// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: paper.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Paper_Status int32

const (
	Paper_ocr         Paper_Status = 0
	Paper_translation Paper_Status = 1
	Paper_finished    Paper_Status = 2
	Paper_failed      Paper_Status = 3
)

// Enum value maps for Paper_Status.
var (
	Paper_Status_name = map[int32]string{
		0: "ocr",
		1: "translation",
		2: "finished",
		3: "failed",
	}
	Paper_Status_value = map[string]int32{
		"ocr":         0,
		"translation": 1,
		"finished":    2,
		"failed":      3,
	}
)

func (x Paper_Status) Enum() *Paper_Status {
	p := new(Paper_Status)
	*p = x
	return p
}

func (x Paper_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Paper_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_paper_proto_enumTypes[0].Descriptor()
}

func (Paper_Status) Type() protoreflect.EnumType {
	return &file_paper_proto_enumTypes[0]
}

func (x Paper_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Paper_Status.Descriptor instead.
func (Paper_Status) EnumDescriptor() ([]byte, []int) {
	return file_paper_proto_rawDescGZIP(), []int{1, 0}
}

type CreatePaper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaperFileHash  string `protobuf:"bytes,1,opt,name=paper_file_hash,json=paperFileHash,proto3" json:"paper_file_hash,omitempty"`
	EmailTo        string `protobuf:"bytes,2,opt,name=email_to,json=emailTo,proto3" json:"email_to,omitempty"`
	TargetLanguage string `protobuf:"bytes,3,opt,name=target_language,json=targetLanguage,proto3" json:"target_language,omitempty"`
}

func (x *CreatePaper) Reset() {
	*x = CreatePaper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePaper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaper) ProtoMessage() {}

func (x *CreatePaper) ProtoReflect() protoreflect.Message {
	mi := &file_paper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaper.ProtoReflect.Descriptor instead.
func (*CreatePaper) Descriptor() ([]byte, []int) {
	return file_paper_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePaper) GetPaperFileHash() string {
	if x != nil {
		return x.PaperFileHash
	}
	return ""
}

func (x *CreatePaper) GetEmailTo() string {
	if x != nil {
		return x.EmailTo
	}
	return ""
}

func (x *CreatePaper) GetTargetLanguage() string {
	if x != nil {
		return x.TargetLanguage
	}
	return ""
}

type Paper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FileHash       string       `protobuf:"bytes,2,opt,name=file_hash,json=fileHash,proto3" json:"file_hash,omitempty"`
	CreateAt       int64        `protobuf:"varint,3,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	Status         Paper_Status `protobuf:"varint,4,opt,name=status,proto3,enum=paper.service.v1.Paper_Status" json:"status,omitempty"`
	TargetLanguage string       `protobuf:"bytes,5,opt,name=target_language,json=targetLanguage,proto3" json:"target_language,omitempty"`
	ResultText     string       `protobuf:"bytes,6,opt,name=result_text,json=resultText,proto3" json:"result_text,omitempty"`
}

func (x *Paper) Reset() {
	*x = Paper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Paper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paper) ProtoMessage() {}

func (x *Paper) ProtoReflect() protoreflect.Message {
	mi := &file_paper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Paper.ProtoReflect.Descriptor instead.
func (*Paper) Descriptor() ([]byte, []int) {
	return file_paper_proto_rawDescGZIP(), []int{1}
}

func (x *Paper) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Paper) GetFileHash() string {
	if x != nil {
		return x.FileHash
	}
	return ""
}

func (x *Paper) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Paper) GetStatus() Paper_Status {
	if x != nil {
		return x.Status
	}
	return Paper_ocr
}

func (x *Paper) GetTargetLanguage() string {
	if x != nil {
		return x.TargetLanguage
	}
	return ""
}

func (x *Paper) GetResultText() string {
	if x != nil {
		return x.ResultText
	}
	return ""
}

type PaperID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PaperID) Reset() {
	*x = PaperID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paper_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaperID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaperID) ProtoMessage() {}

func (x *PaperID) ProtoReflect() protoreflect.Message {
	mi := &file_paper_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaperID.ProtoReflect.Descriptor instead.
func (*PaperID) Descriptor() ([]byte, []int) {
	return file_paper_proto_rawDescGZIP(), []int{2}
}

func (x *PaperID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeletePaper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePaper) Reset() {
	*x = DeletePaper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paper_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePaper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePaper) ProtoMessage() {}

func (x *DeletePaper) ProtoReflect() protoreflect.Message {
	mi := &file_paper_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePaper.ProtoReflect.Descriptor instead.
func (*DeletePaper) Descriptor() ([]byte, []int) {
	return file_paper_proto_rawDescGZIP(), []int{3}
}

type ReqFetchs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReqFetchs) Reset() {
	*x = ReqFetchs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paper_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqFetchs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqFetchs) ProtoMessage() {}

func (x *ReqFetchs) ProtoReflect() protoreflect.Message {
	mi := &file_paper_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqFetchs.ProtoReflect.Descriptor instead.
func (*ReqFetchs) Descriptor() ([]byte, []int) {
	return file_paper_proto_rawDescGZIP(), []int{4}
}

type RespFetchs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total  int32    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Papers []*Paper `protobuf:"bytes,2,rep,name=papers,proto3" json:"papers,omitempty"`
}

func (x *RespFetchs) Reset() {
	*x = RespFetchs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paper_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespFetchs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespFetchs) ProtoMessage() {}

func (x *RespFetchs) ProtoReflect() protoreflect.Message {
	mi := &file_paper_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespFetchs.ProtoReflect.Descriptor instead.
func (*RespFetchs) Descriptor() ([]byte, []int) {
	return file_paper_proto_rawDescGZIP(), []int{5}
}

func (x *RespFetchs) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *RespFetchs) GetPapers() []*Paper {
	if x != nil {
		return x.Papers
	}
	return nil
}

var File_paper_proto protoreflect.FileDescriptor

var file_paper_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70,
	0x61, 0x70, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22,
	0x79, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x70, 0x65, 0x72, 0x12, 0x26,
	0x0a, 0x0f, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x70, 0x65, 0x72, 0x46, 0x69,
	0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f,
	0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x54,
	0x6f, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x91, 0x02, 0x0a, 0x05, 0x50,
	0x61, 0x70, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x36,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e,
	0x2e, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x65, 0x78, 0x74,
	0x22, 0x3c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x07, 0x0a, 0x03, 0x6f, 0x63,
	0x72, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x03, 0x22, 0x19,
	0x0a, 0x07, 0x50, 0x61, 0x70, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x61, 0x70, 0x65, 0x72, 0x22, 0x0b, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x73, 0x22, 0x53, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x70, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x06, 0x70, 0x61, 0x70,
	0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x61, 0x70, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x70,
	0x65, 0x72, 0x52, 0x06, 0x70, 0x61, 0x70, 0x65, 0x72, 0x73, 0x32, 0x96, 0x02, 0x0a, 0x0c, 0x50,
	0x61, 0x70, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x70, 0x65, 0x72, 0x1a, 0x17, 0x2e, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x70, 0x65, 0x72, 0x12, 0x3b, 0x0a,
	0x05, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x70, 0x65, 0x72, 0x49,
	0x44, 0x1a, 0x17, 0x2e, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x70, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x70, 0x65, 0x72, 0x49, 0x44, 0x1a,
	0x1d, 0x2e, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x70, 0x65, 0x72, 0x12, 0x43,
	0x0a, 0x06, 0x46, 0x65, 0x74, 0x63, 0x68, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x61, 0x70, 0x65, 0x72,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x73, 0x1a, 0x1c, 0x2e, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x73, 0x42, 0x1b, 0x5a, 0x19, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x70,
	0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_paper_proto_rawDescOnce sync.Once
	file_paper_proto_rawDescData = file_paper_proto_rawDesc
)

func file_paper_proto_rawDescGZIP() []byte {
	file_paper_proto_rawDescOnce.Do(func() {
		file_paper_proto_rawDescData = protoimpl.X.CompressGZIP(file_paper_proto_rawDescData)
	})
	return file_paper_proto_rawDescData
}

var file_paper_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_paper_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_paper_proto_goTypes = []interface{}{
	(Paper_Status)(0),   // 0: paper.service.v1.Paper.Status
	(*CreatePaper)(nil), // 1: paper.service.v1.CreatePaper
	(*Paper)(nil),       // 2: paper.service.v1.Paper
	(*PaperID)(nil),     // 3: paper.service.v1.PaperID
	(*DeletePaper)(nil), // 4: paper.service.v1.DeletePaper
	(*ReqFetchs)(nil),   // 5: paper.service.v1.ReqFetchs
	(*RespFetchs)(nil),  // 6: paper.service.v1.RespFetchs
}
var file_paper_proto_depIdxs = []int32{
	0, // 0: paper.service.v1.Paper.status:type_name -> paper.service.v1.Paper.Status
	2, // 1: paper.service.v1.RespFetchs.papers:type_name -> paper.service.v1.Paper
	1, // 2: paper.service.v1.PaperService.Create:input_type -> paper.service.v1.CreatePaper
	3, // 3: paper.service.v1.PaperService.Fetch:input_type -> paper.service.v1.PaperID
	3, // 4: paper.service.v1.PaperService.Delete:input_type -> paper.service.v1.PaperID
	5, // 5: paper.service.v1.PaperService.Fetchs:input_type -> paper.service.v1.ReqFetchs
	2, // 6: paper.service.v1.PaperService.Create:output_type -> paper.service.v1.Paper
	2, // 7: paper.service.v1.PaperService.Fetch:output_type -> paper.service.v1.Paper
	4, // 8: paper.service.v1.PaperService.Delete:output_type -> paper.service.v1.DeletePaper
	6, // 9: paper.service.v1.PaperService.Fetchs:output_type -> paper.service.v1.RespFetchs
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_paper_proto_init() }
func file_paper_proto_init() {
	if File_paper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_paper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePaper); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_paper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Paper); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_paper_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaperID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_paper_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePaper); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_paper_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqFetchs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_paper_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespFetchs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_paper_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_paper_proto_goTypes,
		DependencyIndexes: file_paper_proto_depIdxs,
		EnumInfos:         file_paper_proto_enumTypes,
		MessageInfos:      file_paper_proto_msgTypes,
	}.Build()
	File_paper_proto = out.File
	file_paper_proto_rawDesc = nil
	file_paper_proto_goTypes = nil
	file_paper_proto_depIdxs = nil
}
