// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: pbcatalog/v1alpha1/service.proto

package catalogv1alpha1

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

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// workloads is a selector for the workloads this service should represent.
	Workloads *WorkloadSelector `protobuf:"bytes,1,opt,name=workloads,proto3" json:"workloads,omitempty"`
	// ports is the list of mappings of workload ports that this service
	// represents.
	Ports []*ServicePort `protobuf:"bytes,2,rep,name=ports,proto3" json:"ports,omitempty"`
	// virtual_ips is a list of virtual IPs for this service. This is useful when you need to set
	// an IP from an external system (like Kubernetes). This can be an IPv4 or IPv6 string.
	VirtualIps []string `protobuf:"bytes,3,rep,name=virtual_ips,json=virtualIps,proto3" json:"virtual_ips,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbcatalog_v1alpha1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_pbcatalog_v1alpha1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_pbcatalog_v1alpha1_service_proto_rawDescGZIP(), []int{0}
}

func (x *Service) GetWorkloads() *WorkloadSelector {
	if x != nil {
		return x.Workloads
	}
	return nil
}

func (x *Service) GetPorts() []*ServicePort {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *Service) GetVirtualIps() []string {
	if x != nil {
		return x.VirtualIps
	}
	return nil
}

type ServicePort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// virtual_port is the port that could only be used when transparent
	// proxy is used alongside a virtual IP or a virtual DNS address.
	// This value is ignored in other cases. Whether or not using transparent
	// proxy, this value is optional.
	VirtualPort uint32 `protobuf:"varint,1,opt,name=virtual_port,json=virtualPort,proto3" json:"virtual_port,omitempty"`
	// target_port is the name of the workload port.
	TargetPort string `protobuf:"bytes,2,opt,name=target_port,json=targetPort,proto3" json:"target_port,omitempty"`
	// protocol is the port's protocol. This should be set to "mesh"
	// if the target port is the proxy's inbound port.
	Protocol Protocol `protobuf:"varint,3,opt,name=protocol,proto3,enum=hashicorp.consul.catalog.v1alpha1.Protocol" json:"protocol,omitempty"`
}

func (x *ServicePort) Reset() {
	*x = ServicePort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbcatalog_v1alpha1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServicePort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServicePort) ProtoMessage() {}

func (x *ServicePort) ProtoReflect() protoreflect.Message {
	mi := &file_pbcatalog_v1alpha1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServicePort.ProtoReflect.Descriptor instead.
func (*ServicePort) Descriptor() ([]byte, []int) {
	return file_pbcatalog_v1alpha1_service_proto_rawDescGZIP(), []int{1}
}

func (x *ServicePort) GetVirtualPort() uint32 {
	if x != nil {
		return x.VirtualPort
	}
	return 0
}

func (x *ServicePort) GetTargetPort() string {
	if x != nil {
		return x.TargetPort
	}
	return ""
}

func (x *ServicePort) GetProtocol() Protocol {
	if x != nil {
		return x.Protocol
	}
	return Protocol_PROTOCOL_UNSPECIFIED
}

var File_pbcatalog_v1alpha1_service_proto protoreflect.FileDescriptor

var file_pbcatalog_v1alpha1_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x62, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x21, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x21, 0x70, 0x62, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x62, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x62, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x07, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x77,
	0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x44, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x69, 0x70, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x70, 0x73, 0x3a,
	0x06, 0xa2, 0x93, 0x04, 0x02, 0x08, 0x03, 0x22, 0x9a, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x76,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x47, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e,
	0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x42, 0xa9, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x62, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x43,
	0x43, 0xaa, 0x02, 0x21, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x21, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x2d, 0x48, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x43, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x24, 0x48, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbcatalog_v1alpha1_service_proto_rawDescOnce sync.Once
	file_pbcatalog_v1alpha1_service_proto_rawDescData = file_pbcatalog_v1alpha1_service_proto_rawDesc
)

func file_pbcatalog_v1alpha1_service_proto_rawDescGZIP() []byte {
	file_pbcatalog_v1alpha1_service_proto_rawDescOnce.Do(func() {
		file_pbcatalog_v1alpha1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbcatalog_v1alpha1_service_proto_rawDescData)
	})
	return file_pbcatalog_v1alpha1_service_proto_rawDescData
}

var file_pbcatalog_v1alpha1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pbcatalog_v1alpha1_service_proto_goTypes = []interface{}{
	(*Service)(nil),          // 0: hashicorp.consul.catalog.v1alpha1.Service
	(*ServicePort)(nil),      // 1: hashicorp.consul.catalog.v1alpha1.ServicePort
	(*WorkloadSelector)(nil), // 2: hashicorp.consul.catalog.v1alpha1.WorkloadSelector
	(Protocol)(0),            // 3: hashicorp.consul.catalog.v1alpha1.Protocol
}
var file_pbcatalog_v1alpha1_service_proto_depIdxs = []int32{
	2, // 0: hashicorp.consul.catalog.v1alpha1.Service.workloads:type_name -> hashicorp.consul.catalog.v1alpha1.WorkloadSelector
	1, // 1: hashicorp.consul.catalog.v1alpha1.Service.ports:type_name -> hashicorp.consul.catalog.v1alpha1.ServicePort
	3, // 2: hashicorp.consul.catalog.v1alpha1.ServicePort.protocol:type_name -> hashicorp.consul.catalog.v1alpha1.Protocol
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pbcatalog_v1alpha1_service_proto_init() }
func file_pbcatalog_v1alpha1_service_proto_init() {
	if File_pbcatalog_v1alpha1_service_proto != nil {
		return
	}
	file_pbcatalog_v1alpha1_protocol_proto_init()
	file_pbcatalog_v1alpha1_selector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pbcatalog_v1alpha1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_pbcatalog_v1alpha1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServicePort); i {
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
			RawDescriptor: file_pbcatalog_v1alpha1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbcatalog_v1alpha1_service_proto_goTypes,
		DependencyIndexes: file_pbcatalog_v1alpha1_service_proto_depIdxs,
		MessageInfos:      file_pbcatalog_v1alpha1_service_proto_msgTypes,
	}.Build()
	File_pbcatalog_v1alpha1_service_proto = out.File
	file_pbcatalog_v1alpha1_service_proto_rawDesc = nil
	file_pbcatalog_v1alpha1_service_proto_goTypes = nil
	file_pbcatalog_v1alpha1_service_proto_depIdxs = nil
}
