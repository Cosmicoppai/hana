// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: video_message.proto

package show

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VideoId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *VideoId) Reset() {
	*x = VideoId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoId) ProtoMessage() {}

func (x *VideoId) ProtoReflect() protoreflect.Message {
	mi := &file_video_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoId.ProtoReflect.Descriptor instead.
func (*VideoId) Descriptor() ([]byte, []int) {
	return file_video_message_proto_rawDescGZIP(), []int{0}
}

func (x *VideoId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Video []byte `protobuf:"bytes,1,opt,name=video,proto3" json:"video,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_video_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_video_message_proto_rawDescGZIP(), []int{1}
}

func (x *Video) GetVideo() []byte {
	if x != nil {
		return x.Video
	}
	return nil
}

type UploadVideo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*UploadVideo_Video
	//	*UploadVideo_MetaData
	Data isUploadVideo_Data `protobuf_oneof:"data"`
}

func (x *UploadVideo) Reset() {
	*x = UploadVideo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadVideo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadVideo) ProtoMessage() {}

func (x *UploadVideo) ProtoReflect() protoreflect.Message {
	mi := &file_video_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadVideo.ProtoReflect.Descriptor instead.
func (*UploadVideo) Descriptor() ([]byte, []int) {
	return file_video_message_proto_rawDescGZIP(), []int{2}
}

func (m *UploadVideo) GetData() isUploadVideo_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UploadVideo) GetVideo() *Video {
	if x, ok := x.GetData().(*UploadVideo_Video); ok {
		return x.Video
	}
	return nil
}

func (x *UploadVideo) GetMetaData() *MetaData {
	if x, ok := x.GetData().(*UploadVideo_MetaData); ok {
		return x.MetaData
	}
	return nil
}

type isUploadVideo_Data interface {
	isUploadVideo_Data()
}

type UploadVideo_Video struct {
	Video *Video `protobuf:"bytes,1,opt,name=video,proto3,oneof"`
}

type UploadVideo_MetaData struct {
	MetaData *MetaData `protobuf:"bytes,2,opt,name=meta_data,json=metaData,proto3,oneof"`
}

func (*UploadVideo_Video) isUploadVideo_Data() {}

func (*UploadVideo_MetaData) isUploadVideo_Data() {}

type Videos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Videos []*Video `protobuf:"bytes,1,rep,name=videos,proto3" json:"videos,omitempty"`
}

func (x *Videos) Reset() {
	*x = Videos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Videos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Videos) ProtoMessage() {}

func (x *Videos) ProtoReflect() protoreflect.Message {
	mi := &file_video_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Videos.ProtoReflect.Descriptor instead.
func (*Videos) Descriptor() ([]byte, []int) {
	return file_video_message_proto_rawDescGZIP(), []int{3}
}

func (x *Videos) GetVideos() []*Video {
	if x != nil {
		return x.Videos
	}
	return nil
}

type MetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Poster    []byte                 `protobuf:"bytes,1,opt,name=poster,proto3" json:"poster,omitempty"`
	Sub       []byte                 `protobuf:"bytes,2,opt,name=sub,proto3" json:"sub,omitempty"`
	Name      string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *MetaData) Reset() {
	*x = MetaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaData) ProtoMessage() {}

func (x *MetaData) ProtoReflect() protoreflect.Message {
	mi := &file_video_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaData.ProtoReflect.Descriptor instead.
func (*MetaData) Descriptor() ([]byte, []int) {
	return file_video_message_proto_rawDescGZIP(), []int{4}
}

func (x *MetaData) GetPoster() []byte {
	if x != nil {
		return x.Poster
	}
	return nil
}

func (x *MetaData) GetSub() []byte {
	if x != nil {
		return x.Sub
	}
	return nil
}

func (x *MetaData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetaData) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type VideoMetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetaData *MetaData `protobuf:"bytes,1,opt,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty"`
	VideoId  *VideoId  `protobuf:"bytes,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
}

func (x *VideoMetaData) Reset() {
	*x = VideoMetaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoMetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoMetaData) ProtoMessage() {}

func (x *VideoMetaData) ProtoReflect() protoreflect.Message {
	mi := &file_video_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoMetaData.ProtoReflect.Descriptor instead.
func (*VideoMetaData) Descriptor() ([]byte, []int) {
	return file_video_message_proto_rawDescGZIP(), []int{5}
}

func (x *VideoMetaData) GetMetaData() *MetaData {
	if x != nil {
		return x.MetaData
	}
	return nil
}

func (x *VideoMetaData) GetVideoId() *VideoId {
	if x != nil {
		return x.VideoId
	}
	return nil
}

type VideosMetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Videos        []*VideoMetaData `protobuf:"bytes,1,rep,name=videos,proto3" json:"videos,omitempty"`
	TotalResponse uint32           `protobuf:"varint,3,opt,name=total_response,json=totalResponse,proto3" json:"total_response,omitempty"`
}

func (x *VideosMetaData) Reset() {
	*x = VideosMetaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideosMetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideosMetaData) ProtoMessage() {}

func (x *VideosMetaData) ProtoReflect() protoreflect.Message {
	mi := &file_video_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideosMetaData.ProtoReflect.Descriptor instead.
func (*VideosMetaData) Descriptor() ([]byte, []int) {
	return file_video_message_proto_rawDescGZIP(), []int{6}
}

func (x *VideosMetaData) GetVideos() []*VideoMetaData {
	if x != nil {
		return x.Videos
	}
	return nil
}

func (x *VideosMetaData) GetTotalResponse() uint32 {
	if x != nil {
		return x.TotalResponse
	}
	return 0
}

type VideosMetaDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *VideosMetaDataRequest) Reset() {
	*x = VideosMetaDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideosMetaDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideosMetaDataRequest) ProtoMessage() {}

func (x *VideosMetaDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideosMetaDataRequest.ProtoReflect.Descriptor instead.
func (*VideosMetaDataRequest) Descriptor() ([]byte, []int) {
	return file_video_message_proto_rawDescGZIP(), []int{7}
}

func (x *VideosMetaDataRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *VideosMetaDataRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_video_message_proto protoreflect.FileDescriptor

var file_video_message_proto_rawDesc = []byte{
	0x0a, 0x13, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19, 0x0a, 0x07, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x1d, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x22, 0x5f, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x12, 0x1e, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x06, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x48, 0x00, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x12, 0x28, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x48,
	0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x28, 0x0a, 0x06, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x1e, 0x0a,
	0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x22, 0x83, 0x01,
	0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f,
	0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x03, 0x73, 0x75, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x5c, 0x0a, 0x0d, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4d, 0x65, 0x74, 0x61,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x08,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49,
	0x64, 0x22, 0x5f, 0x0a, 0x0e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x45, 0x0a, 0x15, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x4d, 0x65, 0x74, 0x61,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x32, 0xc9, 0x01, 0x0a, 0x0c, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x09, 0x67, 0x65,
	0x74, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x08, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49,
	0x64, 0x1a, 0x06, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2b, 0x0a,
	0x12, 0x67, 0x65, 0x74, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x08, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x1a, 0x09, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x13, 0x67, 0x65,
	0x74, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x16, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x09,
	0x61, 0x64, 0x64, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x0c, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x1a, 0x08, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49,
	0x64, 0x22, 0x00, 0x28, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x73, 0x68, 0x6f, 0x77, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_video_message_proto_rawDescOnce sync.Once
	file_video_message_proto_rawDescData = file_video_message_proto_rawDesc
)

func file_video_message_proto_rawDescGZIP() []byte {
	file_video_message_proto_rawDescOnce.Do(func() {
		file_video_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_message_proto_rawDescData)
	})
	return file_video_message_proto_rawDescData
}

var file_video_message_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_video_message_proto_goTypes = []interface{}{
	(*VideoId)(nil),               // 0: VideoId
	(*Video)(nil),                 // 1: Video
	(*UploadVideo)(nil),           // 2: UploadVideo
	(*Videos)(nil),                // 3: Videos
	(*MetaData)(nil),              // 4: MetaData
	(*VideoMetaData)(nil),         // 5: VideoMetaData
	(*VideosMetaData)(nil),        // 6: VideosMetaData
	(*VideosMetaDataRequest)(nil), // 7: VideosMetaDataRequest
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_video_message_proto_depIdxs = []int32{
	1,  // 0: UploadVideo.video:type_name -> Video
	4,  // 1: UploadVideo.meta_data:type_name -> MetaData
	1,  // 2: Videos.videos:type_name -> Video
	8,  // 3: MetaData.created_at:type_name -> google.protobuf.Timestamp
	4,  // 4: VideoMetaData.meta_data:type_name -> MetaData
	0,  // 5: VideoMetaData.video_id:type_name -> VideoId
	5,  // 6: VideosMetaData.videos:type_name -> VideoMetaData
	0,  // 7: VideoService.get_video:input_type -> VideoId
	0,  // 8: VideoService.get_video_metadata:input_type -> VideoId
	7,  // 9: VideoService.get_videos_metadata:input_type -> VideosMetaDataRequest
	2,  // 10: VideoService.add_video:input_type -> UploadVideo
	1,  // 11: VideoService.get_video:output_type -> Video
	4,  // 12: VideoService.get_video_metadata:output_type -> MetaData
	6,  // 13: VideoService.get_videos_metadata:output_type -> VideosMetaData
	0,  // 14: VideoService.add_video:output_type -> VideoId
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_video_message_proto_init() }
func file_video_message_proto_init() {
	if File_video_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_video_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoId); i {
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
		file_video_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_video_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadVideo); i {
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
		file_video_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Videos); i {
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
		file_video_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaData); i {
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
		file_video_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoMetaData); i {
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
		file_video_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideosMetaData); i {
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
		file_video_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideosMetaDataRequest); i {
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
	file_video_message_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*UploadVideo_Video)(nil),
		(*UploadVideo_MetaData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_video_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_message_proto_goTypes,
		DependencyIndexes: file_video_message_proto_depIdxs,
		MessageInfos:      file_video_message_proto_msgTypes,
	}.Build()
	File_video_message_proto = out.File
	file_video_message_proto_rawDesc = nil
	file_video_message_proto_goTypes = nil
	file_video_message_proto_depIdxs = nil
}
