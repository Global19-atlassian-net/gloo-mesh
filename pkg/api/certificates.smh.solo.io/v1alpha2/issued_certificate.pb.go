// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/certificates/issued_certificate.proto

package v1alpha2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
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

// Possible states in which an IssuedCertificate can exist.
type IssuedCertificateStatus_State int32

const (
	// The IssuedCertificate has yet to be picked up by the agent.
	IssuedCertificateStatus_PENDING IssuedCertificateStatus_State = 0
	// The agent has created a local private key
	// and a CertificateRequest for the IssuedCertificate.
	// In this state, the agent is waiting for the Issuer
	// to issue certificates for the CertificateRequest before proceeding.
	IssuedCertificateStatus_REQUESTED IssuedCertificateStatus_State = 1
	// The certificate has been issued. Any pods that require restarting will be restarted at this point.
	IssuedCertificateStatus_ISSUED IssuedCertificateStatus_State = 2
	// The reply from the Issuer has been processed and
	// the agent has placed the final certificate secret
	// in the target location specified by the IssuedCertificate.
	IssuedCertificateStatus_FINISHED IssuedCertificateStatus_State = 3
	// Processing the certificate workflow failed.
	IssuedCertificateStatus_FAILED IssuedCertificateStatus_State = 4
)

var IssuedCertificateStatus_State_name = map[int32]string{
	0: "PENDING",
	1: "REQUESTED",
	2: "ISSUED",
	3: "FINISHED",
	4: "FAILED",
}

var IssuedCertificateStatus_State_value = map[string]int32{
	"PENDING":   0,
	"REQUESTED": 1,
	"ISSUED":    2,
	"FINISHED":  3,
	"FAILED":    4,
}

func (x IssuedCertificateStatus_State) String() string {
	return proto.EnumName(IssuedCertificateStatus_State_name, int32(x))
}

func (IssuedCertificateStatus_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f4866f21673c5e5b, []int{1, 0}
}

//
//IssuedCertificates are used to issue SSL certificates
//to remote Kubernetes clusters from a central (out-of-cluster) Certificate Authority.
//
//When an IssuedCertificate is created, a certificate is issued to a remote cluster by a central Certificate Authority via
//the following workflow:
//- The Certificate Issuer creates the IssuedCertificate resource on the remote cluster
//- The Certificate Signature Requesting Agent installed to the remote cluster generates
//a Certificate Signing Request and writes it to the status of the IssuedCertificate
//- Finally, the Certificate Issuer generates signed a certificate for the CSR and writes
//it back as Secret in the remote cluster.
//
//Shared trust can therefore be established across clusters without requiring
//private keys to ever leave the node.
type IssuedCertificateSpec struct {
	//
	//A list of hostnames and IPs to generate a certificate for.
	//This can also be set to the identity running the workload,
	//e.g. a Kubernetes service account.
	//
	//Generally for an Istio CA this will take the form `spiffe://cluster.local/ns/istio-system/sa/citadel`.
	//
	//"cluster.local" may be replaced by the root of trust domain for the mesh.
	Hosts []string `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// The organization for this certificate.
	Org string `protobuf:"bytes,2,opt,name=org,proto3" json:"org,omitempty"`
	// The secret containing the root SSL certificate used to sign this IssuedCertificate (located in the Certificate Issuer's cluster).
	SigningCertificateSecret *v1.ObjectRef `protobuf:"bytes,3,opt,name=signing_certificate_secret,json=signingCertificateSecret,proto3" json:"signing_certificate_secret,omitempty"`
	// The secret containing the SSL certificate to be generated for this IssuedCertificate (located in the Certificate Agent's cluster).
	IssuedCertificateSecret *v1.ObjectRef `protobuf:"bytes,4,opt,name=issued_certificate_secret,json=issuedCertificateSecret,proto3" json:"issued_certificate_secret,omitempty"`
	// A list of k8s pods to bounce (delete and cause a restart)
	// when the certificate is issued.
	// This will include the control plane pods as well as any pods
	// which share a data plane with the target mesh.
	PodsToBounce         []*IssuedCertificateSpec_PodSelector `protobuf:"bytes,5,rep,name=pods_to_bounce,json=podsToBounce,proto3" json:"pods_to_bounce,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *IssuedCertificateSpec) Reset()         { *m = IssuedCertificateSpec{} }
func (m *IssuedCertificateSpec) String() string { return proto.CompactTextString(m) }
func (*IssuedCertificateSpec) ProtoMessage()    {}
func (*IssuedCertificateSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4866f21673c5e5b, []int{0}
}
func (m *IssuedCertificateSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssuedCertificateSpec.Unmarshal(m, b)
}
func (m *IssuedCertificateSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssuedCertificateSpec.Marshal(b, m, deterministic)
}
func (m *IssuedCertificateSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssuedCertificateSpec.Merge(m, src)
}
func (m *IssuedCertificateSpec) XXX_Size() int {
	return xxx_messageInfo_IssuedCertificateSpec.Size(m)
}
func (m *IssuedCertificateSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_IssuedCertificateSpec.DiscardUnknown(m)
}

var xxx_messageInfo_IssuedCertificateSpec proto.InternalMessageInfo

func (m *IssuedCertificateSpec) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *IssuedCertificateSpec) GetOrg() string {
	if m != nil {
		return m.Org
	}
	return ""
}

func (m *IssuedCertificateSpec) GetSigningCertificateSecret() *v1.ObjectRef {
	if m != nil {
		return m.SigningCertificateSecret
	}
	return nil
}

func (m *IssuedCertificateSpec) GetIssuedCertificateSecret() *v1.ObjectRef {
	if m != nil {
		return m.IssuedCertificateSecret
	}
	return nil
}

func (m *IssuedCertificateSpec) GetPodsToBounce() []*IssuedCertificateSpec_PodSelector {
	if m != nil {
		return m.PodsToBounce
	}
	return nil
}

// Pods that will be restarted.
type IssuedCertificateSpec_PodSelector struct {
	// The namespace in which the pods live.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Any labels shared by the pods.
	Labels               map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *IssuedCertificateSpec_PodSelector) Reset()         { *m = IssuedCertificateSpec_PodSelector{} }
func (m *IssuedCertificateSpec_PodSelector) String() string { return proto.CompactTextString(m) }
func (*IssuedCertificateSpec_PodSelector) ProtoMessage()    {}
func (*IssuedCertificateSpec_PodSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4866f21673c5e5b, []int{0, 0}
}
func (m *IssuedCertificateSpec_PodSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssuedCertificateSpec_PodSelector.Unmarshal(m, b)
}
func (m *IssuedCertificateSpec_PodSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssuedCertificateSpec_PodSelector.Marshal(b, m, deterministic)
}
func (m *IssuedCertificateSpec_PodSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssuedCertificateSpec_PodSelector.Merge(m, src)
}
func (m *IssuedCertificateSpec_PodSelector) XXX_Size() int {
	return xxx_messageInfo_IssuedCertificateSpec_PodSelector.Size(m)
}
func (m *IssuedCertificateSpec_PodSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_IssuedCertificateSpec_PodSelector.DiscardUnknown(m)
}

var xxx_messageInfo_IssuedCertificateSpec_PodSelector proto.InternalMessageInfo

func (m *IssuedCertificateSpec_PodSelector) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *IssuedCertificateSpec_PodSelector) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// The IssuedCertificate status is written by the CertificateRequesting agent.
type IssuedCertificateStatus struct {
	// The most recent generation observed in the the IssuedCertificate metadata.
	// If the observedGeneration does not match generation, the Certificate Requesting Agent has not processed the most
	// recent version of this IssuedCertificate.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// Any error observed which prevented the CertificateRequest from being processed.
	// If the error is empty, the request has been processed successfully.
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// The current state of the IssuedCertificate workflow, reported by the agent.
	State                IssuedCertificateStatus_State `protobuf:"varint,3,opt,name=state,proto3,enum=certificates.smh.solo.io.IssuedCertificateStatus_State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *IssuedCertificateStatus) Reset()         { *m = IssuedCertificateStatus{} }
func (m *IssuedCertificateStatus) String() string { return proto.CompactTextString(m) }
func (*IssuedCertificateStatus) ProtoMessage()    {}
func (*IssuedCertificateStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4866f21673c5e5b, []int{1}
}
func (m *IssuedCertificateStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssuedCertificateStatus.Unmarshal(m, b)
}
func (m *IssuedCertificateStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssuedCertificateStatus.Marshal(b, m, deterministic)
}
func (m *IssuedCertificateStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssuedCertificateStatus.Merge(m, src)
}
func (m *IssuedCertificateStatus) XXX_Size() int {
	return xxx_messageInfo_IssuedCertificateStatus.Size(m)
}
func (m *IssuedCertificateStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_IssuedCertificateStatus.DiscardUnknown(m)
}

var xxx_messageInfo_IssuedCertificateStatus proto.InternalMessageInfo

func (m *IssuedCertificateStatus) GetObservedGeneration() int64 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func (m *IssuedCertificateStatus) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *IssuedCertificateStatus) GetState() IssuedCertificateStatus_State {
	if m != nil {
		return m.State
	}
	return IssuedCertificateStatus_PENDING
}

func init() {
	proto.RegisterEnum("certificates.smh.solo.io.IssuedCertificateStatus_State", IssuedCertificateStatus_State_name, IssuedCertificateStatus_State_value)
	proto.RegisterType((*IssuedCertificateSpec)(nil), "certificates.smh.solo.io.IssuedCertificateSpec")
	proto.RegisterType((*IssuedCertificateSpec_PodSelector)(nil), "certificates.smh.solo.io.IssuedCertificateSpec.PodSelector")
	proto.RegisterMapType((map[string]string)(nil), "certificates.smh.solo.io.IssuedCertificateSpec.PodSelector.LabelsEntry")
	proto.RegisterType((*IssuedCertificateStatus)(nil), "certificates.smh.solo.io.IssuedCertificateStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/certificates/issued_certificate.proto", fileDescriptor_f4866f21673c5e5b)
}

var fileDescriptor_f4866f21673c5e5b = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x26, 0xed, 0x3a, 0xa8, 0x3b, 0xa6, 0xca, 0x0c, 0x2d, 0x94, 0x69, 0x54, 0xbb, 0xea, 0xcd,
	0x1c, 0xad, 0x5c, 0xf0, 0x77, 0xc5, 0x68, 0x36, 0x22, 0x6d, 0x63, 0x24, 0x9b, 0x84, 0x76, 0x53,
	0x39, 0xe9, 0x59, 0x62, 0x9a, 0xc4, 0x91, 0xed, 0x54, 0xda, 0xeb, 0x70, 0xc5, 0x9b, 0xf0, 0x1e,
	0x3c, 0x06, 0x57, 0xc8, 0x4e, 0xca, 0x22, 0x0a, 0xd2, 0x10, 0x57, 0x3e, 0x3f, 0xf6, 0x77, 0xbe,
	0xef, 0x1c, 0xdb, 0xe8, 0x34, 0x66, 0x2a, 0x29, 0x43, 0x12, 0xf1, 0xcc, 0x91, 0x3c, 0xe5, 0xfb,
	0x8c, 0x3b, 0x12, 0xc4, 0x82, 0x45, 0xb0, 0x9f, 0x81, 0x4c, 0xf6, 0x93, 0x32, 0x74, 0x68, 0xc1,
	0x9c, 0x08, 0x84, 0x62, 0xd7, 0x2c, 0xa2, 0x0a, 0xa4, 0xc3, 0xa4, 0x2c, 0x61, 0x36, 0x6d, 0xc4,
	0x48, 0x21, 0xb8, 0xe2, 0xd8, 0x6e, 0x6e, 0x23, 0x32, 0x4b, 0x88, 0x06, 0x25, 0x8c, 0x0f, 0x9e,
	0xca, 0xf9, 0x62, 0x5c, 0x21, 0x71, 0x01, 0xce, 0xe2, 0xc0, 0xac, 0xd5, 0xb1, 0xc1, 0xb3, 0x98,
	0xf3, 0x38, 0x05, 0xc7, 0x78, 0x61, 0x79, 0xed, 0x28, 0x96, 0x81, 0x54, 0x34, 0x2b, 0xea, 0x0d,
	0xbb, 0xbf, 0x6f, 0x98, 0x95, 0x82, 0x2a, 0xc6, 0xf3, 0x3a, 0xbf, 0x15, 0xf3, 0x98, 0x1b, 0xd3,
	0xd1, 0x56, 0x15, 0xdd, 0xfb, 0xb2, 0x86, 0x1e, 0x7b, 0x86, 0xea, 0xbb, 0x5b, 0x5a, 0x41, 0x01,
	0x11, 0xde, 0x42, 0x9d, 0x84, 0x4b, 0x25, 0x6d, 0x6b, 0xd8, 0x1e, 0x75, 0xfd, 0xca, 0xc1, 0x7d,
	0xd4, 0xe6, 0x22, 0xb6, 0x5b, 0x43, 0x6b, 0xd4, 0xf5, 0xb5, 0x89, 0xaf, 0xd0, 0x40, 0xb2, 0x38,
	0x67, 0x79, 0xdc, 0x14, 0x3b, 0x95, 0x10, 0x09, 0x50, 0x76, 0x7b, 0x68, 0x8d, 0x7a, 0xe3, 0x1d,
	0x62, 0x94, 0x68, 0x7d, 0x4b, 0xb5, 0xe4, 0x43, 0xf8, 0x19, 0x22, 0xe5, 0xc3, 0xb5, 0x6f, 0xd7,
	0xe7, 0x9b, 0x0c, 0xcc, 0x69, 0xfc, 0x09, 0x3d, 0x59, 0xed, 0xe3, 0x12, 0x7a, 0xed, 0x0e, 0xd0,
	0xdb, 0x6c, 0x45, 0x5b, 0x85, 0x4c, 0xd1, 0x66, 0xc1, 0x67, 0x72, 0xaa, 0xf8, 0x34, 0xe4, 0x65,
	0x1e, 0x81, 0xdd, 0x19, 0xb6, 0x47, 0xbd, 0xf1, 0x1b, 0xf2, 0xb7, 0xf1, 0x90, 0x3f, 0xb6, 0x89,
	0x9c, 0xf3, 0x59, 0x00, 0x29, 0x44, 0x8a, 0x0b, 0x7f, 0x43, 0x43, 0x5e, 0xf0, 0x43, 0x03, 0x38,
	0xf8, 0x66, 0xa1, 0x5e, 0x23, 0x8b, 0x77, 0x50, 0x37, 0xa7, 0x19, 0xc8, 0x82, 0x46, 0x60, 0x5b,
	0xa6, 0x81, 0xb7, 0x01, 0x3c, 0x45, 0xeb, 0x29, 0x0d, 0x21, 0x95, 0x76, 0xcb, 0x10, 0x39, 0xfe,
	0x0f, 0x22, 0xe4, 0xc4, 0x20, 0xb9, 0xb9, 0x12, 0x37, 0x7e, 0x0d, 0x3b, 0x78, 0x85, 0x7a, 0x8d,
	0xb0, 0x1e, 0xe4, 0x1c, 0x6e, 0x6a, 0x1e, 0xda, 0xd4, 0x03, 0x5f, 0xd0, 0xb4, 0x84, 0x7a, 0xb8,
	0x95, 0xf3, 0xba, 0xf5, 0xd2, 0xda, 0xfb, 0x61, 0xa1, 0xed, 0xd5, 0xa2, 0x8a, 0xaa, 0x52, 0x62,
	0x07, 0x3d, 0xe2, 0xa1, 0x7e, 0x0e, 0x30, 0x9b, 0xc6, 0x90, 0x43, 0x75, 0xe7, 0x0c, 0x6e, 0xdb,
	0xc7, 0xcb, 0xd4, 0xf1, 0xaf, 0x8c, 0x2e, 0x03, 0x42, 0x70, 0xb1, 0x2c, 0x63, 0x1c, 0x7c, 0x8a,
	0x3a, 0x52, 0x51, 0x05, 0xe6, 0xc2, 0x6c, 0x8e, 0x5f, 0xfc, 0x8b, 0x7a, 0x43, 0x84, 0xe8, 0x05,
	0xfc, 0x0a, 0x65, 0xcf, 0x43, 0x1d, 0xe3, 0xe3, 0x1e, 0xba, 0x7f, 0xee, 0x9e, 0x4d, 0xbc, 0xb3,
	0xe3, 0xfe, 0x3d, 0xfc, 0x10, 0x75, 0x7d, 0xf7, 0xe3, 0xa5, 0x1b, 0x5c, 0xb8, 0x93, 0xbe, 0x85,
	0x11, 0x5a, 0xf7, 0x82, 0xe0, 0xd2, 0x9d, 0xf4, 0x5b, 0x78, 0x03, 0x3d, 0x38, 0xf2, 0xce, 0xbc,
	0xe0, 0xbd, 0x3b, 0xe9, 0xb7, 0x75, 0xe6, 0xe8, 0xad, 0x77, 0xe2, 0x4e, 0xfa, 0x6b, 0x87, 0xc1,
	0xd7, 0xef, 0xbb, 0xd6, 0xd5, 0x9d, 0x3e, 0x81, 0x62, 0x1e, 0xaf, 0x7c, 0x04, 0x4d, 0xee, 0xce,
	0xe2, 0x80, 0xa6, 0x45, 0x42, 0xc7, 0xe1, 0xba, 0x79, 0x7d, 0xcf, 0x7f, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x6f, 0xcb, 0xda, 0x27, 0x5c, 0x04, 0x00, 0x00,
}

func (this *IssuedCertificateSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IssuedCertificateSpec)
	if !ok {
		that2, ok := that.(IssuedCertificateSpec)
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
	if len(this.Hosts) != len(that1.Hosts) {
		return false
	}
	for i := range this.Hosts {
		if this.Hosts[i] != that1.Hosts[i] {
			return false
		}
	}
	if this.Org != that1.Org {
		return false
	}
	if !this.SigningCertificateSecret.Equal(that1.SigningCertificateSecret) {
		return false
	}
	if !this.IssuedCertificateSecret.Equal(that1.IssuedCertificateSecret) {
		return false
	}
	if len(this.PodsToBounce) != len(that1.PodsToBounce) {
		return false
	}
	for i := range this.PodsToBounce {
		if !this.PodsToBounce[i].Equal(that1.PodsToBounce[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *IssuedCertificateSpec_PodSelector) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IssuedCertificateSpec_PodSelector)
	if !ok {
		that2, ok := that.(IssuedCertificateSpec_PodSelector)
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
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if this.Labels[i] != that1.Labels[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *IssuedCertificateStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IssuedCertificateStatus)
	if !ok {
		that2, ok := that.(IssuedCertificateStatus)
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
	if this.Error != that1.Error {
		return false
	}
	if this.State != that1.State {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}