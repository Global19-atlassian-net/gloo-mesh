// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/discovery/v1alpha1/mesh_service.proto

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1/types"
	_ "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1"
	_ "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for MeshServiceSpec
func (this *MeshServiceSpec) MarshalJSON() ([]byte, error) {
	str, err := MeshServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshServiceSpec
func (this *MeshServiceSpec) UnmarshalJSON(b []byte) error {
	return MeshServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshServiceSpec_Federation
func (this *MeshServiceSpec_Federation) MarshalJSON() ([]byte, error) {
	str, err := MeshServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshServiceSpec_Federation
func (this *MeshServiceSpec_Federation) UnmarshalJSON(b []byte) error {
	return MeshServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshServiceSpec_KubeService
func (this *MeshServiceSpec_KubeService) MarshalJSON() ([]byte, error) {
	str, err := MeshServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshServiceSpec_KubeService
func (this *MeshServiceSpec_KubeService) UnmarshalJSON(b []byte) error {
	return MeshServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshServiceSpec_KubeService_KubeServicePort
func (this *MeshServiceSpec_KubeService_KubeServicePort) MarshalJSON() ([]byte, error) {
	str, err := MeshServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshServiceSpec_KubeService_KubeServicePort
func (this *MeshServiceSpec_KubeService_KubeServicePort) UnmarshalJSON(b []byte) error {
	return MeshServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshServiceSpec_Subset
func (this *MeshServiceSpec_Subset) MarshalJSON() ([]byte, error) {
	str, err := MeshServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshServiceSpec_Subset
func (this *MeshServiceSpec_Subset) UnmarshalJSON(b []byte) error {
	return MeshServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshServiceStatus
func (this *MeshServiceStatus) MarshalJSON() ([]byte, error) {
	str, err := MeshServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshServiceStatus
func (this *MeshServiceStatus) UnmarshalJSON(b []byte) error {
	return MeshServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshServiceStatus_AppliedTrafficPolicy
func (this *MeshServiceStatus_AppliedTrafficPolicy) MarshalJSON() ([]byte, error) {
	str, err := MeshServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshServiceStatus_AppliedTrafficPolicy
func (this *MeshServiceStatus_AppliedTrafficPolicy) UnmarshalJSON(b []byte) error {
	return MeshServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshServiceStatus_AppliedAccessPolicy
func (this *MeshServiceStatus_AppliedAccessPolicy) MarshalJSON() ([]byte, error) {
	str, err := MeshServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshServiceStatus_AppliedAccessPolicy
func (this *MeshServiceStatus_AppliedAccessPolicy) UnmarshalJSON(b []byte) error {
	return MeshServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	MeshServiceMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	MeshServiceUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
