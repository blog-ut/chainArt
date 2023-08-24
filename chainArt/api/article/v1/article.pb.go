// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.5
// source: api/article/v1/article.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId    int64    `protobuf:"varint,1,opt,name=CategoryId,proto3" json:"CategoryId,omitempty"`
	Title         string   `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Intro         string   `protobuf:"bytes,3,opt,name=Intro,proto3" json:"Intro,omitempty"`
	Cover         string   `protobuf:"bytes,4,opt,name=Cover,proto3" json:"Cover,omitempty"`
	Content       string   `protobuf:"bytes,5,opt,name=Content,proto3" json:"Content,omitempty"`
	ContentMd     string   `protobuf:"bytes,6,opt,name=ContentMd,proto3" json:"ContentMd,omitempty"`
	KeyWords      []string `protobuf:"bytes,7,rep,name=KeyWords,proto3" json:"KeyWords,omitempty"`
	Sort          int64    `protobuf:"varint,8,opt,name=Sort,proto3" json:"Sort,omitempty"`
	IsElite       int64    `protobuf:"varint,9,opt,name=IsElite,proto3" json:"IsElite,omitempty"`
	ArticleStatus int64    `protobuf:"varint,10,opt,name=ArticleStatus,proto3" json:"ArticleStatus,omitempty"`
}

func (x *CreateArticleRequest) Reset() {
	*x = CreateArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_v1_article_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleRequest) ProtoMessage() {}

func (x *CreateArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_v1_article_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleRequest.ProtoReflect.Descriptor instead.
func (*CreateArticleRequest) Descriptor() ([]byte, []int) {
	return file_api_article_v1_article_proto_rawDescGZIP(), []int{0}
}

func (x *CreateArticleRequest) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *CreateArticleRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateArticleRequest) GetIntro() string {
	if x != nil {
		return x.Intro
	}
	return ""
}

func (x *CreateArticleRequest) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *CreateArticleRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateArticleRequest) GetContentMd() string {
	if x != nil {
		return x.ContentMd
	}
	return ""
}

func (x *CreateArticleRequest) GetKeyWords() []string {
	if x != nil {
		return x.KeyWords
	}
	return nil
}

func (x *CreateArticleRequest) GetSort() int64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *CreateArticleRequest) GetIsElite() int64 {
	if x != nil {
		return x.IsElite
	}
	return 0
}

func (x *CreateArticleRequest) GetArticleStatus() int64 {
	if x != nil {
		return x.ArticleStatus
	}
	return 0
}

type CreateArticleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *ArticleResp `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *CreateArticleReply) Reset() {
	*x = CreateArticleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_v1_article_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleReply) ProtoMessage() {}

func (x *CreateArticleReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_v1_article_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleReply.ProtoReflect.Descriptor instead.
func (*CreateArticleReply) Descriptor() ([]byte, []int) {
	return file_api_article_v1_article_proto_rawDescGZIP(), []int{1}
}

func (x *CreateArticleReply) GetInfo() *ArticleResp {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateArticleRequest) Reset() {
	*x = UpdateArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_v1_article_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArticleRequest) ProtoMessage() {}

func (x *UpdateArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_v1_article_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArticleRequest.ProtoReflect.Descriptor instead.
func (*UpdateArticleRequest) Descriptor() ([]byte, []int) {
	return file_api_article_v1_article_proto_rawDescGZIP(), []int{2}
}

type UpdateArticleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateArticleReply) Reset() {
	*x = UpdateArticleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_v1_article_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateArticleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArticleReply) ProtoMessage() {}

func (x *UpdateArticleReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_v1_article_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArticleReply.ProtoReflect.Descriptor instead.
func (*UpdateArticleReply) Descriptor() ([]byte, []int) {
	return file_api_article_v1_article_proto_rawDescGZIP(), []int{3}
}

type DeleteArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteArticleRequest) Reset() {
	*x = DeleteArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_v1_article_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArticleRequest) ProtoMessage() {}

func (x *DeleteArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_v1_article_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArticleRequest.ProtoReflect.Descriptor instead.
func (*DeleteArticleRequest) Descriptor() ([]byte, []int) {
	return file_api_article_v1_article_proto_rawDescGZIP(), []int{4}
}

type DeleteArticleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteArticleReply) Reset() {
	*x = DeleteArticleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_v1_article_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArticleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArticleReply) ProtoMessage() {}

func (x *DeleteArticleReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_v1_article_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArticleReply.ProtoReflect.Descriptor instead.
func (*DeleteArticleReply) Descriptor() ([]byte, []int) {
	return file_api_article_v1_article_proto_rawDescGZIP(), []int{5}
}

type GetArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *GetArticleRequest) Reset() {
	*x = GetArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_v1_article_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleRequest) ProtoMessage() {}

func (x *GetArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_v1_article_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleRequest.ProtoReflect.Descriptor instead.
func (*GetArticleRequest) Descriptor() ([]byte, []int) {
	return file_api_article_v1_article_proto_rawDescGZIP(), []int{6}
}

func (x *GetArticleRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetArticleRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GetArticleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *ArticleResp `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *GetArticleReply) Reset() {
	*x = GetArticleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_v1_article_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleReply) ProtoMessage() {}

func (x *GetArticleReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_v1_article_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleReply.ProtoReflect.Descriptor instead.
func (*GetArticleReply) Descriptor() ([]byte, []int) {
	return file_api_article_v1_article_proto_rawDescGZIP(), []int{7}
}

func (x *GetArticleReply) GetInfo() *ArticleResp {
	if x != nil {
		return x.Info
	}
	return nil
}

type ListArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize   int64    `protobuf:"varint,1,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	PageNum    int64    `protobuf:"varint,2,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	Title      string   `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	CategoryId int64    `protobuf:"varint,4,opt,name=CategoryId,proto3" json:"CategoryId,omitempty"`
	KeyWords   []string `protobuf:"bytes,5,rep,name=KeyWords,proto3" json:"KeyWords,omitempty"`
}

func (x *ListArticleRequest) Reset() {
	*x = ListArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_v1_article_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArticleRequest) ProtoMessage() {}

func (x *ListArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_v1_article_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArticleRequest.ProtoReflect.Descriptor instead.
func (*ListArticleRequest) Descriptor() ([]byte, []int) {
	return file_api_article_v1_article_proto_rawDescGZIP(), []int{8}
}

func (x *ListArticleRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListArticleRequest) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListArticleRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ListArticleRequest) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *ListArticleRequest) GetKeyWords() []string {
	if x != nil {
		return x.KeyWords
	}
	return nil
}

type ListArticleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64          `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Info  []*ArticleResp `protobuf:"bytes,2,rep,name=info,proto3" json:"info,omitempty"`
}

func (x *ListArticleReply) Reset() {
	*x = ListArticleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_v1_article_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListArticleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArticleReply) ProtoMessage() {}

func (x *ListArticleReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_v1_article_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArticleReply.ProtoReflect.Descriptor instead.
func (*ListArticleReply) Descriptor() ([]byte, []int) {
	return file_api_article_v1_article_proto_rawDescGZIP(), []int{9}
}

func (x *ListArticleReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListArticleReply) GetInfo() []*ArticleResp {
	if x != nil {
		return x.Info
	}
	return nil
}

type ArticleResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId        int64    `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	CategoryId    int64    `protobuf:"varint,3,opt,name=CategoryId,proto3" json:"CategoryId,omitempty"`
	Title         string   `protobuf:"bytes,4,opt,name=Title,proto3" json:"Title,omitempty"`
	Intro         string   `protobuf:"bytes,5,opt,name=Intro,proto3" json:"Intro,omitempty"`
	Cover         string   `protobuf:"bytes,6,opt,name=Cover,proto3" json:"Cover,omitempty"`
	Content       string   `protobuf:"bytes,7,opt,name=Content,proto3" json:"Content,omitempty"`
	ContentMd     string   `protobuf:"bytes,8,opt,name=ContentMd,proto3" json:"ContentMd,omitempty"`
	Code          string   `protobuf:"bytes,9,opt,name=Code,proto3" json:"Code,omitempty"`
	KeyWords      []string `protobuf:"bytes,10,rep,name=KeyWords,proto3" json:"KeyWords,omitempty"`
	IsElite       int64    `protobuf:"varint,11,opt,name=IsElite,proto3" json:"IsElite,omitempty"`
	Hits          int64    `protobuf:"varint,12,opt,name=Hits,proto3" json:"Hits,omitempty"`
	ArticleStatus int64    `protobuf:"varint,13,opt,name=ArticleStatus,proto3" json:"ArticleStatus,omitempty"`
	CreateAt      string   `protobuf:"bytes,14,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`
	UpdateAt      string   `protobuf:"bytes,15,opt,name=UpdateAt,proto3" json:"UpdateAt,omitempty"`
}

func (x *ArticleResp) Reset() {
	*x = ArticleResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_v1_article_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleResp) ProtoMessage() {}

func (x *ArticleResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_v1_article_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleResp.ProtoReflect.Descriptor instead.
func (*ArticleResp) Descriptor() ([]byte, []int) {
	return file_api_article_v1_article_proto_rawDescGZIP(), []int{10}
}

func (x *ArticleResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ArticleResp) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ArticleResp) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *ArticleResp) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ArticleResp) GetIntro() string {
	if x != nil {
		return x.Intro
	}
	return ""
}

func (x *ArticleResp) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *ArticleResp) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ArticleResp) GetContentMd() string {
	if x != nil {
		return x.ContentMd
	}
	return ""
}

func (x *ArticleResp) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ArticleResp) GetKeyWords() []string {
	if x != nil {
		return x.KeyWords
	}
	return nil
}

func (x *ArticleResp) GetIsElite() int64 {
	if x != nil {
		return x.IsElite
	}
	return 0
}

func (x *ArticleResp) GetHits() int64 {
	if x != nil {
		return x.Hits
	}
	return 0
}

func (x *ArticleResp) GetArticleStatus() int64 {
	if x != nil {
		return x.ArticleStatus
	}
	return 0
}

func (x *ArticleResp) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

func (x *ArticleResp) GetUpdateAt() string {
	if x != nil {
		return x.UpdateAt
	}
	return ""
}

var File_api_article_v1_article_proto protoreflect.FileDescriptor

var file_api_article_v1_article_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x02, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x49,
	0x6e, 0x74, 0x72, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6e, 0x74, 0x72,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4d, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4d, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x4b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x53,
	0x6f, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x49, 0x73, 0x45, 0x6c, 0x69, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x49, 0x73, 0x45, 0x6c, 0x69, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x45, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x37, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x42, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22,
	0x9c, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x4b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x59,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x8b, 0x03, 0x0a, 0x0b, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x74, 0x72, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x12, 0x14, 0x0a,
	0x05, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x6f,
	0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4d, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4d, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x4b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x49,
	0x73, 0x45, 0x6c, 0x69, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x49, 0x73,
	0x45, 0x6c, 0x69, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x69, 0x74, 0x73, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x48, 0x69, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x32, 0xab, 0x04, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x12, 0x7f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x12, 0x59, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x59, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x70, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x2f, 0x67, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x77, 0x0a, 0x0b,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x22, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x2e, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1a, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x41, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_article_v1_article_proto_rawDescOnce sync.Once
	file_api_article_v1_article_proto_rawDescData = file_api_article_v1_article_proto_rawDesc
)

func file_api_article_v1_article_proto_rawDescGZIP() []byte {
	file_api_article_v1_article_proto_rawDescOnce.Do(func() {
		file_api_article_v1_article_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_article_v1_article_proto_rawDescData)
	})
	return file_api_article_v1_article_proto_rawDescData
}

var file_api_article_v1_article_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_article_v1_article_proto_goTypes = []interface{}{
	(*CreateArticleRequest)(nil), // 0: api.article.v1.CreateArticleRequest
	(*CreateArticleReply)(nil),   // 1: api.article.v1.CreateArticleReply
	(*UpdateArticleRequest)(nil), // 2: api.article.v1.UpdateArticleRequest
	(*UpdateArticleReply)(nil),   // 3: api.article.v1.UpdateArticleReply
	(*DeleteArticleRequest)(nil), // 4: api.article.v1.DeleteArticleRequest
	(*DeleteArticleReply)(nil),   // 5: api.article.v1.DeleteArticleReply
	(*GetArticleRequest)(nil),    // 6: api.article.v1.GetArticleRequest
	(*GetArticleReply)(nil),      // 7: api.article.v1.GetArticleReply
	(*ListArticleRequest)(nil),   // 8: api.article.v1.ListArticleRequest
	(*ListArticleReply)(nil),     // 9: api.article.v1.ListArticleReply
	(*ArticleResp)(nil),          // 10: api.article.v1.ArticleResp
}
var file_api_article_v1_article_proto_depIdxs = []int32{
	10, // 0: api.article.v1.CreateArticleReply.info:type_name -> api.article.v1.ArticleResp
	10, // 1: api.article.v1.GetArticleReply.info:type_name -> api.article.v1.ArticleResp
	10, // 2: api.article.v1.ListArticleReply.info:type_name -> api.article.v1.ArticleResp
	0,  // 3: api.article.v1.Article.CreateArticle:input_type -> api.article.v1.CreateArticleRequest
	2,  // 4: api.article.v1.Article.UpdateArticle:input_type -> api.article.v1.UpdateArticleRequest
	4,  // 5: api.article.v1.Article.DeleteArticle:input_type -> api.article.v1.DeleteArticleRequest
	6,  // 6: api.article.v1.Article.GetArticle:input_type -> api.article.v1.GetArticleRequest
	8,  // 7: api.article.v1.Article.ListArticle:input_type -> api.article.v1.ListArticleRequest
	1,  // 8: api.article.v1.Article.CreateArticle:output_type -> api.article.v1.CreateArticleReply
	3,  // 9: api.article.v1.Article.UpdateArticle:output_type -> api.article.v1.UpdateArticleReply
	5,  // 10: api.article.v1.Article.DeleteArticle:output_type -> api.article.v1.DeleteArticleReply
	7,  // 11: api.article.v1.Article.GetArticle:output_type -> api.article.v1.GetArticleReply
	9,  // 12: api.article.v1.Article.ListArticle:output_type -> api.article.v1.ListArticleReply
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_article_v1_article_proto_init() }
func file_api_article_v1_article_proto_init() {
	if File_api_article_v1_article_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_article_v1_article_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleRequest); i {
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
		file_api_article_v1_article_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleReply); i {
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
		file_api_article_v1_article_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateArticleRequest); i {
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
		file_api_article_v1_article_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateArticleReply); i {
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
		file_api_article_v1_article_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteArticleRequest); i {
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
		file_api_article_v1_article_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteArticleReply); i {
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
		file_api_article_v1_article_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleRequest); i {
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
		file_api_article_v1_article_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleReply); i {
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
		file_api_article_v1_article_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListArticleRequest); i {
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
		file_api_article_v1_article_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListArticleReply); i {
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
		file_api_article_v1_article_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleResp); i {
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
			RawDescriptor: file_api_article_v1_article_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_article_v1_article_proto_goTypes,
		DependencyIndexes: file_api_article_v1_article_proto_depIdxs,
		MessageInfos:      file_api_article_v1_article_proto_msgTypes,
	}.Build()
	File_api_article_v1_article_proto = out.File
	file_api_article_v1_article_proto_rawDesc = nil
	file_api_article_v1_article_proto_goTypes = nil
	file_api_article_v1_article_proto_depIdxs = nil
}
