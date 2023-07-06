// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: proto/good.proto

package proto

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

// 直播间商品列表请求参数
type GetGoodsByRoomReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId int64 `protobuf:"varint,1,opt,name=RoomId,proto3" json:"RoomId,omitempty"`
	UserId int64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *GetGoodsByRoomReq) Reset() {
	*x = GetGoodsByRoomReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_good_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsByRoomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsByRoomReq) ProtoMessage() {}

func (x *GetGoodsByRoomReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_good_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsByRoomReq.ProtoReflect.Descriptor instead.
func (*GetGoodsByRoomReq) Descriptor() ([]byte, []int) {
	return file_proto_good_proto_rawDescGZIP(), []int{0}
}

func (x *GetGoodsByRoomReq) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *GetGoodsByRoomReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 直播间商品列表返回参数
type GetGoodsByRoomResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentGoodsId int64        `protobuf:"varint,1,opt,name=CurrentGoodsId,proto3" json:"CurrentGoodsId,omitempty"` // 置顶商品id
	Data           []*GoodsInfo `protobuf:"bytes,2,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *GetGoodsByRoomResp) Reset() {
	*x = GetGoodsByRoomResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_good_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsByRoomResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsByRoomResp) ProtoMessage() {}

func (x *GetGoodsByRoomResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_good_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsByRoomResp.ProtoReflect.Descriptor instead.
func (*GetGoodsByRoomResp) Descriptor() ([]byte, []int) {
	return file_proto_good_proto_rawDescGZIP(), []int{1}
}

func (x *GetGoodsByRoomResp) GetCurrentGoodsId() int64 {
	if x != nil {
		return x.CurrentGoodsId
	}
	return 0
}

func (x *GetGoodsByRoomResp) GetData() []*GoodsInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

// 列表页的商品数据
type GoodsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsId     int64    `protobuf:"varint,1,opt,name=GoodsId,proto3" json:"GoodsId,omitempty"`
	CategoryId  int64    `protobuf:"varint,2,opt,name=CategoryId,proto3" json:"CategoryId,omitempty"`
	Status      int32    `protobuf:"varint,3,opt,name=Status,proto3" json:"Status,omitempty"`
	Title       string   `protobuf:"bytes,4,opt,name=Title,proto3" json:"Title,omitempty"`
	MarketPrice string   `protobuf:"bytes,5,opt,name=MarketPrice,proto3" json:"MarketPrice,omitempty"`
	Price       string   `protobuf:"bytes,6,opt,name=Price,proto3" json:"Price,omitempty"`
	Brief       string   `protobuf:"bytes,7,opt,name=Brief,proto3" json:"Brief,omitempty"`
	HeadImgs    []string `protobuf:"bytes,8,rep,name=HeadImgs,proto3" json:"HeadImgs,omitempty"`
}

func (x *GoodsInfo) Reset() {
	*x = GoodsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_good_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsInfo) ProtoMessage() {}

func (x *GoodsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_good_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsInfo.ProtoReflect.Descriptor instead.
func (*GoodsInfo) Descriptor() ([]byte, []int) {
	return file_proto_good_proto_rawDescGZIP(), []int{2}
}

func (x *GoodsInfo) GetGoodsId() int64 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *GoodsInfo) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *GoodsInfo) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GoodsInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GoodsInfo) GetMarketPrice() string {
	if x != nil {
		return x.MarketPrice
	}
	return ""
}

func (x *GoodsInfo) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *GoodsInfo) GetBrief() string {
	if x != nil {
		return x.Brief
	}
	return ""
}

func (x *GoodsInfo) GetHeadImgs() []string {
	if x != nil {
		return x.HeadImgs
	}
	return nil
}

// 商品详情请求参数
type GetGoodDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsId int64 `protobuf:"varint,1,opt,name=GoodsId,proto3" json:"GoodsId,omitempty"`
	UserId  int64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *GetGoodDetailReq) Reset() {
	*x = GetGoodDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_good_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodDetailReq) ProtoMessage() {}

func (x *GetGoodDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_good_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodDetailReq.ProtoReflect.Descriptor instead.
func (*GetGoodDetailReq) Descriptor() ([]byte, []int) {
	return file_proto_good_proto_rawDescGZIP(), []int{3}
}

func (x *GetGoodDetailReq) GetGoodsId() int64 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *GetGoodDetailReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 商品详情返回参数
type GetGoodDetailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsId     int64    `protobuf:"varint,1,opt,name=GoodsId,proto3" json:"GoodsId,omitempty"`
	CategoryId  int64    `protobuf:"varint,2,opt,name=CategoryId,proto3" json:"CategoryId,omitempty"`
	Status      int32    `protobuf:"varint,3,opt,name=Status,proto3" json:"Status,omitempty"`
	Title       string   `protobuf:"bytes,4,opt,name=Title,proto3" json:"Title,omitempty"`
	Code        string   `protobuf:"bytes,5,opt,name=Code,proto3" json:"Code,omitempty"`
	BrandName   string   `protobuf:"bytes,6,opt,name=BrandName,proto3" json:"BrandName,omitempty"`
	MarketPrice string   `protobuf:"bytes,7,opt,name=MarketPrice,proto3" json:"MarketPrice,omitempty"`
	Price       string   `protobuf:"bytes,8,opt,name=Price,proto3" json:"Price,omitempty"`
	Brief       string   `protobuf:"bytes,9,opt,name=Brief,proto3" json:"Brief,omitempty"`
	HeadImgs    []string `protobuf:"bytes,10,rep,name=HeadImgs,proto3" json:"HeadImgs,omitempty"`
	Videos      []string `protobuf:"bytes,11,rep,name=Videos,proto3" json:"Videos,omitempty"`
	Detail      []string `protobuf:"bytes,12,rep,name=Detail,proto3" json:"Detail,omitempty"`
}

func (x *GetGoodDetailResp) Reset() {
	*x = GetGoodDetailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_good_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodDetailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodDetailResp) ProtoMessage() {}

func (x *GetGoodDetailResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_good_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodDetailResp.ProtoReflect.Descriptor instead.
func (*GetGoodDetailResp) Descriptor() ([]byte, []int) {
	return file_proto_good_proto_rawDescGZIP(), []int{4}
}

func (x *GetGoodDetailResp) GetGoodsId() int64 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *GetGoodDetailResp) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *GetGoodDetailResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetGoodDetailResp) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetGoodDetailResp) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetGoodDetailResp) GetBrandName() string {
	if x != nil {
		return x.BrandName
	}
	return ""
}

func (x *GetGoodDetailResp) GetMarketPrice() string {
	if x != nil {
		return x.MarketPrice
	}
	return ""
}

func (x *GetGoodDetailResp) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *GetGoodDetailResp) GetBrief() string {
	if x != nil {
		return x.Brief
	}
	return ""
}

func (x *GetGoodDetailResp) GetHeadImgs() []string {
	if x != nil {
		return x.HeadImgs
	}
	return nil
}

func (x *GetGoodDetailResp) GetVideos() []string {
	if x != nil {
		return x.Videos
	}
	return nil
}

func (x *GetGoodDetailResp) GetDetail() []string {
	if x != nil {
		return x.Detail
	}
	return nil
}

var File_proto_good_proto protoreflect.FileDescriptor

var file_proto_good_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x79, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x62,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x79, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x26, 0x0a, 0x0e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x22, 0xdd, 0x01, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x18, 0x0a, 0x07, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x42, 0x72, 0x69, 0x65, 0x66, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x42, 0x72, 0x69, 0x65, 0x66, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x65, 0x61, 0x64, 0x49, 0x6d,
	0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x48, 0x65, 0x61, 0x64, 0x49, 0x6d,
	0x67, 0x73, 0x22, 0x44, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xc7, 0x02, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x47, 0x6f, 0x6f, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42,
	0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x42, 0x72, 0x69, 0x65, 0x66, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x42, 0x72, 0x69, 0x65, 0x66, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x65, 0x61, 0x64, 0x49, 0x6d,
	0x67, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x48, 0x65, 0x61, 0x64, 0x49, 0x6d,
	0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x32, 0x96, 0x01, 0x0a, 0x05, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x47, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x79, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x42,
	0x79, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x79, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_good_proto_rawDescOnce sync.Once
	file_proto_good_proto_rawDescData = file_proto_good_proto_rawDesc
)

func file_proto_good_proto_rawDescGZIP() []byte {
	file_proto_good_proto_rawDescOnce.Do(func() {
		file_proto_good_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_good_proto_rawDescData)
	})
	return file_proto_good_proto_rawDescData
}

var file_proto_good_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_good_proto_goTypes = []interface{}{
	(*GetGoodsByRoomReq)(nil),  // 0: proto.GetGoodsByRoomReq
	(*GetGoodsByRoomResp)(nil), // 1: proto.GetGoodsByRoomResp
	(*GoodsInfo)(nil),          // 2: proto.GoodsInfo
	(*GetGoodDetailReq)(nil),   // 3: proto.GetGoodDetailReq
	(*GetGoodDetailResp)(nil),  // 4: proto.GetGoodDetailResp
}
var file_proto_good_proto_depIdxs = []int32{
	2, // 0: proto.GetGoodsByRoomResp.Data:type_name -> proto.GoodsInfo
	0, // 1: proto.Goods.GetGoodsByRoom:input_type -> proto.GetGoodsByRoomReq
	3, // 2: proto.Goods.GetGoodDetail:input_type -> proto.GetGoodDetailReq
	1, // 3: proto.Goods.GetGoodsByRoom:output_type -> proto.GetGoodsByRoomResp
	4, // 4: proto.Goods.GetGoodDetail:output_type -> proto.GetGoodDetailResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_good_proto_init() }
func file_proto_good_proto_init() {
	if File_proto_good_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_good_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsByRoomReq); i {
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
		file_proto_good_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsByRoomResp); i {
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
		file_proto_good_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsInfo); i {
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
		file_proto_good_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodDetailReq); i {
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
		file_proto_good_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodDetailResp); i {
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
			RawDescriptor: file_proto_good_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_good_proto_goTypes,
		DependencyIndexes: file_proto_good_proto_depIdxs,
		MessageInfos:      file_proto_good_proto_msgTypes,
	}.Build()
	File_proto_good_proto = out.File
	file_proto_good_proto_rawDesc = nil
	file_proto_good_proto_goTypes = nil
	file_proto_good_proto_depIdxs = nil
}
