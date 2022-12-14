// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: messagepb/message.proto

package messagepb

import (
	simplepb "github.com/IguoChan/proto-prj/genproto/simplepb"
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

type MessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Req *simplepb.SimpleRequest `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
}

func (x *MessageRequest) Reset() {
	*x = MessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagepb_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRequest) ProtoMessage() {}

func (x *MessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messagepb_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageRequest.ProtoReflect.Descriptor instead.
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return file_messagepb_message_proto_rawDescGZIP(), []int{0}
}

func (x *MessageRequest) GetReq() *simplepb.SimpleRequest {
	if x != nil {
		return x.Req
	}
	return nil
}

var File_messagepb_message_proto protoreflect.FileDescriptor

var file_messagepb_message_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x70, 0x62, 0x1a, 0x15, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2f, 0x73,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x0e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a,
	0x03, 0x72, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x03, 0x72, 0x65, 0x71, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x67, 0x75, 0x6f, 0x43, 0x68, 0x61, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x72, 0x6a, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messagepb_message_proto_rawDescOnce sync.Once
	file_messagepb_message_proto_rawDescData = file_messagepb_message_proto_rawDesc
)

func file_messagepb_message_proto_rawDescGZIP() []byte {
	file_messagepb_message_proto_rawDescOnce.Do(func() {
		file_messagepb_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_messagepb_message_proto_rawDescData)
	})
	return file_messagepb_message_proto_rawDescData
}

var file_messagepb_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_messagepb_message_proto_goTypes = []interface{}{
	(*MessageRequest)(nil),         // 0: messagepb.MessageRequest
	(*simplepb.SimpleRequest)(nil), // 1: simplepb.SimpleRequest
}
var file_messagepb_message_proto_depIdxs = []int32{
	1, // 0: messagepb.MessageRequest.req:type_name -> simplepb.SimpleRequest
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_messagepb_message_proto_init() }
func file_messagepb_message_proto_init() {
	if File_messagepb_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messagepb_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageRequest); i {
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
			RawDescriptor: file_messagepb_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messagepb_message_proto_goTypes,
		DependencyIndexes: file_messagepb_message_proto_depIdxs,
		MessageInfos:      file_messagepb_message_proto_msgTypes,
	}.Build()
	File_messagepb_message_proto = out.File
	file_messagepb_message_proto_rawDesc = nil
	file_messagepb_message_proto_goTypes = nil
	file_messagepb_message_proto_depIdxs = nil
}
