// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: pbmesh/v1alpha1/tcp_route.proto

package meshv1alpha1

import (
	_ "github.com/hashicorp/consul/proto-public/pbresource"
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

// NOTE: this should align to the GAMMA/gateway-api version, or at least be
// easily translatable.
//
// https://gateway-api.sigs.k8s.io/references/spec/#gateway.networking.k8s.io/v1alpha2.TCPRoute
//
// This is a Resource type.
type TCPRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ParentRefs references the resources (usually Gateways) that a Route wants
	// to be attached to. Note that the referenced parent resource needs to allow
	// this for the attachment to be complete. For Gateways, that means the
	// Gateway needs to allow attachment from Routes of this kind and namespace.
	//
	// It is invalid to reference an identical parent more than once. It is valid
	// to reference multiple distinct sections within the same parent resource,
	// such as 2 Listeners within a Gateway.
	ParentRefs []*ParentReference `protobuf:"bytes,1,rep,name=parent_refs,json=parentRefs,proto3" json:"parent_refs,omitempty"`
	// Rules are a list of TCP matchers and actions.
	Rules []*TCPRouteRule `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *TCPRoute) Reset() {
	*x = TCPRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v1alpha1_tcp_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TCPRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TCPRoute) ProtoMessage() {}

func (x *TCPRoute) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v1alpha1_tcp_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TCPRoute.ProtoReflect.Descriptor instead.
func (*TCPRoute) Descriptor() ([]byte, []int) {
	return file_pbmesh_v1alpha1_tcp_route_proto_rawDescGZIP(), []int{0}
}

func (x *TCPRoute) GetParentRefs() []*ParentReference {
	if x != nil {
		return x.ParentRefs
	}
	return nil
}

func (x *TCPRoute) GetRules() []*TCPRouteRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type TCPRouteRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// BackendRefs defines the backend(s) where matching requests should be sent.
	// If unspecified or invalid (refers to a non-existent resource or a Service
	// with no endpoints), the underlying implementation MUST actively reject
	// connection attempts to this backend. Connection rejections must respect
	// weight; if an invalid backend is requested to have 80% of connections,
	// then 80% of connections must be rejected instead.
	BackendRefs []*TCPBackendRef `protobuf:"bytes,1,rep,name=backend_refs,json=backendRefs,proto3" json:"backend_refs,omitempty"`
}

func (x *TCPRouteRule) Reset() {
	*x = TCPRouteRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v1alpha1_tcp_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TCPRouteRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TCPRouteRule) ProtoMessage() {}

func (x *TCPRouteRule) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v1alpha1_tcp_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TCPRouteRule.ProtoReflect.Descriptor instead.
func (*TCPRouteRule) Descriptor() ([]byte, []int) {
	return file_pbmesh_v1alpha1_tcp_route_proto_rawDescGZIP(), []int{1}
}

func (x *TCPRouteRule) GetBackendRefs() []*TCPBackendRef {
	if x != nil {
		return x.BackendRefs
	}
	return nil
}

type TCPBackendRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BackendRef *BackendReference `protobuf:"bytes,1,opt,name=backend_ref,json=backendRef,proto3" json:"backend_ref,omitempty"`
	// Weight specifies the proportion of requests forwarded to the referenced
	// backend. This is computed as weight/(sum of all weights in this
	// BackendRefs list). For non-zero values, there may be some epsilon from the
	// exact proportion defined here depending on the precision an implementation
	// supports. Weight is not a percentage and the sum of weights does not need
	// to equal 100.
	//
	// If only one backend is specified and it has a weight greater than 0, 100%
	// of the traffic is forwarded to that backend. If weight is set to 0, no
	// traffic should be forwarded for this entry. If unspecified, weight defaults
	// to 1.
	Weight uint32 `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *TCPBackendRef) Reset() {
	*x = TCPBackendRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v1alpha1_tcp_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TCPBackendRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TCPBackendRef) ProtoMessage() {}

func (x *TCPBackendRef) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v1alpha1_tcp_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TCPBackendRef.ProtoReflect.Descriptor instead.
func (*TCPBackendRef) Descriptor() ([]byte, []int) {
	return file_pbmesh_v1alpha1_tcp_route_proto_rawDescGZIP(), []int{2}
}

func (x *TCPBackendRef) GetBackendRef() *BackendReference {
	if x != nil {
		return x.BackendRef
	}
	return nil
}

func (x *TCPBackendRef) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

var File_pbmesh_v1alpha1_tcp_route_proto protoreflect.FileDescriptor

var file_pbmesh_v1alpha1_tcp_route_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x74, 0x63, 0x70, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x1a, 0x1c, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x70, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01,
	0x0a, 0x08, 0x54, 0x43, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2f, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x73, 0x12, 0x42, 0x0a, 0x05,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x43, 0x50,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x3a, 0x06, 0xa2, 0x93, 0x04, 0x02, 0x08, 0x03, 0x22, 0x60, 0x0a, 0x0c, 0x54, 0x43, 0x50, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x50, 0x0a, 0x0c, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x66, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x54, 0x43, 0x50, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x66, 0x52, 0x0b, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x66, 0x73, 0x22, 0x7a, 0x0a, 0x0d, 0x54, 0x43,
	0x50, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x66, 0x12, 0x51, 0x0a, 0x0b, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x30, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x66, 0x12, 0x16,
	0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x95, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0d, 0x54,
	0x63, 0x70, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x6d, 0x65, 0x73, 0x68, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x43, 0x4d, 0xaa, 0x02, 0x1e, 0x48, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x4d,
	0x65, 0x73, 0x68, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x1e, 0x48,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c,
	0x4d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x2a,
	0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x5c, 0x4d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x48, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a,
	0x4d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbmesh_v1alpha1_tcp_route_proto_rawDescOnce sync.Once
	file_pbmesh_v1alpha1_tcp_route_proto_rawDescData = file_pbmesh_v1alpha1_tcp_route_proto_rawDesc
)

func file_pbmesh_v1alpha1_tcp_route_proto_rawDescGZIP() []byte {
	file_pbmesh_v1alpha1_tcp_route_proto_rawDescOnce.Do(func() {
		file_pbmesh_v1alpha1_tcp_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbmesh_v1alpha1_tcp_route_proto_rawDescData)
	})
	return file_pbmesh_v1alpha1_tcp_route_proto_rawDescData
}

var file_pbmesh_v1alpha1_tcp_route_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pbmesh_v1alpha1_tcp_route_proto_goTypes = []interface{}{
	(*TCPRoute)(nil),         // 0: hashicorp.consul.mesh.v1alpha1.TCPRoute
	(*TCPRouteRule)(nil),     // 1: hashicorp.consul.mesh.v1alpha1.TCPRouteRule
	(*TCPBackendRef)(nil),    // 2: hashicorp.consul.mesh.v1alpha1.TCPBackendRef
	(*ParentReference)(nil),  // 3: hashicorp.consul.mesh.v1alpha1.ParentReference
	(*BackendReference)(nil), // 4: hashicorp.consul.mesh.v1alpha1.BackendReference
}
var file_pbmesh_v1alpha1_tcp_route_proto_depIdxs = []int32{
	3, // 0: hashicorp.consul.mesh.v1alpha1.TCPRoute.parent_refs:type_name -> hashicorp.consul.mesh.v1alpha1.ParentReference
	1, // 1: hashicorp.consul.mesh.v1alpha1.TCPRoute.rules:type_name -> hashicorp.consul.mesh.v1alpha1.TCPRouteRule
	2, // 2: hashicorp.consul.mesh.v1alpha1.TCPRouteRule.backend_refs:type_name -> hashicorp.consul.mesh.v1alpha1.TCPBackendRef
	4, // 3: hashicorp.consul.mesh.v1alpha1.TCPBackendRef.backend_ref:type_name -> hashicorp.consul.mesh.v1alpha1.BackendReference
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pbmesh_v1alpha1_tcp_route_proto_init() }
func file_pbmesh_v1alpha1_tcp_route_proto_init() {
	if File_pbmesh_v1alpha1_tcp_route_proto != nil {
		return
	}
	file_pbmesh_v1alpha1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pbmesh_v1alpha1_tcp_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TCPRoute); i {
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
		file_pbmesh_v1alpha1_tcp_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TCPRouteRule); i {
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
		file_pbmesh_v1alpha1_tcp_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TCPBackendRef); i {
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
			RawDescriptor: file_pbmesh_v1alpha1_tcp_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbmesh_v1alpha1_tcp_route_proto_goTypes,
		DependencyIndexes: file_pbmesh_v1alpha1_tcp_route_proto_depIdxs,
		MessageInfos:      file_pbmesh_v1alpha1_tcp_route_proto_msgTypes,
	}.Build()
	File_pbmesh_v1alpha1_tcp_route_proto = out.File
	file_pbmesh_v1alpha1_tcp_route_proto_rawDesc = nil
	file_pbmesh_v1alpha1_tcp_route_proto_goTypes = nil
	file_pbmesh_v1alpha1_tcp_route_proto_depIdxs = nil
}
