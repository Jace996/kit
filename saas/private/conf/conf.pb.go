// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: saas/private/conf/conf.proto

package conf

import (
	_ "github.com/go-saas/kit/pkg/blob"
	conf "github.com/go-saas/kit/pkg/conf"
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

	Data     *conf.Data      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Security *conf.Security  `protobuf:"bytes,3,opt,name=security,proto3" json:"security,omitempty"`
	Services *conf.Services  `protobuf:"bytes,4,opt,name=services,proto3" json:"services,omitempty"`
	Logging  *conf.Logging   `protobuf:"bytes,6,opt,name=logging,proto3" json:"logging,omitempty"`
	Tracing  *conf.Tracers   `protobuf:"bytes,7,opt,name=tracing,proto3" json:"tracing,omitempty"`
	App      *conf.AppConfig `protobuf:"bytes,8,opt,name=app,proto3" json:"app,omitempty"`
	Dev      *conf.Dev       `protobuf:"bytes,9,opt,name=dev,proto3" json:"dev,omitempty"`
	Saas     *SaasConf       `protobuf:"bytes,20,opt,name=saas,proto3" json:"saas,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saas_private_conf_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_saas_private_conf_conf_proto_msgTypes[0]
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
	return file_saas_private_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetData() *conf.Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetSecurity() *conf.Security {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *Bootstrap) GetServices() *conf.Services {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *Bootstrap) GetLogging() *conf.Logging {
	if x != nil {
		return x.Logging
	}
	return nil
}

func (x *Bootstrap) GetTracing() *conf.Tracers {
	if x != nil {
		return x.Tracing
	}
	return nil
}

func (x *Bootstrap) GetApp() *conf.AppConfig {
	if x != nil {
		return x.App
	}
	return nil
}

func (x *Bootstrap) GetDev() *conf.Dev {
	if x != nil {
		return x.Dev
	}
	return nil
}

func (x *Bootstrap) GetSaas() *SaasConf {
	if x != nil {
		return x.Saas
	}
	return nil
}

type SaasConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database     []*DatabaseTemplate `protobuf:"bytes,1,rep,name=database,proto3" json:"database,omitempty"`
	TenantCookie *conf.Cookie        `protobuf:"bytes,2,opt,name=tenant_cookie,json=tenantCookie,proto3" json:"tenant_cookie,omitempty"`
}

func (x *SaasConf) Reset() {
	*x = SaasConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saas_private_conf_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaasConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaasConf) ProtoMessage() {}

func (x *SaasConf) ProtoReflect() protoreflect.Message {
	mi := &file_saas_private_conf_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaasConf.ProtoReflect.Descriptor instead.
func (*SaasConf) Descriptor() ([]byte, []int) {
	return file_saas_private_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *SaasConf) GetDatabase() []*DatabaseTemplate {
	if x != nil {
		return x.Database
	}
	return nil
}

func (x *SaasConf) GetTenantCookie() *conf.Cookie {
	if x != nil {
		return x.TenantCookie
	}
	return nil
}

type DatabaseTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Template string `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *DatabaseTemplate) Reset() {
	*x = DatabaseTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saas_private_conf_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseTemplate) ProtoMessage() {}

func (x *DatabaseTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_saas_private_conf_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseTemplate.ProtoReflect.Descriptor instead.
func (*DatabaseTemplate) Descriptor() ([]byte, []int) {
	return file_saas_private_conf_conf_proto_rawDescGZIP(), []int{2}
}

func (x *DatabaseTemplate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DatabaseTemplate) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

var File_saas_private_conf_conf_proto protoreflect.FileDescriptor

var file_saas_private_conf_conf_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x73, 0x61, 0x61, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x1a, 0x0f, 0x63,
	0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f,
	0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc2, 0x02, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x1e, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a,
	0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52,
	0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x27,
	0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x72, 0x73, 0x52, 0x07,
	0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x41, 0x70, 0x70, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x1b, 0x0a, 0x03, 0x64, 0x65,
	0x76, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x44,
	0x65, 0x76, 0x52, 0x03, 0x64, 0x65, 0x76, 0x12, 0x2b, 0x0a, 0x04, 0x73, 0x61, 0x61, 0x73, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x53, 0x61, 0x61, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x04,
	0x73, 0x61, 0x61, 0x73, 0x22, 0x7a, 0x0a, 0x08, 0x53, 0x61, 0x61, 0x73, 0x43, 0x6f, 0x6e, 0x66,
	0x12, 0x3b, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x0d, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x43, 0x6f, 0x6f, 0x6b,
	0x69, 0x65, 0x52, 0x0c, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65,
	0x22, 0x42, 0x0a, 0x10, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x6b, 0x69, 0x74, 0x2f, 0x73,
	0x61, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_saas_private_conf_conf_proto_rawDescOnce sync.Once
	file_saas_private_conf_conf_proto_rawDescData = file_saas_private_conf_conf_proto_rawDesc
)

func file_saas_private_conf_conf_proto_rawDescGZIP() []byte {
	file_saas_private_conf_conf_proto_rawDescOnce.Do(func() {
		file_saas_private_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_saas_private_conf_conf_proto_rawDescData)
	})
	return file_saas_private_conf_conf_proto_rawDescData
}

var file_saas_private_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_saas_private_conf_conf_proto_goTypes = []interface{}{
	(*Bootstrap)(nil),        // 0: saas.internal.Bootstrap
	(*SaasConf)(nil),         // 1: saas.internal.SaasConf
	(*DatabaseTemplate)(nil), // 2: saas.internal.DatabaseTemplate
	(*conf.Data)(nil),        // 3: conf.Data
	(*conf.Security)(nil),    // 4: conf.Security
	(*conf.Services)(nil),    // 5: conf.Services
	(*conf.Logging)(nil),     // 6: conf.Logging
	(*conf.Tracers)(nil),     // 7: conf.Tracers
	(*conf.AppConfig)(nil),   // 8: conf.AppConfig
	(*conf.Dev)(nil),         // 9: conf.Dev
	(*conf.Cookie)(nil),      // 10: conf.Cookie
}
var file_saas_private_conf_conf_proto_depIdxs = []int32{
	3,  // 0: saas.internal.Bootstrap.data:type_name -> conf.Data
	4,  // 1: saas.internal.Bootstrap.security:type_name -> conf.Security
	5,  // 2: saas.internal.Bootstrap.services:type_name -> conf.Services
	6,  // 3: saas.internal.Bootstrap.logging:type_name -> conf.Logging
	7,  // 4: saas.internal.Bootstrap.tracing:type_name -> conf.Tracers
	8,  // 5: saas.internal.Bootstrap.app:type_name -> conf.AppConfig
	9,  // 6: saas.internal.Bootstrap.dev:type_name -> conf.Dev
	1,  // 7: saas.internal.Bootstrap.saas:type_name -> saas.internal.SaasConf
	2,  // 8: saas.internal.SaasConf.database:type_name -> saas.internal.DatabaseTemplate
	10, // 9: saas.internal.SaasConf.tenant_cookie:type_name -> conf.Cookie
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_saas_private_conf_conf_proto_init() }
func file_saas_private_conf_conf_proto_init() {
	if File_saas_private_conf_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_saas_private_conf_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_saas_private_conf_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaasConf); i {
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
		file_saas_private_conf_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseTemplate); i {
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
			RawDescriptor: file_saas_private_conf_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_saas_private_conf_conf_proto_goTypes,
		DependencyIndexes: file_saas_private_conf_conf_proto_depIdxs,
		MessageInfos:      file_saas_private_conf_conf_proto_msgTypes,
	}.Build()
	File_saas_private_conf_conf_proto = out.File
	file_saas_private_conf_conf_proto_rawDesc = nil
	file_saas_private_conf_conf_proto_goTypes = nil
	file_saas_private_conf_conf_proto_depIdxs = nil
}
