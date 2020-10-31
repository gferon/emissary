// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/extensions/filters/http/response_map/v4alpha/response_map.proto

package envoy_extensions_filters_http_response_map_v4alpha

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v4alpha "github.com/datawire/ambassador/pkg/api/envoy/config/accesslog/v4alpha"
	v4alpha1 "github.com/datawire/ambassador/pkg/api/envoy/config/core/v4alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/struct"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type ResponseMapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter             *v4alpha.AccessLogFilter           `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	StatusCode         *wrappers.UInt32Value              `protobuf:"bytes,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Body               *v4alpha1.DataSource               `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	BodyFormatOverride *v4alpha1.SubstitutionFormatString `protobuf:"bytes,4,opt,name=body_format_override,json=bodyFormatOverride,proto3" json:"body_format_override,omitempty"`
}

func (x *ResponseMapper) Reset() {
	*x = ResponseMapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMapper) ProtoMessage() {}

func (x *ResponseMapper) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMapper.ProtoReflect.Descriptor instead.
func (*ResponseMapper) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseMapper) GetFilter() *v4alpha.AccessLogFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ResponseMapper) GetStatusCode() *wrappers.UInt32Value {
	if x != nil {
		return x.StatusCode
	}
	return nil
}

func (x *ResponseMapper) GetBody() *v4alpha1.DataSource {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *ResponseMapper) GetBodyFormatOverride() *v4alpha1.SubstitutionFormatString {
	if x != nil {
		return x.BodyFormatOverride
	}
	return nil
}

type ResponseMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mappers    []*ResponseMapper                  `protobuf:"bytes,1,rep,name=mappers,proto3" json:"mappers,omitempty"`
	BodyFormat *v4alpha1.SubstitutionFormatString `protobuf:"bytes,2,opt,name=body_format,json=bodyFormat,proto3" json:"body_format,omitempty"`
}

func (x *ResponseMap) Reset() {
	*x = ResponseMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMap) ProtoMessage() {}

func (x *ResponseMap) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMap.ProtoReflect.Descriptor instead.
func (*ResponseMap) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseMap) GetMappers() []*ResponseMapper {
	if x != nil {
		return x.Mappers
	}
	return nil
}

func (x *ResponseMap) GetBodyFormat() *v4alpha1.SubstitutionFormatString {
	if x != nil {
		return x.BodyFormat
	}
	return nil
}

type ResponseMapPerRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Override:
	//	*ResponseMapPerRoute_Disabled
	Override isResponseMapPerRoute_Override `protobuf_oneof:"override"`
}

func (x *ResponseMapPerRoute) Reset() {
	*x = ResponseMapPerRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMapPerRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMapPerRoute) ProtoMessage() {}

func (x *ResponseMapPerRoute) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMapPerRoute.ProtoReflect.Descriptor instead.
func (*ResponseMapPerRoute) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_rawDescGZIP(), []int{2}
}

func (m *ResponseMapPerRoute) GetOverride() isResponseMapPerRoute_Override {
	if m != nil {
		return m.Override
	}
	return nil
}

func (x *ResponseMapPerRoute) GetDisabled() bool {
	if x, ok := x.GetOverride().(*ResponseMapPerRoute_Disabled); ok {
		return x.Disabled
	}
	return false
}

type isResponseMapPerRoute_Override interface {
	isResponseMapPerRoute_Override()
}

type ResponseMapPerRoute_Disabled struct {
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3,oneof"`
}

func (*ResponseMapPerRoute_Disabled) isResponseMapPerRoute_Override() {}

var File_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_rawDesc = []byte{
	0x0a, 0x45, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x2f, 0x76, 0x34, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x61,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x6d, 0x61, 0x70, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x3a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x75, 0x62,
	0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x03, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x0b,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x0b, 0xfa, 0x42, 0x08, 0x2a, 0x06, 0x10, 0xd8, 0x04, 0x28, 0xc8, 0x01, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x12, 0x65, 0x0a, 0x14, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x33, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x12, 0x62, 0x6f, 0x64, 0x79, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x3a, 0x43, 0x9a, 0xc5, 0x88, 0x1e,
	0x3e, 0x0a, 0x3c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x33,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x22,
	0x83, 0x02, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x61, 0x70, 0x12,
	0x5c, 0x0a, 0x07, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x42, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x34,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x12, 0x54, 0x0a,
	0x0b, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x33, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x62, 0x6f, 0x64, 0x79, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x3a, 0x40, 0x9a, 0xc5, 0x88, 0x1e, 0x3b, 0x0a, 0x39, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x61, 0x70, 0x22, 0x97, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x61, 0x70, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x25, 0x0a,
	0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x6a, 0x02, 0x08, 0x01, 0x48, 0x00, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x3a, 0x48, 0x9a, 0xc5, 0x88, 0x1e, 0x43, 0x0a, 0x41, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4d, 0x61, 0x70, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x42, 0x0f,
	0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42,
	0x5e, 0x0a, 0x40, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x34, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x42, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x61, 0x70,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x03, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_rawDescData = file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_rawDesc
)

func file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_rawDescData
}

var file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_goTypes = []interface{}{
	(*ResponseMapper)(nil),                    // 0: envoy.extensions.filters.http.response_map.v4alpha.ResponseMapper
	(*ResponseMap)(nil),                       // 1: envoy.extensions.filters.http.response_map.v4alpha.ResponseMap
	(*ResponseMapPerRoute)(nil),               // 2: envoy.extensions.filters.http.response_map.v4alpha.ResponseMapPerRoute
	(*v4alpha.AccessLogFilter)(nil),           // 3: envoy.config.accesslog.v4alpha.AccessLogFilter
	(*wrappers.UInt32Value)(nil),              // 4: google.protobuf.UInt32Value
	(*v4alpha1.DataSource)(nil),               // 5: envoy.config.core.v4alpha.DataSource
	(*v4alpha1.SubstitutionFormatString)(nil), // 6: envoy.config.core.v4alpha.SubstitutionFormatString
}
var file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_depIdxs = []int32{
	3, // 0: envoy.extensions.filters.http.response_map.v4alpha.ResponseMapper.filter:type_name -> envoy.config.accesslog.v4alpha.AccessLogFilter
	4, // 1: envoy.extensions.filters.http.response_map.v4alpha.ResponseMapper.status_code:type_name -> google.protobuf.UInt32Value
	5, // 2: envoy.extensions.filters.http.response_map.v4alpha.ResponseMapper.body:type_name -> envoy.config.core.v4alpha.DataSource
	6, // 3: envoy.extensions.filters.http.response_map.v4alpha.ResponseMapper.body_format_override:type_name -> envoy.config.core.v4alpha.SubstitutionFormatString
	0, // 4: envoy.extensions.filters.http.response_map.v4alpha.ResponseMap.mappers:type_name -> envoy.extensions.filters.http.response_map.v4alpha.ResponseMapper
	6, // 5: envoy.extensions.filters.http.response_map.v4alpha.ResponseMap.body_format:type_name -> envoy.config.core.v4alpha.SubstitutionFormatString
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_init() }
func file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_init() {
	if File_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMapper); i {
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
		file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMap); i {
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
		file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMapPerRoute); i {
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
	file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ResponseMapPerRoute_Disabled)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto = out.File
	file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_rawDesc = nil
	file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_goTypes = nil
	file_envoy_extensions_filters_http_response_map_v4alpha_response_map_proto_depIdxs = nil
}
