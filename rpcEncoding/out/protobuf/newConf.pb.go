// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: newConf.proto

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

type NewConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewConfiguration []string `protobuf:"bytes,1,rep,name=new_configuration,json=newConfiguration,proto3" json:"new_configuration,omitempty"`
}

func (x *NewConf) Reset() {
	*x = NewConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_newConf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewConf) ProtoMessage() {}

func (x *NewConf) ProtoReflect() protoreflect.Message {
	mi := &file_newConf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewConf.ProtoReflect.Descriptor instead.
func (*NewConf) Descriptor() ([]byte, []int) {
	return file_newConf_proto_rawDescGZIP(), []int{0}
}

func (x *NewConf) GetNewConfiguration() []string {
	if x != nil {
		return x.NewConfiguration
	}
	return nil
}

var File_newConf_proto protoreflect.FileDescriptor

var file_newConf_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6e, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x36, 0x0a, 0x07, 0x4e, 0x65, 0x77,
	0x43, 0x6f, 0x6e, 0x66, 0x12, 0x2b, 0x0a, 0x11, 0x6e, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x10, 0x6e, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_newConf_proto_rawDescOnce sync.Once
	file_newConf_proto_rawDescData = file_newConf_proto_rawDesc
)

func file_newConf_proto_rawDescGZIP() []byte {
	file_newConf_proto_rawDescOnce.Do(func() {
		file_newConf_proto_rawDescData = protoimpl.X.CompressGZIP(file_newConf_proto_rawDescData)
	})
	return file_newConf_proto_rawDescData
}

var file_newConf_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_newConf_proto_goTypes = []interface{}{
	(*NewConf)(nil), // 0: protobuf.NewConf
}
var file_newConf_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_newConf_proto_init() }
func file_newConf_proto_init() {
	if File_newConf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_newConf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewConf); i {
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
			RawDescriptor: file_newConf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_newConf_proto_goTypes,
		DependencyIndexes: file_newConf_proto_depIdxs,
		MessageInfos:      file_newConf_proto_msgTypes,
	}.Build()
	File_newConf_proto = out.File
	file_newConf_proto_rawDesc = nil
	file_newConf_proto_goTypes = nil
	file_newConf_proto_depIdxs = nil
}
