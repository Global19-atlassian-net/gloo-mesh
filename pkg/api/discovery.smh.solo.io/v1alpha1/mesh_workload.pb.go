// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/discovery/v1alpha1/mesh_workload.proto

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1/types"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//*
//The MeshWorkload is an abstraction for a workload/client which mesh-discovery has discovered to be part of a
//given mesh (i.e. its traffic is managed by an in-mesh sidecar).
//The Mesh object has references to the MeshWorkloads which belong to it.
type MeshWorkloadSpec struct {
	// workload_type specifies configuration to the specific type of workload
	//
	// Types that are valid to be assigned to WorkloadType:
	//	*MeshWorkloadSpec_Kubernetes
	WorkloadType isMeshWorkloadSpec_WorkloadType `protobuf_oneof:"workload_type"`
	// The mesh with which this workload is associated
	Mesh *v1.ObjectRef `protobuf:"bytes,4,opt,name=mesh,proto3" json:"mesh,omitempty"`
	// Appmesh specific metadata
	AppMesh              *MeshWorkloadSpec_AppMesh `protobuf:"bytes,5,opt,name=app_mesh,json=appMesh,proto3" json:"app_mesh,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *MeshWorkloadSpec) Reset()         { *m = MeshWorkloadSpec{} }
func (m *MeshWorkloadSpec) String() string { return proto.CompactTextString(m) }
func (*MeshWorkloadSpec) ProtoMessage()    {}
func (*MeshWorkloadSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_20997816445c0b02, []int{0}
}
func (m *MeshWorkloadSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshWorkloadSpec.Unmarshal(m, b)
}
func (m *MeshWorkloadSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshWorkloadSpec.Marshal(b, m, deterministic)
}
func (m *MeshWorkloadSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshWorkloadSpec.Merge(m, src)
}
func (m *MeshWorkloadSpec) XXX_Size() int {
	return xxx_messageInfo_MeshWorkloadSpec.Size(m)
}
func (m *MeshWorkloadSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshWorkloadSpec.DiscardUnknown(m)
}

var xxx_messageInfo_MeshWorkloadSpec proto.InternalMessageInfo

type isMeshWorkloadSpec_WorkloadType interface {
	isMeshWorkloadSpec_WorkloadType()
	Equal(interface{}) bool
}

type MeshWorkloadSpec_Kubernetes struct {
	Kubernetes *MeshWorkloadSpec_KubernertesWorkload `protobuf:"bytes,1,opt,name=kubernetes,proto3,oneof" json:"kubernetes,omitempty"`
}

func (*MeshWorkloadSpec_Kubernetes) isMeshWorkloadSpec_WorkloadType() {}

func (m *MeshWorkloadSpec) GetWorkloadType() isMeshWorkloadSpec_WorkloadType {
	if m != nil {
		return m.WorkloadType
	}
	return nil
}

func (m *MeshWorkloadSpec) GetKubernetes() *MeshWorkloadSpec_KubernertesWorkload {
	if x, ok := m.GetWorkloadType().(*MeshWorkloadSpec_Kubernetes); ok {
		return x.Kubernetes
	}
	return nil
}

func (m *MeshWorkloadSpec) GetMesh() *v1.ObjectRef {
	if m != nil {
		return m.Mesh
	}
	return nil
}

func (m *MeshWorkloadSpec) GetAppMesh() *MeshWorkloadSpec_AppMesh {
	if m != nil {
		return m.AppMesh
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MeshWorkloadSpec) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MeshWorkloadSpec_Kubernetes)(nil),
	}
}

// information describing a Kubernetes-based workload (e.g. a Deployment or DaemonSet)
type MeshWorkloadSpec_KubernertesWorkload struct {
	//*
	//Resource ref to the underlying kubernetes controller which is managing the pods associated with the workloads.
	//It has the generic name controller as it can represent a deployment, daemonset, or statefulset.
	Controller *v1.ClusterObjectRef `protobuf:"bytes,1,opt,name=controller,proto3" json:"controller,omitempty"`
	//*
	//these are the labels directly from the pods that this controller owns
	//NB: these labels are read directly from the pod template metadata.labels
	//defined in the workload spec.
	//We need these to determine which services are backed by this workload, and
	//the service backing is determined by the pod labels.
	PodLabels map[string]string `protobuf:"bytes,2,rep,name=pod_labels,json=podLabels,proto3" json:"pod_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Service account attached to the pods owned by this controller
	ServiceAccountName   string   `protobuf:"bytes,3,opt,name=service_account_name,json=serviceAccountName,proto3" json:"service_account_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshWorkloadSpec_KubernertesWorkload) Reset()         { *m = MeshWorkloadSpec_KubernertesWorkload{} }
func (m *MeshWorkloadSpec_KubernertesWorkload) String() string { return proto.CompactTextString(m) }
func (*MeshWorkloadSpec_KubernertesWorkload) ProtoMessage()    {}
func (*MeshWorkloadSpec_KubernertesWorkload) Descriptor() ([]byte, []int) {
	return fileDescriptor_20997816445c0b02, []int{0, 0}
}
func (m *MeshWorkloadSpec_KubernertesWorkload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshWorkloadSpec_KubernertesWorkload.Unmarshal(m, b)
}
func (m *MeshWorkloadSpec_KubernertesWorkload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshWorkloadSpec_KubernertesWorkload.Marshal(b, m, deterministic)
}
func (m *MeshWorkloadSpec_KubernertesWorkload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshWorkloadSpec_KubernertesWorkload.Merge(m, src)
}
func (m *MeshWorkloadSpec_KubernertesWorkload) XXX_Size() int {
	return xxx_messageInfo_MeshWorkloadSpec_KubernertesWorkload.Size(m)
}
func (m *MeshWorkloadSpec_KubernertesWorkload) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshWorkloadSpec_KubernertesWorkload.DiscardUnknown(m)
}

var xxx_messageInfo_MeshWorkloadSpec_KubernertesWorkload proto.InternalMessageInfo

func (m *MeshWorkloadSpec_KubernertesWorkload) GetController() *v1.ClusterObjectRef {
	if m != nil {
		return m.Controller
	}
	return nil
}

func (m *MeshWorkloadSpec_KubernertesWorkload) GetPodLabels() map[string]string {
	if m != nil {
		return m.PodLabels
	}
	return nil
}

func (m *MeshWorkloadSpec_KubernertesWorkload) GetServiceAccountName() string {
	if m != nil {
		return m.ServiceAccountName
	}
	return ""
}

// information relevant to AppMesh-injected workloads
type MeshWorkloadSpec_AppMesh struct {
	// The value of the env var APPMESH_VIRTUAL_NODE_NAME on the Appmesh envoy proxy container
	VirtualNodeName string `protobuf:"bytes,1,opt,name=virtual_node_name,json=virtualNodeName,proto3" json:"virtual_node_name,omitempty"`
	// Needed for declaring Appmesh VirtualNode listeners
	Ports                []*MeshWorkloadSpec_AppMesh_ContainerPort `protobuf:"bytes,2,rep,name=ports,proto3" json:"ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *MeshWorkloadSpec_AppMesh) Reset()         { *m = MeshWorkloadSpec_AppMesh{} }
func (m *MeshWorkloadSpec_AppMesh) String() string { return proto.CompactTextString(m) }
func (*MeshWorkloadSpec_AppMesh) ProtoMessage()    {}
func (*MeshWorkloadSpec_AppMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_20997816445c0b02, []int{0, 1}
}
func (m *MeshWorkloadSpec_AppMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshWorkloadSpec_AppMesh.Unmarshal(m, b)
}
func (m *MeshWorkloadSpec_AppMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshWorkloadSpec_AppMesh.Marshal(b, m, deterministic)
}
func (m *MeshWorkloadSpec_AppMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshWorkloadSpec_AppMesh.Merge(m, src)
}
func (m *MeshWorkloadSpec_AppMesh) XXX_Size() int {
	return xxx_messageInfo_MeshWorkloadSpec_AppMesh.Size(m)
}
func (m *MeshWorkloadSpec_AppMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshWorkloadSpec_AppMesh.DiscardUnknown(m)
}

var xxx_messageInfo_MeshWorkloadSpec_AppMesh proto.InternalMessageInfo

func (m *MeshWorkloadSpec_AppMesh) GetVirtualNodeName() string {
	if m != nil {
		return m.VirtualNodeName
	}
	return ""
}

func (m *MeshWorkloadSpec_AppMesh) GetPorts() []*MeshWorkloadSpec_AppMesh_ContainerPort {
	if m != nil {
		return m.Ports
	}
	return nil
}

// k8s application container ports
type MeshWorkloadSpec_AppMesh_ContainerPort struct {
	Port                 uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	Protocol             string   `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshWorkloadSpec_AppMesh_ContainerPort) Reset() {
	*m = MeshWorkloadSpec_AppMesh_ContainerPort{}
}
func (m *MeshWorkloadSpec_AppMesh_ContainerPort) String() string { return proto.CompactTextString(m) }
func (*MeshWorkloadSpec_AppMesh_ContainerPort) ProtoMessage()    {}
func (*MeshWorkloadSpec_AppMesh_ContainerPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_20997816445c0b02, []int{0, 1, 0}
}
func (m *MeshWorkloadSpec_AppMesh_ContainerPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshWorkloadSpec_AppMesh_ContainerPort.Unmarshal(m, b)
}
func (m *MeshWorkloadSpec_AppMesh_ContainerPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshWorkloadSpec_AppMesh_ContainerPort.Marshal(b, m, deterministic)
}
func (m *MeshWorkloadSpec_AppMesh_ContainerPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshWorkloadSpec_AppMesh_ContainerPort.Merge(m, src)
}
func (m *MeshWorkloadSpec_AppMesh_ContainerPort) XXX_Size() int {
	return xxx_messageInfo_MeshWorkloadSpec_AppMesh_ContainerPort.Size(m)
}
func (m *MeshWorkloadSpec_AppMesh_ContainerPort) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshWorkloadSpec_AppMesh_ContainerPort.DiscardUnknown(m)
}

var xxx_messageInfo_MeshWorkloadSpec_AppMesh_ContainerPort proto.InternalMessageInfo

func (m *MeshWorkloadSpec_AppMesh_ContainerPort) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *MeshWorkloadSpec_AppMesh_ContainerPort) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

type MeshWorkloadStatus struct {
	// The observed generation of the MeshWorkload.
	// When this matches the MeshWorkload's metadata.generation, it indicates that mesh-networking
	// has reconciled the latest version of the MeshWorkload.
	ObservedGeneration   int64    `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshWorkloadStatus) Reset()         { *m = MeshWorkloadStatus{} }
func (m *MeshWorkloadStatus) String() string { return proto.CompactTextString(m) }
func (*MeshWorkloadStatus) ProtoMessage()    {}
func (*MeshWorkloadStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_20997816445c0b02, []int{1}
}
func (m *MeshWorkloadStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshWorkloadStatus.Unmarshal(m, b)
}
func (m *MeshWorkloadStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshWorkloadStatus.Marshal(b, m, deterministic)
}
func (m *MeshWorkloadStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshWorkloadStatus.Merge(m, src)
}
func (m *MeshWorkloadStatus) XXX_Size() int {
	return xxx_messageInfo_MeshWorkloadStatus.Size(m)
}
func (m *MeshWorkloadStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshWorkloadStatus.DiscardUnknown(m)
}

var xxx_messageInfo_MeshWorkloadStatus proto.InternalMessageInfo

func (m *MeshWorkloadStatus) GetObservedGeneration() int64 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func init() {
	proto.RegisterType((*MeshWorkloadSpec)(nil), "discovery.smh.solo.io.MeshWorkloadSpec")
	proto.RegisterType((*MeshWorkloadSpec_KubernertesWorkload)(nil), "discovery.smh.solo.io.MeshWorkloadSpec.KubernertesWorkload")
	proto.RegisterMapType((map[string]string)(nil), "discovery.smh.solo.io.MeshWorkloadSpec.KubernertesWorkload.PodLabelsEntry")
	proto.RegisterType((*MeshWorkloadSpec_AppMesh)(nil), "discovery.smh.solo.io.MeshWorkloadSpec.AppMesh")
	proto.RegisterType((*MeshWorkloadSpec_AppMesh_ContainerPort)(nil), "discovery.smh.solo.io.MeshWorkloadSpec.AppMesh.ContainerPort")
	proto.RegisterType((*MeshWorkloadStatus)(nil), "discovery.smh.solo.io.MeshWorkloadStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/discovery/v1alpha1/mesh_workload.proto", fileDescriptor_20997816445c0b02)
}

var fileDescriptor_20997816445c0b02 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xfd, 0x9c, 0xf4, 0x77, 0xaa, 0x7e, 0x2d, 0xdb, 0x22, 0x45, 0x16, 0x42, 0x55, 0xb9, 0xa9,
	0x90, 0xba, 0x6e, 0xc2, 0x0d, 0xe2, 0x47, 0xa8, 0x8d, 0x2a, 0x50, 0x80, 0x10, 0xb9, 0x17, 0x48,
	0x48, 0xc8, 0x5a, 0xdb, 0x43, 0x6c, 0xe2, 0x78, 0x56, 0xbb, 0x6b, 0xa3, 0xbc, 0x00, 0xcf, 0xc2,
	0x0b, 0xf0, 0x14, 0xbc, 0x05, 0x4f, 0x82, 0xbc, 0x76, 0x42, 0x82, 0x72, 0x51, 0xc4, 0x95, 0x67,
	0xe7, 0xec, 0x39, 0x73, 0x66, 0xbc, 0xbb, 0x30, 0x1c, 0xa7, 0x26, 0x29, 0x42, 0x1e, 0xd1, 0xd4,
	0xd3, 0x94, 0xd1, 0x79, 0x4a, 0x9e, 0x46, 0x55, 0xa6, 0x11, 0x9e, 0x4f, 0x51, 0x27, 0xe7, 0x49,
	0x11, 0x7a, 0x42, 0xa6, 0x5e, 0x9c, 0xea, 0x88, 0x4a, 0x54, 0x33, 0xaf, 0xec, 0x8a, 0x4c, 0x26,
	0xa2, 0xeb, 0x55, 0x78, 0xf0, 0x85, 0xd4, 0x24, 0x23, 0x11, 0x73, 0xa9, 0xc8, 0x10, 0xbb, 0xbb,
	0xd8, 0xc7, 0xf5, 0x34, 0xe1, 0x95, 0x24, 0x4f, 0xc9, 0xe5, 0xeb, 0xca, 0x4c, 0xca, 0x9e, 0x95,
	0x8e, 0x48, 0xa1, 0x57, 0x76, 0xed, 0xb7, 0x96, 0x71, 0xd7, 0x7b, 0x68, 0x36, 0x36, 0xe5, 0xb5,
	0x11, 0xa6, 0xd0, 0x0d, 0xe1, 0x78, 0x4c, 0x63, 0xb2, 0xa1, 0x57, 0x45, 0x75, 0xf6, 0xf4, 0xeb,
	0x16, 0x1c, 0xbe, 0x45, 0x9d, 0xbc, 0x6f, 0x4c, 0xde, 0x48, 0x8c, 0xd8, 0x47, 0x80, 0x49, 0x11,
	0xa2, 0xca, 0xd1, 0xa0, 0xee, 0x38, 0x27, 0xce, 0xd9, 0x5e, 0xef, 0x29, 0x5f, 0xeb, 0x9b, 0xff,
	0x49, 0xe6, 0xaf, 0x6b, 0xa6, 0x32, 0xa8, 0xe7, 0xf9, 0x57, 0xff, 0xf9, 0x4b, 0x82, 0xec, 0x02,
	0x36, 0x2a, 0xd3, 0x9d, 0x0d, 0x2b, 0x7c, 0x8f, 0xdb, 0xae, 0xaa, 0x5e, 0x17, 0xa2, 0xef, 0xc2,
	0xcf, 0x18, 0x19, 0x1f, 0x3f, 0xf9, 0x76, 0x27, 0x1b, 0xc0, 0x8e, 0x90, 0x32, 0xb0, 0xac, 0x4d,
	0xcb, 0xf2, 0x6e, 0x6b, 0xe7, 0x52, 0xca, 0x2a, 0xe7, 0x6f, 0x8b, 0x3a, 0x70, 0xbf, 0xb7, 0xe0,
	0x68, 0x8d, 0x47, 0xd6, 0x07, 0x88, 0x28, 0x37, 0x8a, 0xb2, 0x0c, 0x55, 0xd3, 0xf4, 0x83, 0x35,
	0xde, 0xfa, 0x59, 0xa1, 0x0d, 0xaa, 0xdf, 0x16, 0x97, 0x68, 0x2c, 0x05, 0x90, 0x14, 0x07, 0x99,
	0x08, 0x31, 0xd3, 0x9d, 0xd6, 0x49, 0xfb, 0x6c, 0xaf, 0x37, 0xf8, 0x87, 0xc9, 0xf1, 0x11, 0xc5,
	0x6f, 0xac, 0xd8, 0x75, 0x6e, 0xd4, 0xcc, 0xdf, 0x95, 0xf3, 0x35, 0xbb, 0x80, 0xe3, 0xe6, 0x08,
	0x04, 0x22, 0x8a, 0xa8, 0xc8, 0x4d, 0x90, 0x8b, 0x29, 0x76, 0xda, 0x27, 0xce, 0xd9, 0xae, 0xcf,
	0x1a, 0xec, 0xb2, 0x86, 0x86, 0x62, 0x8a, 0xee, 0x33, 0xf8, 0x7f, 0x55, 0x8e, 0x1d, 0x42, 0x7b,
	0x82, 0x33, 0xdb, 0xec, 0xae, 0x5f, 0x85, 0xec, 0x18, 0x36, 0x4b, 0x91, 0x15, 0xd8, 0x69, 0xd9,
	0x5c, 0xbd, 0x78, 0xd2, 0x7a, 0xec, 0xb8, 0x3f, 0x1c, 0xd8, 0x6e, 0x86, 0xc9, 0x1e, 0xc2, 0x9d,
	0x32, 0x55, 0xa6, 0x10, 0x59, 0x90, 0x53, 0x8c, 0x75, 0xe1, 0x5a, 0xe5, 0xa0, 0x01, 0x86, 0x14,
	0x63, 0x55, 0x95, 0xdd, 0xc0, 0xa6, 0x24, 0x65, 0xe6, 0xd3, 0x78, 0xfe, 0x97, 0x3f, 0x8e, 0xf7,
	0x29, 0x37, 0x22, 0xcd, 0x51, 0x8d, 0x48, 0x19, 0xbf, 0xd6, 0x72, 0x5f, 0xc0, 0xfe, 0x4a, 0x9e,
	0x31, 0xd8, 0xa8, 0x10, 0x6b, 0x62, 0xdf, 0xb7, 0x31, 0x73, 0x61, 0xc7, 0x1e, 0xf2, 0x88, 0xb2,
	0xa6, 0x9d, 0xc5, 0xfa, 0xea, 0x00, 0xf6, 0xe7, 0xf7, 0x32, 0x30, 0x33, 0x89, 0xa7, 0xd7, 0xc0,
	0x56, 0x2c, 0xd8, 0xab, 0xc3, 0x3c, 0x38, 0xa2, 0xb0, 0x1a, 0x25, 0xc6, 0xc1, 0x18, 0x73, 0x54,
	0xc2, 0xa4, 0x94, 0xdb, 0x2a, 0x6d, 0x9f, 0xcd, 0xa1, 0x97, 0x0b, 0xe4, 0x6a, 0xf4, 0xed, 0xe7,
	0x7d, 0xe7, 0xc3, 0xe0, 0x36, 0x6f, 0x86, 0x9c, 0x8c, 0x57, 0xdf, 0x8d, 0xe5, 0x79, 0x2c, 0x2e,
	0x71, 0xb8, 0x65, 0x3d, 0x3f, 0xfa, 0x15, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x6e, 0xa2, 0x1e, 0x88,
	0x04, 0x00, 0x00,
}

func (this *MeshWorkloadSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshWorkloadSpec)
	if !ok {
		that2, ok := that.(MeshWorkloadSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.WorkloadType == nil {
		if this.WorkloadType != nil {
			return false
		}
	} else if this.WorkloadType == nil {
		return false
	} else if !this.WorkloadType.Equal(that1.WorkloadType) {
		return false
	}
	if !this.Mesh.Equal(that1.Mesh) {
		return false
	}
	if !this.AppMesh.Equal(that1.AppMesh) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshWorkloadSpec_Kubernetes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshWorkloadSpec_Kubernetes)
	if !ok {
		that2, ok := that.(MeshWorkloadSpec_Kubernetes)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Kubernetes.Equal(that1.Kubernetes) {
		return false
	}
	return true
}
func (this *MeshWorkloadSpec_KubernertesWorkload) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshWorkloadSpec_KubernertesWorkload)
	if !ok {
		that2, ok := that.(MeshWorkloadSpec_KubernertesWorkload)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Controller.Equal(that1.Controller) {
		return false
	}
	if len(this.PodLabels) != len(that1.PodLabels) {
		return false
	}
	for i := range this.PodLabels {
		if this.PodLabels[i] != that1.PodLabels[i] {
			return false
		}
	}
	if this.ServiceAccountName != that1.ServiceAccountName {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshWorkloadSpec_AppMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshWorkloadSpec_AppMesh)
	if !ok {
		that2, ok := that.(MeshWorkloadSpec_AppMesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.VirtualNodeName != that1.VirtualNodeName {
		return false
	}
	if len(this.Ports) != len(that1.Ports) {
		return false
	}
	for i := range this.Ports {
		if !this.Ports[i].Equal(that1.Ports[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshWorkloadSpec_AppMesh_ContainerPort) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshWorkloadSpec_AppMesh_ContainerPort)
	if !ok {
		that2, ok := that.(MeshWorkloadSpec_AppMesh_ContainerPort)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if this.Protocol != that1.Protocol {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshWorkloadStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshWorkloadStatus)
	if !ok {
		that2, ok := that.(MeshWorkloadStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ObservedGeneration != that1.ObservedGeneration {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
