// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: statistics.proto

// REMOVED omitempty FLAG FROM SPECIFIC_TASK_RESPONSE

package protos

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

type StatisticsType int32

const (
	StatisticsType_View StatisticsType = 0
	StatisticsType_Like StatisticsType = 1
)

// Enum value maps for StatisticsType.
var (
	StatisticsType_name = map[int32]string{
		0: "View",
		1: "Like",
	}
	StatisticsType_value = map[string]int32{
		"View": 0,
		"Like": 1,
	}
)

func (x StatisticsType) Enum() *StatisticsType {
	p := new(StatisticsType)
	*p = x
	return p
}

func (x StatisticsType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatisticsType) Descriptor() protoreflect.EnumDescriptor {
	return file_statistics_proto_enumTypes[0].Descriptor()
}

func (StatisticsType) Type() protoreflect.EnumType {
	return &file_statistics_proto_enumTypes[0]
}

func (x StatisticsType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatisticsType.Descriptor instead.
func (StatisticsType) EnumDescriptor() ([]byte, []int) {
	return file_statistics_proto_rawDescGZIP(), []int{0}
}

type SpecificTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskID int64 `protobuf:"varint,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
}

func (x *SpecificTaskRequest) Reset() {
	*x = SpecificTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecificTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecificTaskRequest) ProtoMessage() {}

func (x *SpecificTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecificTaskRequest.ProtoReflect.Descriptor instead.
func (*SpecificTaskRequest) Descriptor() ([]byte, []int) {
	return file_statistics_proto_rawDescGZIP(), []int{0}
}

func (x *SpecificTaskRequest) GetTaskID() int64 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

type SpecificTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskID     int64 `protobuf:"varint,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
	ViewsCount int64 `protobuf:"varint,2,opt,name=viewsCount,proto3" json:"viewsCount"`
	LikesCount int64 `protobuf:"varint,3,opt,name=likesCount,proto3" json:"likesCount"`
}

func (x *SpecificTaskResponse) Reset() {
	*x = SpecificTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecificTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecificTaskResponse) ProtoMessage() {}

func (x *SpecificTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecificTaskResponse.ProtoReflect.Descriptor instead.
func (*SpecificTaskResponse) Descriptor() ([]byte, []int) {
	return file_statistics_proto_rawDescGZIP(), []int{1}
}

func (x *SpecificTaskResponse) GetTaskID() int64 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

func (x *SpecificTaskResponse) GetViewsCount() int64 {
	if x != nil {
		return x.ViewsCount
	}
	return 0
}

func (x *SpecificTaskResponse) GetLikesCount() int64 {
	if x != nil {
		return x.LikesCount
	}
	return 0
}

type TopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type StatisticsType `protobuf:"varint,1,opt,name=type,proto3,enum=StatisticsType" json:"type,omitempty"`
}

func (x *TopRequest) Reset() {
	*x = TopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopRequest) ProtoMessage() {}

func (x *TopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopRequest.ProtoReflect.Descriptor instead.
func (*TopRequest) Descriptor() ([]byte, []int) {
	return file_statistics_proto_rawDescGZIP(), []int{2}
}

func (x *TopRequest) GetType() StatisticsType {
	if x != nil {
		return x.Type
	}
	return StatisticsType_View
}

type TaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskID        int64  `protobuf:"varint,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
	OwnerUsername string `protobuf:"bytes,2,opt,name=ownerUsername,proto3" json:"ownerUsername,omitempty"`
	Count         int64  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *TaskResponse) Reset() {
	*x = TaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResponse) ProtoMessage() {}

func (x *TaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResponse.ProtoReflect.Descriptor instead.
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return file_statistics_proto_rawDescGZIP(), []int{3}
}

func (x *TaskResponse) GetTaskID() int64 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

func (x *TaskResponse) GetOwnerUsername() string {
	if x != nil {
		return x.OwnerUsername
	}
	return ""
}

func (x *TaskResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerUsername string `protobuf:"bytes,1,opt,name=ownerUsername,proto3" json:"ownerUsername,omitempty"`
	Count         int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_statistics_proto_rawDescGZIP(), []int{4}
}

func (x *UserResponse) GetOwnerUsername() string {
	if x != nil {
		return x.OwnerUsername
	}
	return ""
}

func (x *UserResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type TopTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*TaskResponse `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *TopTasksResponse) Reset() {
	*x = TopTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopTasksResponse) ProtoMessage() {}

func (x *TopTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopTasksResponse.ProtoReflect.Descriptor instead.
func (*TopTasksResponse) Descriptor() ([]byte, []int) {
	return file_statistics_proto_rawDescGZIP(), []int{5}
}

func (x *TopTasksResponse) GetTasks() []*TaskResponse {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type TopUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserResponse `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *TopUsersResponse) Reset() {
	*x = TopUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopUsersResponse) ProtoMessage() {}

func (x *TopUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopUsersResponse.ProtoReflect.Descriptor instead.
func (*TopUsersResponse) Descriptor() ([]byte, []int) {
	return file_statistics_proto_rawDescGZIP(), []int{6}
}

func (x *TopUsersResponse) GetUsers() []*UserResponse {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_statistics_proto protoreflect.FileDescriptor

var file_statistics_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x13, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x44, 0x22, 0x6e, 0x0a, 0x14, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x31, 0x0a, 0x0a, 0x54, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x62, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x37, 0x0a, 0x10, 0x54, 0x6f, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x37, 0x0a,
	0x10, 0x54, 0x6f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x23, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2a, 0x24, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x10, 0x01, 0x32, 0xbb, 0x01, 0x0a,
	0x11, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x48, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12,
	0x14, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x0b, 0x2e, 0x54, 0x6f,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x54, 0x6f, 0x70, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x0b, 0x2e, 0x54, 0x6f, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x54, 0x6f, 0x70, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_statistics_proto_rawDescOnce sync.Once
	file_statistics_proto_rawDescData = file_statistics_proto_rawDesc
)

func file_statistics_proto_rawDescGZIP() []byte {
	file_statistics_proto_rawDescOnce.Do(func() {
		file_statistics_proto_rawDescData = protoimpl.X.CompressGZIP(file_statistics_proto_rawDescData)
	})
	return file_statistics_proto_rawDescData
}

var file_statistics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_statistics_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_statistics_proto_goTypes = []interface{}{
	(StatisticsType)(0),          // 0: StatisticsType
	(*SpecificTaskRequest)(nil),  // 1: SpecificTaskRequest
	(*SpecificTaskResponse)(nil), // 2: SpecificTaskResponse
	(*TopRequest)(nil),           // 3: TopRequest
	(*TaskResponse)(nil),         // 4: TaskResponse
	(*UserResponse)(nil),         // 5: UserResponse
	(*TopTasksResponse)(nil),     // 6: TopTasksResponse
	(*TopUsersResponse)(nil),     // 7: TopUsersResponse
}
var file_statistics_proto_depIdxs = []int32{
	0, // 0: TopRequest.type:type_name -> StatisticsType
	4, // 1: TopTasksResponse.tasks:type_name -> TaskResponse
	5, // 2: TopUsersResponse.users:type_name -> UserResponse
	1, // 3: StatisticsService.GetSpecificTaskStatistics:input_type -> SpecificTaskRequest
	3, // 4: StatisticsService.GetTopTasks:input_type -> TopRequest
	3, // 5: StatisticsService.GetTopUsers:input_type -> TopRequest
	2, // 6: StatisticsService.GetSpecificTaskStatistics:output_type -> SpecificTaskResponse
	6, // 7: StatisticsService.GetTopTasks:output_type -> TopTasksResponse
	7, // 8: StatisticsService.GetTopUsers:output_type -> TopUsersResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_statistics_proto_init() }
func file_statistics_proto_init() {
	if File_statistics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_statistics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecificTaskRequest); i {
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
		file_statistics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecificTaskResponse); i {
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
		file_statistics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopRequest); i {
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
		file_statistics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskResponse); i {
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
		file_statistics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
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
		file_statistics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopTasksResponse); i {
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
		file_statistics_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopUsersResponse); i {
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
			RawDescriptor: file_statistics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_statistics_proto_goTypes,
		DependencyIndexes: file_statistics_proto_depIdxs,
		EnumInfos:         file_statistics_proto_enumTypes,
		MessageInfos:      file_statistics_proto_msgTypes,
	}.Build()
	File_statistics_proto = out.File
	file_statistics_proto_rawDesc = nil
	file_statistics_proto_goTypes = nil
	file_statistics_proto_depIdxs = nil
}
