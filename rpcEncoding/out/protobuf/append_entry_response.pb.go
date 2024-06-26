// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: pkg/rpcEncoding/src/append_entry_response.proto

package protobuf

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

type AppendEntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            *string `protobuf:"bytes,1,req,name=Id" json:"Id,omitempty"`
	Success       *bool   `protobuf:"varint,2,req,name=Success" json:"Success,omitempty"`
	Term          *uint64 `protobuf:"varint,3,req,name=Term" json:"Term,omitempty"`
	LogIndexError *int32  `protobuf:"varint,4,opt,name=LogIndexError" json:"LogIndexError,omitempty"`
}

func (x *AppendEntryResponse) Reset() {
	*x = AppendEntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpcEncoding_src_append_entry_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntryResponse) ProtoMessage() {}

func (x *AppendEntryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpcEncoding_src_append_entry_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntryResponse.ProtoReflect.Descriptor instead.
func (*AppendEntryResponse) Descriptor() ([]byte, []int) {
	return file_pkg_rpcEncoding_src_append_entry_response_proto_rawDescGZIP(), []int{0}
}

func (x *AppendEntryResponse) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *AppendEntryResponse) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

func (x *AppendEntryResponse) GetTerm() uint64 {
	if x != nil && x.Term != nil {
		return *x.Term
	}
	return 0
}

func (x *AppendEntryResponse) GetLogIndexError() int32 {
	if x != nil && x.LogIndexError != nil {
		return *x.LogIndexError
	}
	return 0
}

var File_pkg_rpcEncoding_src_append_entry_response_proto protoreflect.FileDescriptor

var file_pkg_rpcEncoding_src_append_entry_response_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x5f, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x79, 0x0a, 0x13, 0x41,
	0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x65, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x02, 0x28, 0x04, 0x52, 0x04, 0x54, 0x65, 0x72, 0x6d,
	0x12, 0x24, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f,
}

var (
	file_pkg_rpcEncoding_src_append_entry_response_proto_rawDescOnce sync.Once
	file_pkg_rpcEncoding_src_append_entry_response_proto_rawDescData = file_pkg_rpcEncoding_src_append_entry_response_proto_rawDesc
)

func file_pkg_rpcEncoding_src_append_entry_response_proto_rawDescGZIP() []byte {
	file_pkg_rpcEncoding_src_append_entry_response_proto_rawDescOnce.Do(func() {
		file_pkg_rpcEncoding_src_append_entry_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_rpcEncoding_src_append_entry_response_proto_rawDescData)
	})
	return file_pkg_rpcEncoding_src_append_entry_response_proto_rawDescData
}

var file_pkg_rpcEncoding_src_append_entry_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_rpcEncoding_src_append_entry_response_proto_goTypes = []interface{}{
	(*AppendEntryResponse)(nil), // 0: protobuf.AppendEntryResponse
}
var file_pkg_rpcEncoding_src_append_entry_response_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_rpcEncoding_src_append_entry_response_proto_init() }
func file_pkg_rpcEncoding_src_append_entry_response_proto_init() {
	if File_pkg_rpcEncoding_src_append_entry_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_rpcEncoding_src_append_entry_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntryResponse); i {
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
			RawDescriptor: file_pkg_rpcEncoding_src_append_entry_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_rpcEncoding_src_append_entry_response_proto_goTypes,
		DependencyIndexes: file_pkg_rpcEncoding_src_append_entry_response_proto_depIdxs,
		MessageInfos:      file_pkg_rpcEncoding_src_append_entry_response_proto_msgTypes,
	}.Build()
	File_pkg_rpcEncoding_src_append_entry_response_proto = out.File
	file_pkg_rpcEncoding_src_append_entry_response_proto_rawDesc = nil
	file_pkg_rpcEncoding_src_append_entry_response_proto_goTypes = nil
	file_pkg_rpcEncoding_src_append_entry_response_proto_depIdxs = nil
}
