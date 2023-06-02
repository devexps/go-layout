// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: common/conf/bootstrap.proto

package conf

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

type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server   *Server       `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Data     *Data         `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Auth     *Auth         `protobuf:"bytes,3,opt,name=auth,proto3" json:"auth,omitempty"`
	Trace    *Tracer       `protobuf:"bytes,4,opt,name=trace,proto3" json:"trace,omitempty"`
	Logger   *Logger       `protobuf:"bytes,5,opt,name=logger,proto3" json:"logger,omitempty"`
	Registry *Registry     `protobuf:"bytes,6,opt,name=registry,proto3" json:"registry,omitempty"`
	Config   *RemoteConfig `protobuf:"bytes,8,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_conf_bootstrap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_common_conf_bootstrap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_common_conf_bootstrap_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *Bootstrap) GetTrace() *Tracer {
	if x != nil {
		return x.Trace
	}
	return nil
}

func (x *Bootstrap) GetLogger() *Logger {
	if x != nil {
		return x.Logger
	}
	return nil
}

func (x *Bootstrap) GetRegistry() *Registry {
	if x != nil {
		return x.Registry
	}
	return nil
}

func (x *Bootstrap) GetConfig() *RemoteConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_common_conf_bootstrap_proto protoreflect.FileDescriptor

var file_common_conf_bootstrap_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x62, 0x6f,
	0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x02,
	0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x2b, 0x0a, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x25, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x29, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x72, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x12, 0x2b, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e,
	0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x06, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x12, 0x31,
	0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x12, 0x31, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x42, 0x2f, 0x5a, 0x2d, 0x79, 0x6f, 0x75, 0x72, 0x5f, 0x67, 0x6f, 0x5f,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_conf_bootstrap_proto_rawDescOnce sync.Once
	file_common_conf_bootstrap_proto_rawDescData = file_common_conf_bootstrap_proto_rawDesc
)

func file_common_conf_bootstrap_proto_rawDescGZIP() []byte {
	file_common_conf_bootstrap_proto_rawDescOnce.Do(func() {
		file_common_conf_bootstrap_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_conf_bootstrap_proto_rawDescData)
	})
	return file_common_conf_bootstrap_proto_rawDescData
}

var file_common_conf_bootstrap_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_conf_bootstrap_proto_goTypes = []interface{}{
	(*Bootstrap)(nil),    // 0: common.conf.Bootstrap
	(*Server)(nil),       // 1: common.conf.Server
	(*Data)(nil),         // 2: common.conf.Data
	(*Auth)(nil),         // 3: common.conf.Auth
	(*Tracer)(nil),       // 4: common.conf.Tracer
	(*Logger)(nil),       // 5: common.conf.Logger
	(*Registry)(nil),     // 6: common.conf.Registry
	(*RemoteConfig)(nil), // 7: common.conf.RemoteConfig
}
var file_common_conf_bootstrap_proto_depIdxs = []int32{
	1, // 0: common.conf.Bootstrap.server:type_name -> common.conf.Server
	2, // 1: common.conf.Bootstrap.data:type_name -> common.conf.Data
	3, // 2: common.conf.Bootstrap.auth:type_name -> common.conf.Auth
	4, // 3: common.conf.Bootstrap.trace:type_name -> common.conf.Tracer
	5, // 4: common.conf.Bootstrap.logger:type_name -> common.conf.Logger
	6, // 5: common.conf.Bootstrap.registry:type_name -> common.conf.Registry
	7, // 6: common.conf.Bootstrap.config:type_name -> common.conf.RemoteConfig
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_common_conf_bootstrap_proto_init() }
func file_common_conf_bootstrap_proto_init() {
	if File_common_conf_bootstrap_proto != nil {
		return
	}
	file_common_conf_server_proto_init()
	file_common_conf_data_proto_init()
	file_common_conf_auth_proto_init()
	file_common_conf_trace_proto_init()
	file_common_conf_logger_proto_init()
	file_common_conf_registry_proto_init()
	file_common_conf_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_conf_bootstrap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap); i {
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
			RawDescriptor: file_common_conf_bootstrap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_conf_bootstrap_proto_goTypes,
		DependencyIndexes: file_common_conf_bootstrap_proto_depIdxs,
		MessageInfos:      file_common_conf_bootstrap_proto_msgTypes,
	}.Build()
	File_common_conf_bootstrap_proto = out.File
	file_common_conf_bootstrap_proto_rawDesc = nil
	file_common_conf_bootstrap_proto_goTypes = nil
	file_common_conf_bootstrap_proto_depIdxs = nil
}
