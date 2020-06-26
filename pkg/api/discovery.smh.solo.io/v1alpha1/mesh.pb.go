// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/discovery/v1alpha1/mesh.proto

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
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

//
//Meshes represent a currently registered service mesh.
type MeshSpec struct {
	// Types that are valid to be assigned to MeshType:
	//	*MeshSpec_Istio_
	//	*MeshSpec_AwsAppMesh_
	//	*MeshSpec_Linkerd
	//	*MeshSpec_ConsulConnect
	MeshType             isMeshSpec_MeshType `protobuf_oneof:"mesh_type"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MeshSpec) Reset()         { *m = MeshSpec{} }
func (m *MeshSpec) String() string { return proto.CompactTextString(m) }
func (*MeshSpec) ProtoMessage()    {}
func (*MeshSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_345fdec2269fec48, []int{0}
}
func (m *MeshSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshSpec.Unmarshal(m, b)
}
func (m *MeshSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshSpec.Marshal(b, m, deterministic)
}
func (m *MeshSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshSpec.Merge(m, src)
}
func (m *MeshSpec) XXX_Size() int {
	return xxx_messageInfo_MeshSpec.Size(m)
}
func (m *MeshSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshSpec.DiscardUnknown(m)
}

var xxx_messageInfo_MeshSpec proto.InternalMessageInfo

type isMeshSpec_MeshType interface {
	isMeshSpec_MeshType()
	Equal(interface{}) bool
}

type MeshSpec_Istio_ struct {
	Istio *MeshSpec_Istio `protobuf:"bytes,1,opt,name=istio,proto3,oneof" json:"istio,omitempty"`
}
type MeshSpec_AwsAppMesh_ struct {
	AwsAppMesh *MeshSpec_AwsAppMesh `protobuf:"bytes,2,opt,name=aws_app_mesh,json=awsAppMesh,proto3,oneof" json:"aws_app_mesh,omitempty"`
}
type MeshSpec_Linkerd struct {
	Linkerd *MeshSpec_LinkerdMesh `protobuf:"bytes,3,opt,name=linkerd,proto3,oneof" json:"linkerd,omitempty"`
}
type MeshSpec_ConsulConnect struct {
	ConsulConnect *MeshSpec_ConsulConnectMesh `protobuf:"bytes,4,opt,name=consul_connect,json=consulConnect,proto3,oneof" json:"consul_connect,omitempty"`
}

func (*MeshSpec_Istio_) isMeshSpec_MeshType()        {}
func (*MeshSpec_AwsAppMesh_) isMeshSpec_MeshType()   {}
func (*MeshSpec_Linkerd) isMeshSpec_MeshType()       {}
func (*MeshSpec_ConsulConnect) isMeshSpec_MeshType() {}

func (m *MeshSpec) GetMeshType() isMeshSpec_MeshType {
	if m != nil {
		return m.MeshType
	}
	return nil
}

func (m *MeshSpec) GetIstio() *MeshSpec_Istio {
	if x, ok := m.GetMeshType().(*MeshSpec_Istio_); ok {
		return x.Istio
	}
	return nil
}

func (m *MeshSpec) GetAwsAppMesh() *MeshSpec_AwsAppMesh {
	if x, ok := m.GetMeshType().(*MeshSpec_AwsAppMesh_); ok {
		return x.AwsAppMesh
	}
	return nil
}

func (m *MeshSpec) GetLinkerd() *MeshSpec_LinkerdMesh {
	if x, ok := m.GetMeshType().(*MeshSpec_Linkerd); ok {
		return x.Linkerd
	}
	return nil
}

func (m *MeshSpec) GetConsulConnect() *MeshSpec_ConsulConnectMesh {
	if x, ok := m.GetMeshType().(*MeshSpec_ConsulConnect); ok {
		return x.ConsulConnect
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MeshSpec) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MeshSpec_Istio_)(nil),
		(*MeshSpec_AwsAppMesh_)(nil),
		(*MeshSpec_Linkerd)(nil),
		(*MeshSpec_ConsulConnect)(nil),
	}
}

// Mesh object representing an installed Istio control plane
type MeshSpec_Istio struct {
	// where istio control plane components are installed.
	Installation *MeshSpec_MeshInstallation `protobuf:"bytes,1,opt,name=installation,proto3" json:"installation,omitempty"`
	// configuration for Istio Citadel, Istio's security component.
	CitadelInfo          *MeshSpec_Istio_CitadelInfo `protobuf:"bytes,2,opt,name=citadel_info,json=citadelInfo,proto3" json:"citadel_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *MeshSpec_Istio) Reset()         { *m = MeshSpec_Istio{} }
func (m *MeshSpec_Istio) String() string { return proto.CompactTextString(m) }
func (*MeshSpec_Istio) ProtoMessage()    {}
func (*MeshSpec_Istio) Descriptor() ([]byte, []int) {
	return fileDescriptor_345fdec2269fec48, []int{0, 0}
}
func (m *MeshSpec_Istio) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshSpec_Istio.Unmarshal(m, b)
}
func (m *MeshSpec_Istio) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshSpec_Istio.Marshal(b, m, deterministic)
}
func (m *MeshSpec_Istio) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshSpec_Istio.Merge(m, src)
}
func (m *MeshSpec_Istio) XXX_Size() int {
	return xxx_messageInfo_MeshSpec_Istio.Size(m)
}
func (m *MeshSpec_Istio) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshSpec_Istio.DiscardUnknown(m)
}

var xxx_messageInfo_MeshSpec_Istio proto.InternalMessageInfo

func (m *MeshSpec_Istio) GetInstallation() *MeshSpec_MeshInstallation {
	if m != nil {
		return m.Installation
	}
	return nil
}

func (m *MeshSpec_Istio) GetCitadelInfo() *MeshSpec_Istio_CitadelInfo {
	if m != nil {
		return m.CitadelInfo
	}
	return nil
}

type MeshSpec_Istio_CitadelInfo struct {
	//
	//Istio trust domain used for https/spiffe identity.
	//https://spiffe.io/spiffe/concepts/#trust-domain
	//https://istio.io/docs/reference/glossary/#identity
	//
	//If empty will default to "cluster.local"
	TrustDomain string `protobuf:"bytes,1,opt,name=trust_domain,json=trustDomain,proto3" json:"trust_domain,omitempty"`
	//
	//istio-citadel service account, used to determine identity for the Istio CA cert.
	//If empty will default to "istio-citadel"
	CitadelServiceAccount string   `protobuf:"bytes,2,opt,name=citadel_service_account,json=citadelServiceAccount,proto3" json:"citadel_service_account,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *MeshSpec_Istio_CitadelInfo) Reset()         { *m = MeshSpec_Istio_CitadelInfo{} }
func (m *MeshSpec_Istio_CitadelInfo) String() string { return proto.CompactTextString(m) }
func (*MeshSpec_Istio_CitadelInfo) ProtoMessage()    {}
func (*MeshSpec_Istio_CitadelInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_345fdec2269fec48, []int{0, 0, 0}
}
func (m *MeshSpec_Istio_CitadelInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshSpec_Istio_CitadelInfo.Unmarshal(m, b)
}
func (m *MeshSpec_Istio_CitadelInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshSpec_Istio_CitadelInfo.Marshal(b, m, deterministic)
}
func (m *MeshSpec_Istio_CitadelInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshSpec_Istio_CitadelInfo.Merge(m, src)
}
func (m *MeshSpec_Istio_CitadelInfo) XXX_Size() int {
	return xxx_messageInfo_MeshSpec_Istio_CitadelInfo.Size(m)
}
func (m *MeshSpec_Istio_CitadelInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshSpec_Istio_CitadelInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MeshSpec_Istio_CitadelInfo proto.InternalMessageInfo

func (m *MeshSpec_Istio_CitadelInfo) GetTrustDomain() string {
	if m != nil {
		return m.TrustDomain
	}
	return ""
}

func (m *MeshSpec_Istio_CitadelInfo) GetCitadelServiceAccount() string {
	if m != nil {
		return m.CitadelServiceAccount
	}
	return ""
}

// Mesh object representing AWS AppMesh
type MeshSpec_AwsAppMesh struct {
	// AWS name for the AppMesh instance, must be unique across the AWS account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The AWS region the AWS App Mesh control plane resources exist in.
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// The AWS Account ID associated with the Mesh. Populated at REST API registration time.
	AwsAccountId string `protobuf:"bytes,3,opt,name=aws_account_id,json=awsAccountId,proto3" json:"aws_account_id,omitempty"`
	// The k8s clusters on which sidecars for this AppMesh instance have been discovered.
	Clusters             []string `protobuf:"bytes,4,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshSpec_AwsAppMesh) Reset()         { *m = MeshSpec_AwsAppMesh{} }
func (m *MeshSpec_AwsAppMesh) String() string { return proto.CompactTextString(m) }
func (*MeshSpec_AwsAppMesh) ProtoMessage()    {}
func (*MeshSpec_AwsAppMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_345fdec2269fec48, []int{0, 1}
}
func (m *MeshSpec_AwsAppMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshSpec_AwsAppMesh.Unmarshal(m, b)
}
func (m *MeshSpec_AwsAppMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshSpec_AwsAppMesh.Marshal(b, m, deterministic)
}
func (m *MeshSpec_AwsAppMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshSpec_AwsAppMesh.Merge(m, src)
}
func (m *MeshSpec_AwsAppMesh) XXX_Size() int {
	return xxx_messageInfo_MeshSpec_AwsAppMesh.Size(m)
}
func (m *MeshSpec_AwsAppMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshSpec_AwsAppMesh.DiscardUnknown(m)
}

var xxx_messageInfo_MeshSpec_AwsAppMesh proto.InternalMessageInfo

func (m *MeshSpec_AwsAppMesh) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MeshSpec_AwsAppMesh) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *MeshSpec_AwsAppMesh) GetAwsAccountId() string {
	if m != nil {
		return m.AwsAccountId
	}
	return ""
}

func (m *MeshSpec_AwsAppMesh) GetClusters() []string {
	if m != nil {
		return m.Clusters
	}
	return nil
}

// Mesh object representing an installed Linkerd control plane
type MeshSpec_LinkerdMesh struct {
	Installation *MeshSpec_MeshInstallation `protobuf:"bytes,1,opt,name=installation,proto3" json:"installation,omitempty"`
	// The cluster domain suffix this Linkerd mesh is configured with.
	// See https://linkerd.io/2/tasks/using-custom-domain/ for info
	ClusterDomain        string   `protobuf:"bytes,2,opt,name=cluster_domain,json=clusterDomain,proto3" json:"cluster_domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshSpec_LinkerdMesh) Reset()         { *m = MeshSpec_LinkerdMesh{} }
func (m *MeshSpec_LinkerdMesh) String() string { return proto.CompactTextString(m) }
func (*MeshSpec_LinkerdMesh) ProtoMessage()    {}
func (*MeshSpec_LinkerdMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_345fdec2269fec48, []int{0, 2}
}
func (m *MeshSpec_LinkerdMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshSpec_LinkerdMesh.Unmarshal(m, b)
}
func (m *MeshSpec_LinkerdMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshSpec_LinkerdMesh.Marshal(b, m, deterministic)
}
func (m *MeshSpec_LinkerdMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshSpec_LinkerdMesh.Merge(m, src)
}
func (m *MeshSpec_LinkerdMesh) XXX_Size() int {
	return xxx_messageInfo_MeshSpec_LinkerdMesh.Size(m)
}
func (m *MeshSpec_LinkerdMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshSpec_LinkerdMesh.DiscardUnknown(m)
}

var xxx_messageInfo_MeshSpec_LinkerdMesh proto.InternalMessageInfo

func (m *MeshSpec_LinkerdMesh) GetInstallation() *MeshSpec_MeshInstallation {
	if m != nil {
		return m.Installation
	}
	return nil
}

func (m *MeshSpec_LinkerdMesh) GetClusterDomain() string {
	if m != nil {
		return m.ClusterDomain
	}
	return ""
}

type MeshSpec_ConsulConnectMesh struct {
	Installation         *MeshSpec_MeshInstallation `protobuf:"bytes,1,opt,name=installation,proto3" json:"installation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *MeshSpec_ConsulConnectMesh) Reset()         { *m = MeshSpec_ConsulConnectMesh{} }
func (m *MeshSpec_ConsulConnectMesh) String() string { return proto.CompactTextString(m) }
func (*MeshSpec_ConsulConnectMesh) ProtoMessage()    {}
func (*MeshSpec_ConsulConnectMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_345fdec2269fec48, []int{0, 3}
}
func (m *MeshSpec_ConsulConnectMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshSpec_ConsulConnectMesh.Unmarshal(m, b)
}
func (m *MeshSpec_ConsulConnectMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshSpec_ConsulConnectMesh.Marshal(b, m, deterministic)
}
func (m *MeshSpec_ConsulConnectMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshSpec_ConsulConnectMesh.Merge(m, src)
}
func (m *MeshSpec_ConsulConnectMesh) XXX_Size() int {
	return xxx_messageInfo_MeshSpec_ConsulConnectMesh.Size(m)
}
func (m *MeshSpec_ConsulConnectMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshSpec_ConsulConnectMesh.DiscardUnknown(m)
}

var xxx_messageInfo_MeshSpec_ConsulConnectMesh proto.InternalMessageInfo

func (m *MeshSpec_ConsulConnectMesh) GetInstallation() *MeshSpec_MeshInstallation {
	if m != nil {
		return m.Installation
	}
	return nil
}

//
//The cluster on which the control plane for this mesh is deployed.
//Not all MeshTypes have a MeshInstallation. Only self-hosted
//control planes such as Istio and Linkerd will have installation metadata.
type MeshSpec_MeshInstallation struct {
	// Namespace in which the control plane has been installed.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Cluster in which the control plane has been installed.
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// Version of the Mesh that has been installed.
	// Determined using the image tag on the
	// Mesh's primary control plane image (e.g. the istio-pilot image tag).
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshSpec_MeshInstallation) Reset()         { *m = MeshSpec_MeshInstallation{} }
func (m *MeshSpec_MeshInstallation) String() string { return proto.CompactTextString(m) }
func (*MeshSpec_MeshInstallation) ProtoMessage()    {}
func (*MeshSpec_MeshInstallation) Descriptor() ([]byte, []int) {
	return fileDescriptor_345fdec2269fec48, []int{0, 4}
}
func (m *MeshSpec_MeshInstallation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshSpec_MeshInstallation.Unmarshal(m, b)
}
func (m *MeshSpec_MeshInstallation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshSpec_MeshInstallation.Marshal(b, m, deterministic)
}
func (m *MeshSpec_MeshInstallation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshSpec_MeshInstallation.Merge(m, src)
}
func (m *MeshSpec_MeshInstallation) XXX_Size() int {
	return xxx_messageInfo_MeshSpec_MeshInstallation.Size(m)
}
func (m *MeshSpec_MeshInstallation) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshSpec_MeshInstallation.DiscardUnknown(m)
}

var xxx_messageInfo_MeshSpec_MeshInstallation proto.InternalMessageInfo

func (m *MeshSpec_MeshInstallation) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *MeshSpec_MeshInstallation) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *MeshSpec_MeshInstallation) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type MeshStatus struct {
	// The observed generation of the Mesh.
	// When this matches the Mesh's metadata.generation, it indicates that mesh-networking
	// has reconciled the latest version of the Mesh.
	ObservedGeneration   int64    `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshStatus) Reset()         { *m = MeshStatus{} }
func (m *MeshStatus) String() string { return proto.CompactTextString(m) }
func (*MeshStatus) ProtoMessage()    {}
func (*MeshStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_345fdec2269fec48, []int{1}
}
func (m *MeshStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshStatus.Unmarshal(m, b)
}
func (m *MeshStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshStatus.Marshal(b, m, deterministic)
}
func (m *MeshStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshStatus.Merge(m, src)
}
func (m *MeshStatus) XXX_Size() int {
	return xxx_messageInfo_MeshStatus.Size(m)
}
func (m *MeshStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshStatus.DiscardUnknown(m)
}

var xxx_messageInfo_MeshStatus proto.InternalMessageInfo

func (m *MeshStatus) GetObservedGeneration() int64 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func init() {
	proto.RegisterType((*MeshSpec)(nil), "discovery.smh.solo.io.MeshSpec")
	proto.RegisterType((*MeshSpec_Istio)(nil), "discovery.smh.solo.io.MeshSpec.Istio")
	proto.RegisterType((*MeshSpec_Istio_CitadelInfo)(nil), "discovery.smh.solo.io.MeshSpec.Istio.CitadelInfo")
	proto.RegisterType((*MeshSpec_AwsAppMesh)(nil), "discovery.smh.solo.io.MeshSpec.AwsAppMesh")
	proto.RegisterType((*MeshSpec_LinkerdMesh)(nil), "discovery.smh.solo.io.MeshSpec.LinkerdMesh")
	proto.RegisterType((*MeshSpec_ConsulConnectMesh)(nil), "discovery.smh.solo.io.MeshSpec.ConsulConnectMesh")
	proto.RegisterType((*MeshSpec_MeshInstallation)(nil), "discovery.smh.solo.io.MeshSpec.MeshInstallation")
	proto.RegisterType((*MeshStatus)(nil), "discovery.smh.solo.io.MeshStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/discovery/v1alpha1/mesh.proto", fileDescriptor_345fdec2269fec48)
}

var fileDescriptor_345fdec2269fec48 = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6e, 0xd4, 0x3c,
	0x14, 0xed, 0xf4, 0x7f, 0xee, 0xb4, 0xd5, 0xf7, 0x19, 0x0a, 0x51, 0x84, 0xaa, 0x52, 0x51, 0xa9,
	0x02, 0x35, 0x61, 0x8a, 0xc4, 0xae, 0x8b, 0xb6, 0xa0, 0x76, 0x10, 0x20, 0x94, 0xb2, 0xea, 0x26,
	0xf2, 0x38, 0x6e, 0x62, 0x35, 0x13, 0x5b, 0xb6, 0x93, 0xaa, 0x3c, 0x02, 0xef, 0xc0, 0x9e, 0x17,
	0xe1, 0x45, 0x78, 0x12, 0xe4, 0x9f, 0xcc, 0x4c, 0xa1, 0xd2, 0x74, 0xd3, 0x55, 0x7c, 0xef, 0xf1,
	0x3d, 0x3e, 0x3e, 0xbe, 0xb9, 0xf0, 0x3e, 0x67, 0xba, 0xa8, 0x87, 0x11, 0xe1, 0xa3, 0x58, 0xf1,
	0x92, 0xef, 0x33, 0x1e, 0x2b, 0x2a, 0x1b, 0x46, 0xe8, 0xfe, 0x88, 0xaa, 0x62, 0xbf, 0xa8, 0x87,
	0x31, 0x16, 0x2c, 0xce, 0x98, 0x22, 0xbc, 0xa1, 0xf2, 0x26, 0x6e, 0xfa, 0xb8, 0x14, 0x05, 0xee,
	0xc7, 0x06, 0x8f, 0x84, 0xe4, 0x9a, 0xa3, 0xcd, 0x31, 0x1c, 0xa9, 0x51, 0x11, 0x19, 0xa6, 0x88,
	0xf1, 0x30, 0xba, 0x8b, 0xfd, 0xaa, 0x39, 0xb0, 0x8c, 0x84, 0x4b, 0x1a, 0x37, 0x7d, 0xfb, 0x75,
	0x34, 0xe1, 0x56, 0xce, 0x79, 0x5e, 0xd2, 0xd8, 0x46, 0xc3, 0xfa, 0x32, 0xbe, 0x96, 0x58, 0x08,
	0x2a, 0x95, 0xc7, 0x1f, 0xe7, 0x3c, 0xe7, 0x76, 0x19, 0x9b, 0x95, 0xcb, 0xee, 0xfc, 0x5a, 0x85,
	0xd5, 0x4f, 0x54, 0x15, 0xe7, 0x82, 0x12, 0x74, 0x08, 0x4b, 0x4c, 0x69, 0xc6, 0x83, 0xce, 0x76,
	0x67, 0xaf, 0x77, 0xb0, 0x1b, 0xdd, 0xa9, 0x2c, 0x6a, 0xf7, 0x47, 0x03, 0xb3, 0xf9, 0x6c, 0x2e,
	0x71, 0x55, 0xe8, 0x33, 0xac, 0xe1, 0x6b, 0x95, 0x62, 0x21, 0x52, 0x73, 0xbd, 0x60, 0xde, 0xb2,
	0xbc, 0x9c, 0xc5, 0x72, 0x74, 0xad, 0x8e, 0x84, 0x30, 0xe1, 0xd9, 0x5c, 0x02, 0x78, 0x1c, 0xa1,
	0x53, 0x58, 0x29, 0x59, 0x75, 0x45, 0x65, 0x16, 0x2c, 0x58, 0xaa, 0x57, 0xb3, 0xa8, 0x3e, 0xba,
	0xed, 0x9e, 0xab, 0xad, 0x46, 0x17, 0xb0, 0x41, 0x78, 0xa5, 0xea, 0x32, 0x25, 0xbc, 0xaa, 0x28,
	0xd1, 0xc1, 0xa2, 0xe5, 0xeb, 0xcf, 0xe2, 0x3b, 0xb1, 0x55, 0x27, 0xae, 0xc8, 0xb3, 0xae, 0x93,
	0xe9, 0x64, 0xf8, 0x63, 0x1e, 0x96, 0xac, 0x0f, 0xe8, 0x2b, 0xac, 0xb1, 0x4a, 0x69, 0x5c, 0x96,
	0x58, 0x33, 0x5e, 0x79, 0x13, 0x5f, 0xcf, 0x3a, 0xc3, 0x2c, 0x06, 0x53, 0x75, 0xc9, 0x2d, 0x16,
	0xc3, 0x4a, 0x98, 0xc6, 0x19, 0x2d, 0x53, 0x56, 0x5d, 0x72, 0x6f, 0x6a, 0xff, 0x5e, 0x4f, 0x13,
	0x9d, 0xb8, 0xca, 0x41, 0x75, 0xc9, 0x93, 0x1e, 0x99, 0x04, 0x61, 0x01, 0xbd, 0x29, 0x0c, 0x3d,
	0x87, 0x35, 0x2d, 0x6b, 0xa5, 0xd3, 0x8c, 0x8f, 0x30, 0x73, 0xd2, 0xbb, 0x49, 0xcf, 0xe6, 0xde,
	0xd9, 0x14, 0x7a, 0x0b, 0x4f, 0x5b, 0x1d, 0xbe, 0xc7, 0x53, 0x4c, 0x08, 0xaf, 0x2b, 0x6d, 0x25,
	0x75, 0x93, 0x4d, 0x0f, 0x9f, 0x3b, 0xf4, 0xc8, 0x81, 0xe1, 0x37, 0x80, 0xc9, 0x03, 0x23, 0x04,
	0x8b, 0x15, 0x1e, 0x51, 0x7f, 0x80, 0x5d, 0xa3, 0x27, 0xb0, 0x2c, 0x69, 0x6e, 0x1c, 0x73, 0x44,
	0x3e, 0x42, 0x2f, 0x60, 0xc3, 0xb6, 0x93, 0x23, 0x4a, 0x99, 0xeb, 0x82, 0x6e, 0x62, 0x9a, 0xcc,
	0xb3, 0x0f, 0x32, 0x14, 0xc2, 0x2a, 0x29, 0x6b, 0xa5, 0xa9, 0x54, 0xc1, 0xe2, 0xf6, 0xc2, 0x5e,
	0x37, 0x19, 0xc7, 0xe1, 0xf7, 0x0e, 0xf4, 0xa6, 0x5a, 0xe2, 0x81, 0x5e, 0x68, 0x17, 0x36, 0xfc,
	0x89, 0xad, 0x7d, 0xee, 0x1e, 0xeb, 0x3e, 0xeb, 0x0c, 0x0c, 0x19, 0xfc, 0xff, 0x4f, 0x3b, 0x3d,
	0x8c, 0xa2, 0x30, 0x83, 0xff, 0xfe, 0xde, 0x81, 0x9e, 0x41, 0xd7, 0xb8, 0xad, 0x04, 0x26, 0xad,
	0xfd, 0x93, 0x04, 0x0a, 0x60, 0xc5, 0xab, 0xf5, 0xe2, 0xdb, 0xd0, 0x20, 0x0d, 0x95, 0xca, 0x88,
	0x73, 0xf6, 0xb7, 0xe1, 0x71, 0x0f, 0xba, 0xe6, 0x37, 0x4f, 0xf5, 0x8d, 0xa0, 0x3b, 0x87, 0x00,
	0x56, 0x9d, 0xc6, 0xba, 0x56, 0x28, 0x86, 0x47, 0x7c, 0x68, 0xda, 0x84, 0x66, 0x69, 0x4e, 0x2b,
	0x2a, 0x27, 0xb7, 0x5b, 0x48, 0x50, 0x0b, 0x9d, 0x8e, 0x91, 0xe3, 0x2f, 0x3f, 0x7f, 0x6f, 0x75,
	0x2e, 0x3e, 0xdc, 0x67, 0xa0, 0x8a, 0xab, 0xfc, 0xf6, 0x50, 0x9d, 0xb6, 0x68, 0x3c, 0x60, 0x87,
	0xcb, 0x76, 0xbe, 0xbd, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0xba, 0xe0, 0xde, 0xf3, 0xa5, 0x05,
	0x00, 0x00,
}

func (this *MeshSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshSpec)
	if !ok {
		that2, ok := that.(MeshSpec)
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
	if that1.MeshType == nil {
		if this.MeshType != nil {
			return false
		}
	} else if this.MeshType == nil {
		return false
	} else if !this.MeshType.Equal(that1.MeshType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshSpec_Istio_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshSpec_Istio_)
	if !ok {
		that2, ok := that.(MeshSpec_Istio_)
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
	if !this.Istio.Equal(that1.Istio) {
		return false
	}
	return true
}
func (this *MeshSpec_AwsAppMesh_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshSpec_AwsAppMesh_)
	if !ok {
		that2, ok := that.(MeshSpec_AwsAppMesh_)
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
	if !this.AwsAppMesh.Equal(that1.AwsAppMesh) {
		return false
	}
	return true
}
func (this *MeshSpec_Linkerd) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshSpec_Linkerd)
	if !ok {
		that2, ok := that.(MeshSpec_Linkerd)
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
	if !this.Linkerd.Equal(that1.Linkerd) {
		return false
	}
	return true
}
func (this *MeshSpec_ConsulConnect) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshSpec_ConsulConnect)
	if !ok {
		that2, ok := that.(MeshSpec_ConsulConnect)
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
	if !this.ConsulConnect.Equal(that1.ConsulConnect) {
		return false
	}
	return true
}
func (this *MeshSpec_Istio) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshSpec_Istio)
	if !ok {
		that2, ok := that.(MeshSpec_Istio)
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
	if !this.Installation.Equal(that1.Installation) {
		return false
	}
	if !this.CitadelInfo.Equal(that1.CitadelInfo) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshSpec_Istio_CitadelInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshSpec_Istio_CitadelInfo)
	if !ok {
		that2, ok := that.(MeshSpec_Istio_CitadelInfo)
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
	if this.TrustDomain != that1.TrustDomain {
		return false
	}
	if this.CitadelServiceAccount != that1.CitadelServiceAccount {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshSpec_AwsAppMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshSpec_AwsAppMesh)
	if !ok {
		that2, ok := that.(MeshSpec_AwsAppMesh)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Region != that1.Region {
		return false
	}
	if this.AwsAccountId != that1.AwsAccountId {
		return false
	}
	if len(this.Clusters) != len(that1.Clusters) {
		return false
	}
	for i := range this.Clusters {
		if this.Clusters[i] != that1.Clusters[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshSpec_LinkerdMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshSpec_LinkerdMesh)
	if !ok {
		that2, ok := that.(MeshSpec_LinkerdMesh)
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
	if !this.Installation.Equal(that1.Installation) {
		return false
	}
	if this.ClusterDomain != that1.ClusterDomain {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshSpec_ConsulConnectMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshSpec_ConsulConnectMesh)
	if !ok {
		that2, ok := that.(MeshSpec_ConsulConnectMesh)
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
	if !this.Installation.Equal(that1.Installation) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshSpec_MeshInstallation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshSpec_MeshInstallation)
	if !ok {
		that2, ok := that.(MeshSpec_MeshInstallation)
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
	if this.Namespace != that1.Namespace {
		return false
	}
	if this.Cluster != that1.Cluster {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshStatus)
	if !ok {
		that2, ok := that.(MeshStatus)
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
