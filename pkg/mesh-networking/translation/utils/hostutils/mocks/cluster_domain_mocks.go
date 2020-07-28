// Code generated by MockGen. DO NOT EDIT.
// Source: ./cluster_domain.go

// Package mock_hostutils is a generated GoMock package.
package mock_hostutils

import (
	gomock "github.com/golang/mock/gomock"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	reflect "reflect"
)

// MockClusterDomainRegistry is a mock of ClusterDomainRegistry interface.
type MockClusterDomainRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockClusterDomainRegistryMockRecorder
}

// MockClusterDomainRegistryMockRecorder is the mock recorder for MockClusterDomainRegistry.
type MockClusterDomainRegistryMockRecorder struct {
	mock *MockClusterDomainRegistry
}

// NewMockClusterDomainRegistry creates a new mock instance.
func NewMockClusterDomainRegistry(ctrl *gomock.Controller) *MockClusterDomainRegistry {
	mock := &MockClusterDomainRegistry{ctrl: ctrl}
	mock.recorder = &MockClusterDomainRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterDomainRegistry) EXPECT() *MockClusterDomainRegistryMockRecorder {
	return m.recorder
}

// GetClusterDomain mocks base method.
func (m *MockClusterDomainRegistry) GetClusterDomain(clusterName string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterDomain", clusterName)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetClusterDomain indicates an expected call of GetClusterDomain.
func (mr *MockClusterDomainRegistryMockRecorder) GetClusterDomain(clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterDomain", reflect.TypeOf((*MockClusterDomainRegistry)(nil).GetClusterDomain), clusterName)
}

// GetServiceLocalFQDN mocks base method.
func (m *MockClusterDomainRegistry) GetServiceLocalFQDN(serviceRef ezkube.ClusterResourceId) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceLocalFQDN", serviceRef)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetServiceLocalFQDN indicates an expected call of GetServiceLocalFQDN.
func (mr *MockClusterDomainRegistryMockRecorder) GetServiceLocalFQDN(serviceRef interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceLocalFQDN", reflect.TypeOf((*MockClusterDomainRegistry)(nil).GetServiceLocalFQDN), serviceRef)
}

// GetServiceGlobalFQDN mocks base method.
func (m *MockClusterDomainRegistry) GetServiceGlobalFQDN(serviceRef ezkube.ClusterResourceId) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceGlobalFQDN", serviceRef)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetServiceGlobalFQDN indicates an expected call of GetServiceGlobalFQDN.
func (mr *MockClusterDomainRegistryMockRecorder) GetServiceGlobalFQDN(serviceRef interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceGlobalFQDN", reflect.TypeOf((*MockClusterDomainRegistry)(nil).GetServiceGlobalFQDN), serviceRef)
}

// GetDestinationServiceFQDN mocks base method.
func (m *MockClusterDomainRegistry) GetDestinationServiceFQDN(originatingCluster string, serviceRef ezkube.ClusterResourceId) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDestinationServiceFQDN", originatingCluster, serviceRef)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDestinationServiceFQDN indicates an expected call of GetDestinationServiceFQDN.
func (mr *MockClusterDomainRegistryMockRecorder) GetDestinationServiceFQDN(originatingCluster, serviceRef interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDestinationServiceFQDN", reflect.TypeOf((*MockClusterDomainRegistry)(nil).GetDestinationServiceFQDN), originatingCluster, serviceRef)
}
