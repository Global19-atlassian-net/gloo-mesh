// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/discovery/v1alpha1/mesh.proto

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1"
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
	CitadelInfo *MeshSpec_Istio_CitadelInfo `protobuf:"bytes,2,opt,name=citadel_info,json=citadelInfo,proto3" json:"citadel_info,omitempty"`
	// configuration for Istio IngressGateway, the Istio Ingress
	IngressGateways      []*MeshSpec_Istio_IngressGatewayInfo `protobuf:"bytes,3,rep,name=ingress_gateways,json=ingressGateways,proto3" json:"ingress_gateways,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
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

func (m *MeshSpec_Istio) GetIngressGateways() []*MeshSpec_Istio_IngressGatewayInfo {
	if m != nil {
		return m.IngressGateways
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

type MeshSpec_Istio_IngressGatewayInfo struct {
	// labels matching the workload which backs the gateway,
	// defaults to {"istio": "ingressgateway"}
	WorkloadLabels map[string]string `protobuf:"bytes,1,rep,name=workload_labels,json=workloadLabels,proto3" json:"workload_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// the externally-reachable address
	// this will be the addressed used for cross-cluster connectivity.
	// defaults to the LoadBalancer Address or NodeIP of the Kubernetes Service (depending on its type)
	ExternalAddress string `protobuf:"bytes,2,opt,name=external_address,json=externalAddress,proto3" json:"external_address,omitempty"`
	// service port on which the gateway is listening for TLS connections.
	// List of common ports: https://istio.io/latest/docs/ops/deployment/requirements/#ports-used-by-istio
	// Defaults to 15443
	TlsServicePort uint32 `protobuf:"varint,3,opt,name=tls_service_port,json=tlsServicePort,proto3" json:"tls_service_port,omitempty"`
	// container port on which the gateway is listening for TLS connections.
	// Defaults to 15443
	TlsContainerPort     uint32   `protobuf:"varint,4,opt,name=tls_container_port,json=tlsContainerPort,proto3" json:"tls_container_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshSpec_Istio_IngressGatewayInfo) Reset()         { *m = MeshSpec_Istio_IngressGatewayInfo{} }
func (m *MeshSpec_Istio_IngressGatewayInfo) String() string { return proto.CompactTextString(m) }
func (*MeshSpec_Istio_IngressGatewayInfo) ProtoMessage()    {}
func (*MeshSpec_Istio_IngressGatewayInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_345fdec2269fec48, []int{0, 0, 1}
}
func (m *MeshSpec_Istio_IngressGatewayInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshSpec_Istio_IngressGatewayInfo.Unmarshal(m, b)
}
func (m *MeshSpec_Istio_IngressGatewayInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshSpec_Istio_IngressGatewayInfo.Marshal(b, m, deterministic)
}
func (m *MeshSpec_Istio_IngressGatewayInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshSpec_Istio_IngressGatewayInfo.Merge(m, src)
}
func (m *MeshSpec_Istio_IngressGatewayInfo) XXX_Size() int {
	return xxx_messageInfo_MeshSpec_Istio_IngressGatewayInfo.Size(m)
}
func (m *MeshSpec_Istio_IngressGatewayInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshSpec_Istio_IngressGatewayInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MeshSpec_Istio_IngressGatewayInfo proto.InternalMessageInfo

func (m *MeshSpec_Istio_IngressGatewayInfo) GetWorkloadLabels() map[string]string {
	if m != nil {
		return m.WorkloadLabels
	}
	return nil
}

func (m *MeshSpec_Istio_IngressGatewayInfo) GetExternalAddress() string {
	if m != nil {
		return m.ExternalAddress
	}
	return ""
}

func (m *MeshSpec_Istio_IngressGatewayInfo) GetTlsServicePort() uint32 {
	if m != nil {
		return m.TlsServicePort
	}
	return 0
}

func (m *MeshSpec_Istio_IngressGatewayInfo) GetTlsContainerPort() uint32 {
	if m != nil {
		return m.TlsContainerPort
	}
	return 0
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
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// The VirtualMesh, if any, which contains this mesh.
	AppliedVirtualMeshes []*MeshStatus_AppliedVirtualMesh `protobuf:"bytes,5,rep,name=applied_virtual_meshes,json=appliedVirtualMeshes,proto3" json:"applied_virtual_meshes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
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

func (m *MeshStatus) GetAppliedVirtualMeshes() []*MeshStatus_AppliedVirtualMesh {
	if m != nil {
		return m.AppliedVirtualMeshes
	}
	return nil
}

// AppliedTrafficPolicy represents a traffic policy that has been applied to the MeshService.
// if an existing Traffic Policy becomes invalid, the last applied policy will be used
type MeshStatus_AppliedVirtualMesh struct {
	// reference to the traffic policy
	Ref *v1.ObjectRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	// the observed generation of the accepted traffic policy
	ObservedGeneration int64 `protobuf:"varint,2,opt,name=observedGeneration,proto3" json:"observedGeneration,omitempty"`
	// the last known valid spec of the traffic policy
	Spec                 *v1alpha1.VirtualMeshSpec `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *MeshStatus_AppliedVirtualMesh) Reset()         { *m = MeshStatus_AppliedVirtualMesh{} }
func (m *MeshStatus_AppliedVirtualMesh) String() string { return proto.CompactTextString(m) }
func (*MeshStatus_AppliedVirtualMesh) ProtoMessage()    {}
func (*MeshStatus_AppliedVirtualMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_345fdec2269fec48, []int{1, 0}
}
func (m *MeshStatus_AppliedVirtualMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshStatus_AppliedVirtualMesh.Unmarshal(m, b)
}
func (m *MeshStatus_AppliedVirtualMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshStatus_AppliedVirtualMesh.Marshal(b, m, deterministic)
}
func (m *MeshStatus_AppliedVirtualMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshStatus_AppliedVirtualMesh.Merge(m, src)
}
func (m *MeshStatus_AppliedVirtualMesh) XXX_Size() int {
	return xxx_messageInfo_MeshStatus_AppliedVirtualMesh.Size(m)
}
func (m *MeshStatus_AppliedVirtualMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshStatus_AppliedVirtualMesh.DiscardUnknown(m)
}

var xxx_messageInfo_MeshStatus_AppliedVirtualMesh proto.InternalMessageInfo

func (m *MeshStatus_AppliedVirtualMesh) GetRef() *v1.ObjectRef {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *MeshStatus_AppliedVirtualMesh) GetObservedGeneration() int64 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func (m *MeshStatus_AppliedVirtualMesh) GetSpec() *v1alpha1.VirtualMeshSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func init() {
	proto.RegisterType((*MeshSpec)(nil), "discovery.smh.solo.io.MeshSpec")
	proto.RegisterType((*MeshSpec_Istio)(nil), "discovery.smh.solo.io.MeshSpec.Istio")
	proto.RegisterType((*MeshSpec_Istio_CitadelInfo)(nil), "discovery.smh.solo.io.MeshSpec.Istio.CitadelInfo")
	proto.RegisterType((*MeshSpec_Istio_IngressGatewayInfo)(nil), "discovery.smh.solo.io.MeshSpec.Istio.IngressGatewayInfo")
	proto.RegisterMapType((map[string]string)(nil), "discovery.smh.solo.io.MeshSpec.Istio.IngressGatewayInfo.WorkloadLabelsEntry")
	proto.RegisterType((*MeshSpec_AwsAppMesh)(nil), "discovery.smh.solo.io.MeshSpec.AwsAppMesh")
	proto.RegisterType((*MeshSpec_LinkerdMesh)(nil), "discovery.smh.solo.io.MeshSpec.LinkerdMesh")
	proto.RegisterType((*MeshSpec_ConsulConnectMesh)(nil), "discovery.smh.solo.io.MeshSpec.ConsulConnectMesh")
	proto.RegisterType((*MeshSpec_MeshInstallation)(nil), "discovery.smh.solo.io.MeshSpec.MeshInstallation")
	proto.RegisterType((*MeshStatus)(nil), "discovery.smh.solo.io.MeshStatus")
	proto.RegisterType((*MeshStatus_AppliedVirtualMesh)(nil), "discovery.smh.solo.io.MeshStatus.AppliedVirtualMesh")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/discovery/v1alpha1/mesh.proto", fileDescriptor_345fdec2269fec48)
}

var fileDescriptor_345fdec2269fec48 = []byte{
	// 913 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xae, 0xe3, 0xb8, 0xad, 0x8f, 0x13, 0xc7, 0x4c, 0xd3, 0x62, 0xad, 0xaa, 0x2a, 0x54, 0x54,
	0x84, 0x9f, 0xec, 0xe2, 0x80, 0x50, 0x45, 0xc5, 0x85, 0x1b, 0xaa, 0xd4, 0x28, 0x40, 0xb5, 0x45,
	0x20, 0xf5, 0x66, 0x35, 0x9e, 0x1d, 0xaf, 0xa7, 0x5e, 0xcf, 0xac, 0x66, 0x66, 0x6d, 0xcc, 0x23,
	0xf0, 0x24, 0xdc, 0xf2, 0x28, 0xdc, 0x70, 0xc7, 0x0d, 0x0f, 0x82, 0xd0, 0xfc, 0xac, 0x7f, 0x88,
	0x45, 0xa2, 0x4a, 0xbd, 0xf2, 0x9c, 0xbf, 0xef, 0x7c, 0xe7, 0xcc, 0x9c, 0x3d, 0x86, 0x67, 0x19,
	0xd3, 0xe3, 0x72, 0x18, 0x12, 0x31, 0x8d, 0x94, 0xc8, 0xc5, 0x09, 0x13, 0x91, 0xa2, 0x72, 0xc6,
	0x08, 0x3d, 0x99, 0x52, 0x35, 0x3e, 0x19, 0x97, 0xc3, 0x08, 0x17, 0x2c, 0x4a, 0x99, 0x22, 0x62,
	0x46, 0xe5, 0x22, 0x9a, 0xf5, 0x70, 0x5e, 0x8c, 0x71, 0x2f, 0x32, 0xf6, 0xb0, 0x90, 0x42, 0x0b,
	0x74, 0x77, 0x69, 0x0e, 0xd5, 0x74, 0x1c, 0x1a, 0xa4, 0x90, 0x89, 0x20, 0xdc, 0x86, 0x3e, 0x99,
	0x9d, 0x5a, 0x44, 0x22, 0x24, 0x8d, 0x66, 0x3d, 0xfb, 0xeb, 0x60, 0x82, 0x27, 0x5b, 0x53, 0x73,
	0xaa, 0xe7, 0x42, 0x4e, 0x18, 0xcf, 0x56, 0xb9, 0x67, 0x4c, 0xea, 0x12, 0xe7, 0xc9, 0x8a, 0x43,
	0xf0, 0x20, 0x13, 0x22, 0xcb, 0x69, 0x64, 0xa5, 0x61, 0x39, 0x8a, 0xe6, 0x12, 0x17, 0x05, 0x95,
	0xca, 0xdb, 0x0f, 0x33, 0x91, 0x09, 0x7b, 0x8c, 0xcc, 0xc9, 0x69, 0x1f, 0xfe, 0xd3, 0x82, 0xdb,
	0xdf, 0x52, 0x35, 0x7e, 0x59, 0x50, 0x82, 0xbe, 0x82, 0x06, 0x53, 0x9a, 0x89, 0x6e, 0xed, 0xa8,
	0x76, 0xdc, 0x3a, 0x7d, 0x14, 0x6e, 0x2d, 0x2b, 0xac, 0xfc, 0xc3, 0x81, 0x71, 0x7e, 0x7e, 0x23,
	0x76, 0x51, 0xe8, 0x3b, 0xd8, 0xc3, 0x73, 0x95, 0xe0, 0xa2, 0xb0, 0xbc, 0xba, 0x3b, 0x16, 0xe5,
	0xa3, 0xab, 0x50, 0xfa, 0x73, 0xd5, 0x2f, 0x0a, 0x23, 0x3e, 0xbf, 0x11, 0x03, 0x5e, 0x4a, 0xe8,
	0x1c, 0x6e, 0xe5, 0x8c, 0x4f, 0xa8, 0x4c, 0xbb, 0x75, 0x0b, 0xf5, 0xf1, 0x55, 0x50, 0x17, 0xce,
	0xdd, 0x63, 0x55, 0xd1, 0xe8, 0x15, 0xb4, 0x89, 0xe0, 0xaa, 0xcc, 0x13, 0x22, 0x38, 0xa7, 0x44,
	0x77, 0x77, 0x2d, 0x5e, 0xef, 0x2a, 0xbc, 0x33, 0x1b, 0x75, 0xe6, 0x82, 0x3c, 0xea, 0x3e, 0x59,
	0x57, 0x06, 0x7f, 0x36, 0xa0, 0x61, 0xfb, 0x80, 0x7e, 0x80, 0x3d, 0xc6, 0x95, 0xc6, 0x79, 0x8e,
	0x35, 0x13, 0xdc, 0x37, 0xf1, 0xd3, 0xab, 0x72, 0x98, 0xc3, 0x60, 0x2d, 0x2e, 0xde, 0x40, 0x31,
	0xa8, 0x84, 0x69, 0x9c, 0xd2, 0x3c, 0x61, 0x7c, 0x24, 0x7c, 0x53, 0x7b, 0xd7, 0xba, 0x9a, 0xf0,
	0xcc, 0x45, 0x0e, 0xf8, 0x48, 0xc4, 0x2d, 0xb2, 0x12, 0x10, 0x81, 0x0e, 0xe3, 0x99, 0xa4, 0x4a,
	0x25, 0x19, 0xd6, 0x74, 0x8e, 0x17, 0xaa, 0x5b, 0x3f, 0xaa, 0x1f, 0xb7, 0x4e, 0x1f, 0x5f, 0x0f,
	0x79, 0xe0, 0xa2, 0xcf, 0x5d, 0xb0, 0x4d, 0x70, 0xc0, 0x36, 0x74, 0x2a, 0x18, 0x43, 0x6b, 0x8d,
	0x00, 0x7a, 0x0f, 0xf6, 0xb4, 0x2c, 0x95, 0x4e, 0x52, 0x31, 0xc5, 0xcc, 0xf5, 0xa7, 0x19, 0xb7,
	0xac, 0xee, 0x6b, 0xab, 0x42, 0x5f, 0xc0, 0xbb, 0x55, 0xb1, 0x7e, 0x14, 0x12, 0x4c, 0x88, 0x28,
	0xb9, 0xb6, 0x75, 0x37, 0xe3, 0xbb, 0xde, 0xfc, 0xd2, 0x59, 0xfb, 0xce, 0x18, 0xfc, 0xb1, 0x03,
	0xe8, 0x32, 0x23, 0x54, 0xc2, 0x81, 0x99, 0x9c, 0x5c, 0xe0, 0x34, 0xc9, 0xf1, 0x90, 0xe6, 0xaa,
	0x5b, 0xb3, 0x45, 0x5e, 0xbc, 0x69, 0x91, 0xe1, 0x4f, 0x1e, 0xef, 0xc2, 0xc2, 0x3d, 0xe3, 0x5a,
	0x2e, 0xe2, 0xf6, 0x7c, 0x43, 0x89, 0x3e, 0x84, 0x0e, 0xfd, 0x59, 0x53, 0xc9, 0x71, 0x9e, 0xe0,
	0x34, 0x35, 0x10, 0x9e, 0xfe, 0x41, 0xa5, 0xef, 0x3b, 0x35, 0x3a, 0x86, 0x8e, 0xce, 0xd5, 0xb2,
	0xd8, 0x42, 0x48, 0x6d, 0xdf, 0xfa, 0x7e, 0xdc, 0xd6, 0xb9, 0xf2, 0x55, 0xbe, 0x10, 0x52, 0xa3,
	0x4f, 0x00, 0x19, 0x4f, 0x22, 0xb8, 0xc6, 0x8c, 0x53, 0xe9, 0x7c, 0x77, 0xad, 0xaf, 0xc1, 0x38,
	0xab, 0x0c, 0xc6, 0x3b, 0xe8, 0xc3, 0x9d, 0x2d, 0x4c, 0x51, 0x07, 0xea, 0x13, 0xba, 0xf0, 0x9d,
	0x37, 0x47, 0x74, 0x08, 0x8d, 0x19, 0xce, 0x4b, 0xea, 0x09, 0x3a, 0xe1, 0xcb, 0x9d, 0xc7, 0xb5,
	0xe0, 0x17, 0x80, 0xd5, 0x64, 0x22, 0x04, 0xbb, 0x1c, 0x4f, 0xa9, 0x0f, 0xb5, 0x67, 0x74, 0x0f,
	0x6e, 0x4a, 0x9a, 0x99, 0xa7, 0xee, 0x82, 0xbd, 0x84, 0xde, 0x87, 0xb6, 0xfd, 0x0e, 0xb8, 0xcb,
	0x49, 0x98, 0x1b, 0xdf, 0x66, 0x6c, 0xbe, 0x0e, 0xfe, 0xc6, 0x06, 0x29, 0x0a, 0xe0, 0x36, 0xc9,
	0x4b, 0xa5, 0xa9, 0x54, 0xdd, 0xdd, 0xa3, 0xfa, 0x71, 0x33, 0x5e, 0xca, 0xc1, 0xaf, 0x35, 0x68,
	0xad, 0xcd, 0xf2, 0x5b, 0x1a, 0xad, 0x47, 0xd0, 0xf6, 0x19, 0xab, 0x27, 0xe9, 0xea, 0xd8, 0xf7,
	0x5a, 0xf7, 0x28, 0x03, 0x06, 0xef, 0x5c, 0xfa, 0x0e, 0xbc, 0x1d, 0x46, 0x41, 0x0a, 0x9d, 0xff,
	0x7a, 0xa0, 0xfb, 0xd0, 0x34, 0xdd, 0x56, 0x05, 0x26, 0x55, 0xfb, 0x57, 0x0a, 0xd4, 0x85, 0x5b,
	0x9e, 0xad, 0x27, 0x5f, 0x89, 0xc6, 0x32, 0xa3, 0x52, 0x19, 0x72, 0xae, 0xfd, 0x95, 0xf8, 0xb4,
	0x05, 0x4d, 0xf3, 0x7d, 0x4e, 0xf4, 0xa2, 0xa0, 0x0f, 0xff, 0xda, 0x01, 0xb0, 0xf4, 0x34, 0xd6,
	0xa5, 0x42, 0x11, 0xdc, 0x11, 0x43, 0xf3, 0x1c, 0x69, 0x9a, 0x64, 0x94, 0x53, 0xb9, 0x2a, 0xaf,
	0x1e, 0xa3, 0xca, 0x74, 0xbe, 0xb4, 0xa0, 0xd7, 0x70, 0x0f, 0x17, 0x45, 0xce, 0x68, 0x9a, 0xac,
	0x2f, 0x25, 0xaa, 0xba, 0x0d, 0x3b, 0x6a, 0x9f, 0xff, 0x5f, 0x4b, 0x6c, 0xce, 0xb0, 0xef, 0xe2,
	0x7f, 0x74, 0xe1, 0xc6, 0x10, 0x1f, 0xe2, 0x4b, 0x3a, 0xaa, 0x82, 0xdf, 0x6b, 0x80, 0x2e, 0x3b,
	0xa3, 0x10, 0xea, 0x92, 0x8e, 0xfc, 0x15, 0xdc, 0x0f, 0xed, 0x42, 0x35, 0x6b, 0x76, 0x99, 0xeb,
	0xfb, 0xe1, 0x6b, 0x4a, 0x74, 0x4c, 0x47, 0xb1, 0x71, 0x44, 0x21, 0x6c, 0x29, 0xc4, 0xb6, 0x6f,
	0x7b, 0x89, 0x4f, 0x60, 0x57, 0x15, 0x94, 0xf8, 0x25, 0xf4, 0x41, 0xb8, 0x5a, 0xc8, 0x1b, 0x15,
	0xad, 0x51, 0x32, 0x77, 0x1d, 0xdb, 0xa0, 0xa7, 0x2f, 0x7e, 0xfb, 0xfb, 0x41, 0xed, 0xd5, 0x37,
	0xd7, 0xf9, 0x9f, 0x51, 0x4c, 0xb2, 0xcd, 0xff, 0x1a, 0xeb, 0xf0, 0xcb, 0xdd, 0x3f, 0xbc, 0x69,
	0x37, 0xf7, 0x67, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x36, 0x1e, 0xe5, 0xc9, 0xbc, 0x08, 0x00,
	0x00,
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
	if len(this.IngressGateways) != len(that1.IngressGateways) {
		return false
	}
	for i := range this.IngressGateways {
		if !this.IngressGateways[i].Equal(that1.IngressGateways[i]) {
			return false
		}
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
func (this *MeshSpec_Istio_IngressGatewayInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshSpec_Istio_IngressGatewayInfo)
	if !ok {
		that2, ok := that.(MeshSpec_Istio_IngressGatewayInfo)
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
	if len(this.WorkloadLabels) != len(that1.WorkloadLabels) {
		return false
	}
	for i := range this.WorkloadLabels {
		if this.WorkloadLabels[i] != that1.WorkloadLabels[i] {
			return false
		}
	}
	if this.ExternalAddress != that1.ExternalAddress {
		return false
	}
	if this.TlsServicePort != that1.TlsServicePort {
		return false
	}
	if this.TlsContainerPort != that1.TlsContainerPort {
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
	if len(this.AppliedVirtualMeshes) != len(that1.AppliedVirtualMeshes) {
		return false
	}
	for i := range this.AppliedVirtualMeshes {
		if !this.AppliedVirtualMeshes[i].Equal(that1.AppliedVirtualMeshes[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshStatus_AppliedVirtualMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshStatus_AppliedVirtualMesh)
	if !ok {
		that2, ok := that.(MeshStatus_AppliedVirtualMesh)
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
	if !this.Ref.Equal(that1.Ref) {
		return false
	}
	if this.ObservedGeneration != that1.ObservedGeneration {
		return false
	}
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
