// Code generated by MockGen. DO NOT EDIT.
// Source: ./snapshot.go

// Package mock_input is a generated GoMock package.
package mock_input

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta2sets "github.com/solo-io/external-apis/pkg/api/appmesh/appmesh.k8s.aws/v1beta2/sets"
	v1sets "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1/sets"
	v1sets0 "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	input "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/input"
	v1alpha2sets "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1alpha2/sets"
	multicluster "github.com/solo-io/skv2/pkg/multicluster"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockSnapshot is a mock of Snapshot interface
type MockSnapshot struct {
	ctrl     *gomock.Controller
	recorder *MockSnapshotMockRecorder
}

// MockSnapshotMockRecorder is the mock recorder for MockSnapshot
type MockSnapshotMockRecorder struct {
	mock *MockSnapshot
}

// NewMockSnapshot creates a new mock instance
func NewMockSnapshot(ctrl *gomock.Controller) *MockSnapshot {
	mock := &MockSnapshot{ctrl: ctrl}
	mock.recorder = &MockSnapshotMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSnapshot) EXPECT() *MockSnapshotMockRecorder {
	return m.recorder
}

// Settings mocks base method
func (m *MockSnapshot) Settings() v1alpha2sets.SettingsSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Settings")
	ret0, _ := ret[0].(v1alpha2sets.SettingsSet)
	return ret0
}

// Settings indicates an expected call of Settings
func (mr *MockSnapshotMockRecorder) Settings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Settings", reflect.TypeOf((*MockSnapshot)(nil).Settings))
}

// Meshes mocks base method
func (m *MockSnapshot) Meshes() v1beta2sets.MeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Meshes")
	ret0, _ := ret[0].(v1beta2sets.MeshSet)
	return ret0
}

// Meshes indicates an expected call of Meshes
func (mr *MockSnapshotMockRecorder) Meshes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Meshes", reflect.TypeOf((*MockSnapshot)(nil).Meshes))
}

// ConfigMaps mocks base method
func (m *MockSnapshot) ConfigMaps() v1sets0.ConfigMapSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigMaps")
	ret0, _ := ret[0].(v1sets0.ConfigMapSet)
	return ret0
}

// ConfigMaps indicates an expected call of ConfigMaps
func (mr *MockSnapshotMockRecorder) ConfigMaps() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigMaps", reflect.TypeOf((*MockSnapshot)(nil).ConfigMaps))
}

// Services mocks base method
func (m *MockSnapshot) Services() v1sets0.ServiceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Services")
	ret0, _ := ret[0].(v1sets0.ServiceSet)
	return ret0
}

// Services indicates an expected call of Services
func (mr *MockSnapshotMockRecorder) Services() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Services", reflect.TypeOf((*MockSnapshot)(nil).Services))
}

// Pods mocks base method
func (m *MockSnapshot) Pods() v1sets0.PodSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pods")
	ret0, _ := ret[0].(v1sets0.PodSet)
	return ret0
}

// Pods indicates an expected call of Pods
func (mr *MockSnapshotMockRecorder) Pods() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pods", reflect.TypeOf((*MockSnapshot)(nil).Pods))
}

// Nodes mocks base method
func (m *MockSnapshot) Nodes() v1sets0.NodeSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Nodes")
	ret0, _ := ret[0].(v1sets0.NodeSet)
	return ret0
}

// Nodes indicates an expected call of Nodes
func (mr *MockSnapshotMockRecorder) Nodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nodes", reflect.TypeOf((*MockSnapshot)(nil).Nodes))
}

// Deployments mocks base method
func (m *MockSnapshot) Deployments() v1sets.DeploymentSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deployments")
	ret0, _ := ret[0].(v1sets.DeploymentSet)
	return ret0
}

// Deployments indicates an expected call of Deployments
func (mr *MockSnapshotMockRecorder) Deployments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deployments", reflect.TypeOf((*MockSnapshot)(nil).Deployments))
}

// ReplicaSets mocks base method
func (m *MockSnapshot) ReplicaSets() v1sets.ReplicaSetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplicaSets")
	ret0, _ := ret[0].(v1sets.ReplicaSetSet)
	return ret0
}

// ReplicaSets indicates an expected call of ReplicaSets
func (mr *MockSnapshotMockRecorder) ReplicaSets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicaSets", reflect.TypeOf((*MockSnapshot)(nil).ReplicaSets))
}

// DaemonSets mocks base method
func (m *MockSnapshot) DaemonSets() v1sets.DaemonSetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DaemonSets")
	ret0, _ := ret[0].(v1sets.DaemonSetSet)
	return ret0
}

// DaemonSets indicates an expected call of DaemonSets
func (mr *MockSnapshotMockRecorder) DaemonSets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DaemonSets", reflect.TypeOf((*MockSnapshot)(nil).DaemonSets))
}

// StatefulSets mocks base method
func (m *MockSnapshot) StatefulSets() v1sets.StatefulSetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatefulSets")
	ret0, _ := ret[0].(v1sets.StatefulSetSet)
	return ret0
}

// StatefulSets indicates an expected call of StatefulSets
func (mr *MockSnapshotMockRecorder) StatefulSets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatefulSets", reflect.TypeOf((*MockSnapshot)(nil).StatefulSets))
}

// SyncStatuses mocks base method
func (m *MockSnapshot) SyncStatuses(ctx context.Context, c client.Client) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStatuses", ctx, c)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncStatuses indicates an expected call of SyncStatuses
func (mr *MockSnapshotMockRecorder) SyncStatuses(ctx, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStatuses", reflect.TypeOf((*MockSnapshot)(nil).SyncStatuses), ctx, c)
}

// SyncStatusesMultiCluster mocks base method
func (m *MockSnapshot) SyncStatusesMultiCluster(ctx context.Context, mcClient multicluster.Client) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStatusesMultiCluster", ctx, mcClient)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncStatusesMultiCluster indicates an expected call of SyncStatusesMultiCluster
func (mr *MockSnapshotMockRecorder) SyncStatusesMultiCluster(ctx, mcClient interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStatusesMultiCluster", reflect.TypeOf((*MockSnapshot)(nil).SyncStatusesMultiCluster), ctx, mcClient)
}

// MarshalJSON mocks base method
func (m *MockSnapshot) MarshalJSON() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalJSON")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalJSON indicates an expected call of MarshalJSON
func (mr *MockSnapshotMockRecorder) MarshalJSON() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalJSON", reflect.TypeOf((*MockSnapshot)(nil).MarshalJSON))
}

// MockBuilder is a mock of Builder interface
type MockBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockBuilderMockRecorder
}

// MockBuilderMockRecorder is the mock recorder for MockBuilder
type MockBuilderMockRecorder struct {
	mock *MockBuilder
}

// NewMockBuilder creates a new mock instance
func NewMockBuilder(ctrl *gomock.Controller) *MockBuilder {
	mock := &MockBuilder{ctrl: ctrl}
	mock.recorder = &MockBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBuilder) EXPECT() *MockBuilderMockRecorder {
	return m.recorder
}

// BuildSnapshot mocks base method
func (m *MockBuilder) BuildSnapshot(ctx context.Context, name string, opts input.BuildOptions) (input.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildSnapshot", ctx, name, opts)
	ret0, _ := ret[0].(input.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildSnapshot indicates an expected call of BuildSnapshot
func (mr *MockBuilderMockRecorder) BuildSnapshot(ctx, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildSnapshot", reflect.TypeOf((*MockBuilder)(nil).BuildSnapshot), ctx, name, opts)
}
