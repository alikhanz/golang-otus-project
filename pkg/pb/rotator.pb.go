// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: rotator.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SelectBannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlotId int32 `protobuf:"varint,1,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`
	SdgId  int32 `protobuf:"varint,2,opt,name=sdg_id,json=sdgId,proto3" json:"sdg_id,omitempty"`
}

func (x *SelectBannerRequest) Reset() {
	*x = SelectBannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rotator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectBannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectBannerRequest) ProtoMessage() {}

func (x *SelectBannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rotator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectBannerRequest.ProtoReflect.Descriptor instead.
func (*SelectBannerRequest) Descriptor() ([]byte, []int) {
	return file_rotator_proto_rawDescGZIP(), []int{0}
}

func (x *SelectBannerRequest) GetSlotId() int32 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

func (x *SelectBannerRequest) GetSdgId() int32 {
	if x != nil {
		return x.SdgId
	}
	return 0
}

type SelectBannerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BannerId int32 `protobuf:"varint,2,opt,name=banner_id,json=bannerId,proto3" json:"banner_id,omitempty"`
}

func (x *SelectBannerResponse) Reset() {
	*x = SelectBannerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rotator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectBannerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectBannerResponse) ProtoMessage() {}

func (x *SelectBannerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rotator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectBannerResponse.ProtoReflect.Descriptor instead.
func (*SelectBannerResponse) Descriptor() ([]byte, []int) {
	return file_rotator_proto_rawDescGZIP(), []int{1}
}

func (x *SelectBannerResponse) GetBannerId() int32 {
	if x != nil {
		return x.BannerId
	}
	return 0
}

type AddBannerToSlotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BannerId int32 `protobuf:"varint,1,opt,name=banner_id,json=bannerId,proto3" json:"banner_id,omitempty"`
	SlotId   int32 `protobuf:"varint,2,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`
}

func (x *AddBannerToSlotRequest) Reset() {
	*x = AddBannerToSlotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rotator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBannerToSlotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBannerToSlotRequest) ProtoMessage() {}

func (x *AddBannerToSlotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rotator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBannerToSlotRequest.ProtoReflect.Descriptor instead.
func (*AddBannerToSlotRequest) Descriptor() ([]byte, []int) {
	return file_rotator_proto_rawDescGZIP(), []int{2}
}

func (x *AddBannerToSlotRequest) GetBannerId() int32 {
	if x != nil {
		return x.BannerId
	}
	return 0
}

func (x *AddBannerToSlotRequest) GetSlotId() int32 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

type AddBannerToSlotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddBannerToSlotResponse) Reset() {
	*x = AddBannerToSlotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rotator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBannerToSlotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBannerToSlotResponse) ProtoMessage() {}

func (x *AddBannerToSlotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rotator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBannerToSlotResponse.ProtoReflect.Descriptor instead.
func (*AddBannerToSlotResponse) Descriptor() ([]byte, []int) {
	return file_rotator_proto_rawDescGZIP(), []int{3}
}

type RemoveBannerFromSlotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BannerId int32 `protobuf:"varint,1,opt,name=banner_id,json=bannerId,proto3" json:"banner_id,omitempty"`
	SlotId   int32 `protobuf:"varint,2,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`
}

func (x *RemoveBannerFromSlotRequest) Reset() {
	*x = RemoveBannerFromSlotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rotator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveBannerFromSlotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveBannerFromSlotRequest) ProtoMessage() {}

func (x *RemoveBannerFromSlotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rotator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveBannerFromSlotRequest.ProtoReflect.Descriptor instead.
func (*RemoveBannerFromSlotRequest) Descriptor() ([]byte, []int) {
	return file_rotator_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveBannerFromSlotRequest) GetBannerId() int32 {
	if x != nil {
		return x.BannerId
	}
	return 0
}

func (x *RemoveBannerFromSlotRequest) GetSlotId() int32 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

type RemoveBannerFromSlotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveBannerFromSlotResponse) Reset() {
	*x = RemoveBannerFromSlotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rotator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveBannerFromSlotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveBannerFromSlotResponse) ProtoMessage() {}

func (x *RemoveBannerFromSlotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rotator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveBannerFromSlotResponse.ProtoReflect.Descriptor instead.
func (*RemoveBannerFromSlotResponse) Descriptor() ([]byte, []int) {
	return file_rotator_proto_rawDescGZIP(), []int{5}
}

type AddBanner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *AddBanner) Reset() {
	*x = AddBanner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rotator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBanner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBanner) ProtoMessage() {}

func (x *AddBanner) ProtoReflect() protoreflect.Message {
	mi := &file_rotator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBanner.ProtoReflect.Descriptor instead.
func (*AddBanner) Descriptor() ([]byte, []int) {
	return file_rotator_proto_rawDescGZIP(), []int{6}
}

func (x *AddBanner) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type AddSlot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *AddSlot) Reset() {
	*x = AddSlot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rotator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSlot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSlot) ProtoMessage() {}

func (x *AddSlot) ProtoReflect() protoreflect.Message {
	mi := &file_rotator_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSlot.ProtoReflect.Descriptor instead.
func (*AddSlot) Descriptor() ([]byte, []int) {
	return file_rotator_proto_rawDescGZIP(), []int{7}
}

func (x *AddSlot) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Социально-демографическая группа
type AddSdg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *AddSdg) Reset() {
	*x = AddSdg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rotator_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSdg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSdg) ProtoMessage() {}

func (x *AddSdg) ProtoReflect() protoreflect.Message {
	mi := &file_rotator_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSdg.ProtoReflect.Descriptor instead.
func (*AddSdg) Descriptor() ([]byte, []int) {
	return file_rotator_proto_rawDescGZIP(), []int{8}
}

func (x *AddSdg) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_rotator_proto protoreflect.FileDescriptor

var file_rotator_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x45, 0x0a, 0x13, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x64, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x64, 0x67, 0x49, 0x64, 0x22,
	0x33, 0x0a, 0x14, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x54, 0x6f, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73,
	0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x6c,
	0x6f, 0x74, 0x49, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x54, 0x6f, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x53, 0x0a, 0x1b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x46,
	0x72, 0x6f, 0x6d, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73,
	0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x6c,
	0x6f, 0x74, 0x49, 0x64, 0x22, 0x1e, 0x0a, 0x1c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x2b, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x2a, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x53, 0x64, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x91, 0x02, 0x0a,
	0x07, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x4b, 0x0a, 0x0c, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x72, 0x6f, 0x74, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x54, 0x6f, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x1f, 0x2e, 0x72, 0x6f, 0x74, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x54, 0x6f, 0x53, 0x6c,
	0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72, 0x6f, 0x74, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x54, 0x6f, 0x53,
	0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x14, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x53,
	0x6c, 0x6f, 0x74, 0x12, 0x24, 0x2e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x6c,
	0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x72, 0x6f, 0x74, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x46, 0x72, 0x6f, 0x6d, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x6c, 0x69, 0x6b, 0x68, 0x61, 0x6e, 0x7a, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x6f,
	0x74, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rotator_proto_rawDescOnce sync.Once
	file_rotator_proto_rawDescData = file_rotator_proto_rawDesc
)

func file_rotator_proto_rawDescGZIP() []byte {
	file_rotator_proto_rawDescOnce.Do(func() {
		file_rotator_proto_rawDescData = protoimpl.X.CompressGZIP(file_rotator_proto_rawDescData)
	})
	return file_rotator_proto_rawDescData
}

var file_rotator_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rotator_proto_goTypes = []interface{}{
	(*SelectBannerRequest)(nil),          // 0: rotator.SelectBannerRequest
	(*SelectBannerResponse)(nil),         // 1: rotator.SelectBannerResponse
	(*AddBannerToSlotRequest)(nil),       // 2: rotator.AddBannerToSlotRequest
	(*AddBannerToSlotResponse)(nil),      // 3: rotator.AddBannerToSlotResponse
	(*RemoveBannerFromSlotRequest)(nil),  // 4: rotator.RemoveBannerFromSlotRequest
	(*RemoveBannerFromSlotResponse)(nil), // 5: rotator.RemoveBannerFromSlotResponse
	(*AddBanner)(nil),                    // 6: rotator.AddBanner
	(*AddSlot)(nil),                      // 7: rotator.AddSlot
	(*AddSdg)(nil),                       // 8: rotator.AddSdg
}
var file_rotator_proto_depIdxs = []int32{
	0, // 0: rotator.Rotator.SelectBanner:input_type -> rotator.SelectBannerRequest
	2, // 1: rotator.Rotator.AddBannerToSlot:input_type -> rotator.AddBannerToSlotRequest
	4, // 2: rotator.Rotator.RemoveBannerFromSlot:input_type -> rotator.RemoveBannerFromSlotRequest
	1, // 3: rotator.Rotator.SelectBanner:output_type -> rotator.SelectBannerResponse
	3, // 4: rotator.Rotator.AddBannerToSlot:output_type -> rotator.AddBannerToSlotResponse
	5, // 5: rotator.Rotator.RemoveBannerFromSlot:output_type -> rotator.RemoveBannerFromSlotResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rotator_proto_init() }
func file_rotator_proto_init() {
	if File_rotator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rotator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectBannerRequest); i {
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
		file_rotator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectBannerResponse); i {
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
		file_rotator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBannerToSlotRequest); i {
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
		file_rotator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBannerToSlotResponse); i {
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
		file_rotator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveBannerFromSlotRequest); i {
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
		file_rotator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveBannerFromSlotResponse); i {
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
		file_rotator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBanner); i {
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
		file_rotator_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSlot); i {
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
		file_rotator_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSdg); i {
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
			RawDescriptor: file_rotator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rotator_proto_goTypes,
		DependencyIndexes: file_rotator_proto_depIdxs,
		MessageInfos:      file_rotator_proto_msgTypes,
	}.Build()
	File_rotator_proto = out.File
	file_rotator_proto_rawDesc = nil
	file_rotator_proto_goTypes = nil
	file_rotator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RotatorClient is the client API for Rotator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RotatorClient interface {
	SelectBanner(ctx context.Context, in *SelectBannerRequest, opts ...grpc.CallOption) (*SelectBannerResponse, error)
	AddBannerToSlot(ctx context.Context, in *AddBannerToSlotRequest, opts ...grpc.CallOption) (*AddBannerToSlotResponse, error)
	RemoveBannerFromSlot(ctx context.Context, in *RemoveBannerFromSlotRequest, opts ...grpc.CallOption) (*RemoveBannerFromSlotResponse, error)
}

type rotatorClient struct {
	cc grpc.ClientConnInterface
}

func NewRotatorClient(cc grpc.ClientConnInterface) RotatorClient {
	return &rotatorClient{cc}
}

func (c *rotatorClient) SelectBanner(ctx context.Context, in *SelectBannerRequest, opts ...grpc.CallOption) (*SelectBannerResponse, error) {
	out := new(SelectBannerResponse)
	err := c.cc.Invoke(ctx, "/rotator.Rotator/SelectBanner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rotatorClient) AddBannerToSlot(ctx context.Context, in *AddBannerToSlotRequest, opts ...grpc.CallOption) (*AddBannerToSlotResponse, error) {
	out := new(AddBannerToSlotResponse)
	err := c.cc.Invoke(ctx, "/rotator.Rotator/AddBannerToSlot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rotatorClient) RemoveBannerFromSlot(ctx context.Context, in *RemoveBannerFromSlotRequest, opts ...grpc.CallOption) (*RemoveBannerFromSlotResponse, error) {
	out := new(RemoveBannerFromSlotResponse)
	err := c.cc.Invoke(ctx, "/rotator.Rotator/RemoveBannerFromSlot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RotatorServer is the server API for Rotator service.
type RotatorServer interface {
	SelectBanner(context.Context, *SelectBannerRequest) (*SelectBannerResponse, error)
	AddBannerToSlot(context.Context, *AddBannerToSlotRequest) (*AddBannerToSlotResponse, error)
	RemoveBannerFromSlot(context.Context, *RemoveBannerFromSlotRequest) (*RemoveBannerFromSlotResponse, error)
}

// UnimplementedRotatorServer can be embedded to have forward compatible implementations.
type UnimplementedRotatorServer struct {
}

func (*UnimplementedRotatorServer) SelectBanner(context.Context, *SelectBannerRequest) (*SelectBannerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectBanner not implemented")
}
func (*UnimplementedRotatorServer) AddBannerToSlot(context.Context, *AddBannerToSlotRequest) (*AddBannerToSlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBannerToSlot not implemented")
}
func (*UnimplementedRotatorServer) RemoveBannerFromSlot(context.Context, *RemoveBannerFromSlotRequest) (*RemoveBannerFromSlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBannerFromSlot not implemented")
}

func RegisterRotatorServer(s *grpc.Server, srv RotatorServer) {
	s.RegisterService(&_Rotator_serviceDesc, srv)
}

func _Rotator_SelectBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectBannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RotatorServer).SelectBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rotator.Rotator/SelectBanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RotatorServer).SelectBanner(ctx, req.(*SelectBannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rotator_AddBannerToSlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBannerToSlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RotatorServer).AddBannerToSlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rotator.Rotator/AddBannerToSlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RotatorServer).AddBannerToSlot(ctx, req.(*AddBannerToSlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rotator_RemoveBannerFromSlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveBannerFromSlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RotatorServer).RemoveBannerFromSlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rotator.Rotator/RemoveBannerFromSlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RotatorServer).RemoveBannerFromSlot(ctx, req.(*RemoveBannerFromSlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rotator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rotator.Rotator",
	HandlerType: (*RotatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SelectBanner",
			Handler:    _Rotator_SelectBanner_Handler,
		},
		{
			MethodName: "AddBannerToSlot",
			Handler:    _Rotator_AddBannerToSlot_Handler,
		},
		{
			MethodName: "RemoveBannerFromSlot",
			Handler:    _Rotator_RemoveBannerFromSlot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rotator.proto",
}
