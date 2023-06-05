// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.0
// source: protos/chat/chat.proto

package chatpb

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

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primaryKey"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chat_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chat_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_protos_chat_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Room) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Room) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetRoomByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRoomByIdRequest) Reset() {
	*x = GetRoomByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chat_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomByIdRequest) ProtoMessage() {}

func (x *GetRoomByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chat_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomByIdRequest.ProtoReflect.Descriptor instead.
func (*GetRoomByIdRequest) Descriptor() ([]byte, []int) {
	return file_protos_chat_chat_proto_rawDescGZIP(), []int{1}
}

func (x *GetRoomByIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetAllRoomsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rooms []*Room `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *GetAllRoomsResponse) Reset() {
	*x = GetAllRoomsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chat_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRoomsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRoomsResponse) ProtoMessage() {}

func (x *GetAllRoomsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chat_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRoomsResponse.ProtoReflect.Descriptor instead.
func (*GetAllRoomsResponse) Descriptor() ([]byte, []int) {
	return file_protos_chat_chat_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllRoomsResponse) GetRooms() []*Room {
	if x != nil {
		return x.Rooms
	}
	return nil
}

type CreateRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateRoomRequest) Reset() {
	*x = CreateRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chat_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomRequest) ProtoMessage() {}

func (x *CreateRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chat_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomRequest.ProtoReflect.Descriptor instead.
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return file_protos_chat_chat_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRoomRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetMessagesByRoomIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId int64 `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
}

func (x *GetMessagesByRoomIdRequest) Reset() {
	*x = GetMessagesByRoomIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chat_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessagesByRoomIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesByRoomIdRequest) ProtoMessage() {}

func (x *GetMessagesByRoomIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chat_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesByRoomIdRequest.ProtoReflect.Descriptor instead.
func (*GetMessagesByRoomIdRequest) Descriptor() ([]byte, []int) {
	return file_protos_chat_chat_proto_rawDescGZIP(), []int{4}
}

func (x *GetMessagesByRoomIdRequest) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

type GetMessagesByRoomIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*ChatMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *GetMessagesByRoomIdResponse) Reset() {
	*x = GetMessagesByRoomIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chat_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessagesByRoomIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesByRoomIdResponse) ProtoMessage() {}

func (x *GetMessagesByRoomIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chat_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesByRoomIdResponse.ProtoReflect.Descriptor instead.
func (*GetMessagesByRoomIdResponse) Descriptor() ([]byte, []int) {
	return file_protos_chat_chat_proto_rawDescGZIP(), []int{5}
}

func (x *GetMessagesByRoomIdResponse) GetMessages() []*ChatMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

type JoinRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *JoinRoomRequest) Reset() {
	*x = JoinRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chat_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomRequest) ProtoMessage() {}

func (x *JoinRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chat_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoomRequest.ProtoReflect.Descriptor instead.
func (*JoinRoomRequest) Descriptor() ([]byte, []int) {
	return file_protos_chat_chat_proto_rawDescGZIP(), []int{6}
}

func (x *JoinRoomRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId   int64  `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	UserId   int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Content  string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Id       int64  `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty" gorm:"primaryKey"`
	SendedAt int64  `protobuf:"varint,5,opt,name=sendedAt,proto3" json:"sendedAt,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chat_chat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chat_chat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_protos_chat_chat_proto_rawDescGZIP(), []int{7}
}

func (x *ChatMessage) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *ChatMessage) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ChatMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ChatMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChatMessage) GetSendedAt() int64 {
	if x != nil {
		return x.SendedAt
	}
	return 0
}

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId  int64  `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	UserId  int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chat_chat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chat_chat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_protos_chat_chat_proto_rawDescGZIP(), []int{8}
}

func (x *SendMessageRequest) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *SendMessageRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SendMessageRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SendMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chat_chat_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chat_chat_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_protos_chat_chat_proto_rawDescGZIP(), []int{9}
}

func (x *SendMessageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chat_chat_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chat_chat_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_protos_chat_chat_proto_rawDescGZIP(), []int{10}
}

var File_protos_chat_chat_proto protoreflect.FileDescriptor

var file_protos_chat_chat_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62,
	0x22, 0x2a, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x39, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x6f, 0x6f, 0x6d,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x72, 0x6f, 0x6f,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70,
	0x62, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x22, 0x27, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x42, 0x79, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x1b,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42, 0x79, 0x52, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x21, 0x0a, 0x0f,
	0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x83, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x64, 0x41, 0x74, 0x22, 0x5e, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x2f, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32,
	0x9c, 0x03, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x39, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x0d,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x6f, 0x6f,
	0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2e, 0x52,
	0x6f, 0x6f, 0x6d, 0x12, 0x5e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x42, 0x79, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x22, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42,
	0x79, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x42, 0x79, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f,
	0x6d, 0x12, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x3a, 0x0a, 0x08, 0x4a, 0x6f,
	0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2e,
	0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x12, 0x46, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_chat_chat_proto_rawDescOnce sync.Once
	file_protos_chat_chat_proto_rawDescData = file_protos_chat_chat_proto_rawDesc
)

func file_protos_chat_chat_proto_rawDescGZIP() []byte {
	file_protos_chat_chat_proto_rawDescOnce.Do(func() {
		file_protos_chat_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_chat_chat_proto_rawDescData)
	})
	return file_protos_chat_chat_proto_rawDescData
}

var file_protos_chat_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_protos_chat_chat_proto_goTypes = []interface{}{
	(*Room)(nil),                        // 0: chatpb.Room
	(*GetRoomByIdRequest)(nil),          // 1: chatpb.GetRoomByIdRequest
	(*GetAllRoomsResponse)(nil),         // 2: chatpb.GetAllRoomsResponse
	(*CreateRoomRequest)(nil),           // 3: chatpb.CreateRoomRequest
	(*GetMessagesByRoomIdRequest)(nil),  // 4: chatpb.GetMessagesByRoomIdRequest
	(*GetMessagesByRoomIdResponse)(nil), // 5: chatpb.GetMessagesByRoomIdResponse
	(*JoinRoomRequest)(nil),             // 6: chatpb.JoinRoomRequest
	(*ChatMessage)(nil),                 // 7: chatpb.ChatMessage
	(*SendMessageRequest)(nil),          // 8: chatpb.SendMessageRequest
	(*SendMessageResponse)(nil),         // 9: chatpb.SendMessageResponse
	(*Empty)(nil),                       // 10: chatpb.Empty
}
var file_protos_chat_chat_proto_depIdxs = []int32{
	0,  // 0: chatpb.GetAllRoomsResponse.rooms:type_name -> chatpb.Room
	7,  // 1: chatpb.GetMessagesByRoomIdResponse.messages:type_name -> chatpb.ChatMessage
	10, // 2: chatpb.ChatService.GetAllRooms:input_type -> chatpb.Empty
	1,  // 3: chatpb.ChatService.GetRoomById:input_type -> chatpb.GetRoomByIdRequest
	4,  // 4: chatpb.ChatService.GetMessagesByRoomId:input_type -> chatpb.GetMessagesByRoomIdRequest
	3,  // 5: chatpb.ChatService.CreateRoom:input_type -> chatpb.CreateRoomRequest
	6,  // 6: chatpb.ChatService.JoinRoom:input_type -> chatpb.JoinRoomRequest
	8,  // 7: chatpb.ChatService.SendMessage:input_type -> chatpb.SendMessageRequest
	2,  // 8: chatpb.ChatService.GetAllRooms:output_type -> chatpb.GetAllRoomsResponse
	0,  // 9: chatpb.ChatService.GetRoomById:output_type -> chatpb.Room
	5,  // 10: chatpb.ChatService.GetMessagesByRoomId:output_type -> chatpb.GetMessagesByRoomIdResponse
	0,  // 11: chatpb.ChatService.CreateRoom:output_type -> chatpb.Room
	7,  // 12: chatpb.ChatService.JoinRoom:output_type -> chatpb.ChatMessage
	9,  // 13: chatpb.ChatService.SendMessage:output_type -> chatpb.SendMessageResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_protos_chat_chat_proto_init() }
func file_protos_chat_chat_proto_init() {
	if File_protos_chat_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_chat_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room); i {
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
		file_protos_chat_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomByIdRequest); i {
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
		file_protos_chat_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRoomsResponse); i {
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
		file_protos_chat_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomRequest); i {
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
		file_protos_chat_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessagesByRoomIdRequest); i {
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
		file_protos_chat_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessagesByRoomIdResponse); i {
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
		file_protos_chat_chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRoomRequest); i {
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
		file_protos_chat_chat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessage); i {
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
		file_protos_chat_chat_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
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
		file_protos_chat_chat_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageResponse); i {
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
		file_protos_chat_chat_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_protos_chat_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_chat_chat_proto_goTypes,
		DependencyIndexes: file_protos_chat_chat_proto_depIdxs,
		MessageInfos:      file_protos_chat_chat_proto_msgTypes,
	}.Build()
	File_protos_chat_chat_proto = out.File
	file_protos_chat_chat_proto_rawDesc = nil
	file_protos_chat_chat_proto_goTypes = nil
	file_protos_chat_chat_proto_depIdxs = nil
}
