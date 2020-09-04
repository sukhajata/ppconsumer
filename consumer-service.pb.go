// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.0
// source: consumer-service.proto

package ppconsumer

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

type Consumer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username      string          `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email         string          `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Installations []*Installation `protobuf:"bytes,3,rep,name=installations,proto3" json:"installations,omitempty"`
	Type          string          `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Consumer) Reset() {
	*x = Consumer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Consumer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Consumer) ProtoMessage() {}

func (x *Consumer) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Consumer.ProtoReflect.Descriptor instead.
func (*Consumer) Descriptor() ([]byte, []int) {
	return file_consumer_service_proto_rawDescGZIP(), []int{0}
}

func (x *Consumer) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Consumer) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Consumer) GetInstallations() []*Installation {
	if x != nil {
		return x.Installations
	}
	return nil
}

func (x *Consumer) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Installation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDNumber  string   `protobuf:"bytes,1,opt,name=IDNumber,proto3" json:"IDNumber,omitempty"`
	DeviceEUI []string `protobuf:"bytes,2,rep,name=deviceEUI,proto3" json:"deviceEUI,omitempty"`
}

func (x *Installation) Reset() {
	*x = Installation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Installation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Installation) ProtoMessage() {}

func (x *Installation) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Installation.ProtoReflect.Descriptor instead.
func (*Installation) Descriptor() ([]byte, []int) {
	return file_consumer_service_proto_rawDescGZIP(), []int{1}
}

func (x *Installation) GetIDNumber() string {
	if x != nil {
		return x.IDNumber
	}
	return ""
}

func (x *Installation) GetDeviceEUI() []string {
	if x != nil {
		return x.DeviceEUI
	}
	return nil
}

type ConsumerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ConsumerRequest) Reset() {
	*x = ConsumerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumerRequest) ProtoMessage() {}

func (x *ConsumerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumerRequest.ProtoReflect.Descriptor instead.
func (*ConsumerRequest) Descriptor() ([]byte, []int) {
	return file_consumer_service_proto_rawDescGZIP(), []int{2}
}

func (x *ConsumerRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type CreateConsumerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateConsumerResponse) Reset() {
	*x = CreateConsumerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateConsumerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConsumerResponse) ProtoMessage() {}

func (x *CreateConsumerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConsumerResponse.ProtoReflect.Descriptor instead.
func (*CreateConsumerResponse) Descriptor() ([]byte, []int) {
	return file_consumer_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateConsumerResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply string `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_consumer_service_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

var File_consumer_service_proto protoreflect.FileDescriptor

var file_consumer_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x22, 0x90, 0x01, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x3e, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x70, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x48, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x44, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x44, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x45, 0x55, 0x49,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x45, 0x55,
	0x49, 0x22, 0x2d, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x34, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x20, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xaa, 0x02, 0x0a, 0x0f, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x12, 0x14,
	0x2e, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x1a, 0x22, 0x2e, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x70,
	0x70, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x1a, 0x14, 0x2e, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x70, 0x70, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x22, 0x00, 0x12, 0x45,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x12, 0x1b, 0x2e, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x70, 0x70, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_consumer_service_proto_rawDescOnce sync.Once
	file_consumer_service_proto_rawDescData = file_consumer_service_proto_rawDesc
)

func file_consumer_service_proto_rawDescGZIP() []byte {
	file_consumer_service_proto_rawDescOnce.Do(func() {
		file_consumer_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_consumer_service_proto_rawDescData)
	})
	return file_consumer_service_proto_rawDescData
}

var file_consumer_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_consumer_service_proto_goTypes = []interface{}{
	(*Consumer)(nil),               // 0: ppconsumer.Consumer
	(*Installation)(nil),           // 1: ppconsumer.Installation
	(*ConsumerRequest)(nil),        // 2: ppconsumer.ConsumerRequest
	(*CreateConsumerResponse)(nil), // 3: ppconsumer.CreateConsumerResponse
	(*Response)(nil),               // 4: ppconsumer.Response
}
var file_consumer_service_proto_depIdxs = []int32{
	1, // 0: ppconsumer.Consumer.installations:type_name -> ppconsumer.Installation
	0, // 1: ppconsumer.ConsumerService.CreateConsumer:input_type -> ppconsumer.Consumer
	0, // 2: ppconsumer.ConsumerService.UpdateConsumer:input_type -> ppconsumer.Consumer
	2, // 3: ppconsumer.ConsumerService.GetConsumer:input_type -> ppconsumer.ConsumerRequest
	2, // 4: ppconsumer.ConsumerService.DeleteConsumer:input_type -> ppconsumer.ConsumerRequest
	3, // 5: ppconsumer.ConsumerService.CreateConsumer:output_type -> ppconsumer.CreateConsumerResponse
	0, // 6: ppconsumer.ConsumerService.UpdateConsumer:output_type -> ppconsumer.Consumer
	0, // 7: ppconsumer.ConsumerService.GetConsumer:output_type -> ppconsumer.Consumer
	4, // 8: ppconsumer.ConsumerService.DeleteConsumer:output_type -> ppconsumer.Response
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_consumer_service_proto_init() }
func file_consumer_service_proto_init() {
	if File_consumer_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_consumer_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Consumer); i {
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
		file_consumer_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Installation); i {
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
		file_consumer_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumerRequest); i {
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
		file_consumer_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateConsumerResponse); i {
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
		file_consumer_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_consumer_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_consumer_service_proto_goTypes,
		DependencyIndexes: file_consumer_service_proto_depIdxs,
		MessageInfos:      file_consumer_service_proto_msgTypes,
	}.Build()
	File_consumer_service_proto = out.File
	file_consumer_service_proto_rawDesc = nil
	file_consumer_service_proto_goTypes = nil
	file_consumer_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConsumerServiceClient is the client API for ConsumerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConsumerServiceClient interface {
	CreateConsumer(ctx context.Context, in *Consumer, opts ...grpc.CallOption) (*CreateConsumerResponse, error)
	UpdateConsumer(ctx context.Context, in *Consumer, opts ...grpc.CallOption) (*Consumer, error)
	GetConsumer(ctx context.Context, in *ConsumerRequest, opts ...grpc.CallOption) (*Consumer, error)
	DeleteConsumer(ctx context.Context, in *ConsumerRequest, opts ...grpc.CallOption) (*Response, error)
}

type consumerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConsumerServiceClient(cc grpc.ClientConnInterface) ConsumerServiceClient {
	return &consumerServiceClient{cc}
}

func (c *consumerServiceClient) CreateConsumer(ctx context.Context, in *Consumer, opts ...grpc.CallOption) (*CreateConsumerResponse, error) {
	out := new(CreateConsumerResponse)
	err := c.cc.Invoke(ctx, "/ppconsumer.ConsumerService/CreateConsumer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) UpdateConsumer(ctx context.Context, in *Consumer, opts ...grpc.CallOption) (*Consumer, error) {
	out := new(Consumer)
	err := c.cc.Invoke(ctx, "/ppconsumer.ConsumerService/UpdateConsumer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) GetConsumer(ctx context.Context, in *ConsumerRequest, opts ...grpc.CallOption) (*Consumer, error) {
	out := new(Consumer)
	err := c.cc.Invoke(ctx, "/ppconsumer.ConsumerService/GetConsumer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) DeleteConsumer(ctx context.Context, in *ConsumerRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ppconsumer.ConsumerService/DeleteConsumer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsumerServiceServer is the server API for ConsumerService service.
type ConsumerServiceServer interface {
	CreateConsumer(context.Context, *Consumer) (*CreateConsumerResponse, error)
	UpdateConsumer(context.Context, *Consumer) (*Consumer, error)
	GetConsumer(context.Context, *ConsumerRequest) (*Consumer, error)
	DeleteConsumer(context.Context, *ConsumerRequest) (*Response, error)
}

// UnimplementedConsumerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConsumerServiceServer struct {
}

func (*UnimplementedConsumerServiceServer) CreateConsumer(context.Context, *Consumer) (*CreateConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConsumer not implemented")
}
func (*UnimplementedConsumerServiceServer) UpdateConsumer(context.Context, *Consumer) (*Consumer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConsumer not implemented")
}
func (*UnimplementedConsumerServiceServer) GetConsumer(context.Context, *ConsumerRequest) (*Consumer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsumer not implemented")
}
func (*UnimplementedConsumerServiceServer) DeleteConsumer(context.Context, *ConsumerRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConsumer not implemented")
}

func RegisterConsumerServiceServer(s *grpc.Server, srv ConsumerServiceServer) {
	s.RegisterService(&_ConsumerService_serviceDesc, srv)
}

func _ConsumerService_CreateConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Consumer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).CreateConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ppconsumer.ConsumerService/CreateConsumer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).CreateConsumer(ctx, req.(*Consumer))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_UpdateConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Consumer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).UpdateConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ppconsumer.ConsumerService/UpdateConsumer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).UpdateConsumer(ctx, req.(*Consumer))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_GetConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).GetConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ppconsumer.ConsumerService/GetConsumer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).GetConsumer(ctx, req.(*ConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_DeleteConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).DeleteConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ppconsumer.ConsumerService/DeleteConsumer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).DeleteConsumer(ctx, req.(*ConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConsumerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ppconsumer.ConsumerService",
	HandlerType: (*ConsumerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConsumer",
			Handler:    _ConsumerService_CreateConsumer_Handler,
		},
		{
			MethodName: "UpdateConsumer",
			Handler:    _ConsumerService_UpdateConsumer_Handler,
		},
		{
			MethodName: "GetConsumer",
			Handler:    _ConsumerService_GetConsumer_Handler,
		},
		{
			MethodName: "DeleteConsumer",
			Handler:    _ConsumerService_DeleteConsumer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consumer-service.proto",
}