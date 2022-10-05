// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: protos/time.proto

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

type ClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
	Ack  int32  `protobuf:"varint,2,opt,name=Ack,proto3" json:"Ack,omitempty"`
	Seq  int32  `protobuf:"varint,3,opt,name=Seq,proto3" json:"Seq,omitempty"`
}

func (x *ClientRequest) Reset() {
	*x = ClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_time_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientRequest) ProtoMessage() {}

func (x *ClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_time_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientRequest.ProtoReflect.Descriptor instead.
func (*ClientRequest) Descriptor() ([]byte, []int) {
	return file_protos_time_proto_rawDescGZIP(), []int{0}
}

func (x *ClientRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ClientRequest) GetAck() int32 {
	if x != nil {
		return x.Ack
	}
	return 0
}

func (x *ClientRequest) GetSeq() int32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

type ServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
	Ack  int32  `protobuf:"varint,2,opt,name=Ack,proto3" json:"Ack,omitempty"`
	Seq  int32  `protobuf:"varint,3,opt,name=Seq,proto3" json:"Seq,omitempty"`
}

func (x *ServerResponse) Reset() {
	*x = ServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_time_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerResponse) ProtoMessage() {}

func (x *ServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_time_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerResponse.ProtoReflect.Descriptor instead.
func (*ServerResponse) Descriptor() ([]byte, []int) {
	return file_protos_time_proto_rawDescGZIP(), []int{1}
}

func (x *ServerResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ServerResponse) GetAck() int32 {
	if x != nil {
		return x.Ack
	}
	return 0
}

func (x *ServerResponse) GetSeq() int32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

var File_protos_time_proto protoreflect.FileDescriptor

var file_protos_time_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x47, 0x0a, 0x0d, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x41,
	0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x53, 0x65, 0x71, 0x22, 0x48, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x63,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03,
	0x53, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x53, 0x65, 0x71, 0x32, 0x49,
	0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x42, 0x5a, 0x40, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x47, 0x72, 0x75, 0x6d, 0x6c, 0x65, 0x62, 0x6f, 0x62, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x50, 0x68,
	0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x2f,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_time_proto_rawDescOnce sync.Once
	file_protos_time_proto_rawDescData = file_protos_time_proto_rawDesc
)

func file_protos_time_proto_rawDescGZIP() []byte {
	file_protos_time_proto_rawDescOnce.Do(func() {
		file_protos_time_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_time_proto_rawDescData)
	})
	return file_protos_time_proto_rawDescData
}

var file_protos_time_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_time_proto_goTypes = []interface{}{
	(*ClientRequest)(nil),  // 0: protos.ClientRequest
	(*ServerResponse)(nil), // 1: protos.ServerResponse
}
var file_protos_time_proto_depIdxs = []int32{
	0, // 0: protos.ChatService.GetTime:input_type -> protos.ClientRequest
	1, // 1: protos.ChatService.GetTime:output_type -> protos.ServerResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_time_proto_init() }
func file_protos_time_proto_init() {
	if File_protos_time_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_time_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientRequest); i {
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
		file_protos_time_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerResponse); i {
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
			RawDescriptor: file_protos_time_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_time_proto_goTypes,
		DependencyIndexes: file_protos_time_proto_depIdxs,
		MessageInfos:      file_protos_time_proto_msgTypes,
	}.Build()
	File_protos_time_proto = out.File
	file_protos_time_proto_rawDesc = nil
	file_protos_time_proto_goTypes = nil
	file_protos_time_proto_depIdxs = nil
}
