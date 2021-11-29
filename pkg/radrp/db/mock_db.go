// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/radius/pkg/radrp/db (interfaces: RadrpDB)

// Package db is a generated GoMock package.
package db

import (
	context "context"
	reflect "reflect"

	azresources "github.com/Azure/radius/pkg/azure/azresources"
	gomock "github.com/golang/mock/gomock"
)

// MockRadrpDB is a mock of RadrpDB interface.
type MockRadrpDB struct {
	ctrl     *gomock.Controller
	recorder *MockRadrpDBMockRecorder
}

// MockRadrpDBMockRecorder is the mock recorder for MockRadrpDB.
type MockRadrpDBMockRecorder struct {
	mock *MockRadrpDB
}

// NewMockRadrpDB creates a new mock instance.
func NewMockRadrpDB(ctrl *gomock.Controller) *MockRadrpDB {
	mock := &MockRadrpDB{ctrl: ctrl}
	mock.recorder = &MockRadrpDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRadrpDB) EXPECT() *MockRadrpDBMockRecorder {
	return m.recorder
}

// AddAzureResourceConnection mocks base method.
func (m *MockRadrpDB) AddAzureResourceConnection(arg0 context.Context, arg1 string, arg2 AzureResource) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAzureResourceConnection", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAzureResourceConnection indicates an expected call of AddAzureResourceConnection.
func (mr *MockRadrpDBMockRecorder) AddAzureResourceConnection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAzureResourceConnection", reflect.TypeOf((*MockRadrpDB)(nil).AddAzureResourceConnection), arg0, arg1, arg2)
}

// DeleteAzureResource mocks base method.
func (m *MockRadrpDB) DeleteAzureResource(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAzureResource", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAzureResource indicates an expected call of DeleteAzureResource.
func (mr *MockRadrpDBMockRecorder) DeleteAzureResource(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAzureResource", reflect.TypeOf((*MockRadrpDB)(nil).DeleteAzureResource), arg0, arg1, arg2)
}

// DeleteOperationByID mocks base method.
func (m *MockRadrpDB) DeleteOperationByID(arg0 context.Context, arg1 azresources.ResourceID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOperationByID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOperationByID indicates an expected call of DeleteOperationByID.
func (mr *MockRadrpDBMockRecorder) DeleteOperationByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOperationByID", reflect.TypeOf((*MockRadrpDB)(nil).DeleteOperationByID), arg0, arg1)
}

// DeleteV3Application mocks base method.
func (m *MockRadrpDB) DeleteV3Application(arg0 context.Context, arg1 azresources.ResourceID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteV3Application", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteV3Application indicates an expected call of DeleteV3Application.
func (mr *MockRadrpDBMockRecorder) DeleteV3Application(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteV3Application", reflect.TypeOf((*MockRadrpDB)(nil).DeleteV3Application), arg0, arg1)
}

// DeleteV3Resource mocks base method.
func (m *MockRadrpDB) DeleteV3Resource(arg0 context.Context, arg1 azresources.ResourceID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteV3Resource", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteV3Resource indicates an expected call of DeleteV3Resource.
func (mr *MockRadrpDBMockRecorder) DeleteV3Resource(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteV3Resource", reflect.TypeOf((*MockRadrpDB)(nil).DeleteV3Resource), arg0, arg1)
}

// GetAzureResource mocks base method.
func (m *MockRadrpDB) GetAzureResource(arg0 context.Context, arg1, arg2 string) (AzureResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureResource", arg0, arg1, arg2)
	ret0, _ := ret[0].(AzureResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureResource indicates an expected call of GetAzureResource.
func (mr *MockRadrpDBMockRecorder) GetAzureResource(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureResource", reflect.TypeOf((*MockRadrpDB)(nil).GetAzureResource), arg0, arg1, arg2)
}

// GetOperationByID mocks base method.
func (m *MockRadrpDB) GetOperationByID(arg0 context.Context, arg1 azresources.ResourceID) (*Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperationByID", arg0, arg1)
	ret0, _ := ret[0].(*Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperationByID indicates an expected call of GetOperationByID.
func (mr *MockRadrpDBMockRecorder) GetOperationByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperationByID", reflect.TypeOf((*MockRadrpDB)(nil).GetOperationByID), arg0, arg1)
}

// GetV3Application mocks base method.
func (m *MockRadrpDB) GetV3Application(arg0 context.Context, arg1 azresources.ResourceID) (ApplicationResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetV3Application", arg0, arg1)
	ret0, _ := ret[0].(ApplicationResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetV3Application indicates an expected call of GetV3Application.
func (mr *MockRadrpDBMockRecorder) GetV3Application(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetV3Application", reflect.TypeOf((*MockRadrpDB)(nil).GetV3Application), arg0, arg1)
}

// GetV3Resource mocks base method.
func (m *MockRadrpDB) GetV3Resource(arg0 context.Context, arg1 azresources.ResourceID) (RadiusResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetV3Resource", arg0, arg1)
	ret0, _ := ret[0].(RadiusResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetV3Resource indicates an expected call of GetV3Resource.
func (mr *MockRadrpDBMockRecorder) GetV3Resource(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetV3Resource", reflect.TypeOf((*MockRadrpDB)(nil).GetV3Resource), arg0, arg1)
}

// ListAllAzureResourcesForApplication mocks base method.
func (m *MockRadrpDB) ListAllAzureResourcesForApplication(arg0 context.Context, arg1 azresources.ResourceID, arg2 string) ([]AzureResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllAzureResourcesForApplication", arg0, arg1, arg2)
	ret0, _ := ret[0].([]AzureResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllAzureResourcesForApplication indicates an expected call of ListAllAzureResourcesForApplication.
func (mr *MockRadrpDBMockRecorder) ListAllAzureResourcesForApplication(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllAzureResourcesForApplication", reflect.TypeOf((*MockRadrpDB)(nil).ListAllAzureResourcesForApplication), arg0, arg1, arg2)
}

// ListAllV3ResourcesByApplication mocks base method.
func (m *MockRadrpDB) ListAllV3ResourcesByApplication(arg0 context.Context, arg1 azresources.ResourceID) ([]RadiusResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllV3ResourcesByApplication", arg0, arg1)
	ret0, _ := ret[0].([]RadiusResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllV3ResourcesByApplication indicates an expected call of ListAllV3ResourcesByApplication.
func (mr *MockRadrpDBMockRecorder) ListAllV3ResourcesByApplication(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllV3ResourcesByApplication", reflect.TypeOf((*MockRadrpDB)(nil).ListAllV3ResourcesByApplication), arg0, arg1)
}

// ListAzureResourcesForResourceType mocks base method.
func (m *MockRadrpDB) ListAzureResourcesForResourceType(arg0 context.Context, arg1 azresources.ResourceID, arg2 string) ([]AzureResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureResourcesForResourceType", arg0, arg1, arg2)
	ret0, _ := ret[0].([]AzureResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAzureResourcesForResourceType indicates an expected call of ListAzureResourcesForResourceType.
func (mr *MockRadrpDBMockRecorder) ListAzureResourcesForResourceType(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureResourcesForResourceType", reflect.TypeOf((*MockRadrpDB)(nil).ListAzureResourcesForResourceType), arg0, arg1, arg2)
}

// ListV3Applications mocks base method.
func (m *MockRadrpDB) ListV3Applications(arg0 context.Context, arg1 azresources.ResourceID) ([]ApplicationResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListV3Applications", arg0, arg1)
	ret0, _ := ret[0].([]ApplicationResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListV3Applications indicates an expected call of ListV3Applications.
func (mr *MockRadrpDBMockRecorder) ListV3Applications(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListV3Applications", reflect.TypeOf((*MockRadrpDB)(nil).ListV3Applications), arg0, arg1)
}

// ListV3Resources mocks base method.
func (m *MockRadrpDB) ListV3Resources(arg0 context.Context, arg1 azresources.ResourceID) ([]RadiusResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListV3Resources", arg0, arg1)
	ret0, _ := ret[0].([]RadiusResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListV3Resources indicates an expected call of ListV3Resources.
func (mr *MockRadrpDBMockRecorder) ListV3Resources(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListV3Resources", reflect.TypeOf((*MockRadrpDB)(nil).ListV3Resources), arg0, arg1)
}

// PatchOperationByID mocks base method.
func (m *MockRadrpDB) PatchOperationByID(arg0 context.Context, arg1 azresources.ResourceID, arg2 *Operation) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchOperationByID", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchOperationByID indicates an expected call of PatchOperationByID.
func (mr *MockRadrpDBMockRecorder) PatchOperationByID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchOperationByID", reflect.TypeOf((*MockRadrpDB)(nil).PatchOperationByID), arg0, arg1, arg2)
}

// RemoveAzureResourceConnection mocks base method.
func (m *MockRadrpDB) RemoveAzureResourceConnection(arg0 context.Context, arg1, arg2, arg3 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAzureResourceConnection", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveAzureResourceConnection indicates an expected call of RemoveAzureResourceConnection.
func (mr *MockRadrpDBMockRecorder) RemoveAzureResourceConnection(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAzureResourceConnection", reflect.TypeOf((*MockRadrpDB)(nil).RemoveAzureResourceConnection), arg0, arg1, arg2, arg3)
}

// UpdateAzureResource mocks base method.
func (m *MockRadrpDB) UpdateAzureResource(arg0 context.Context, arg1 AzureResource) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAzureResource", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAzureResource indicates an expected call of UpdateAzureResource.
func (mr *MockRadrpDBMockRecorder) UpdateAzureResource(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAzureResource", reflect.TypeOf((*MockRadrpDB)(nil).UpdateAzureResource), arg0, arg1)
}

// UpdateV3ApplicationDefinition mocks base method.
func (m *MockRadrpDB) UpdateV3ApplicationDefinition(arg0 context.Context, arg1 ApplicationResource) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateV3ApplicationDefinition", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateV3ApplicationDefinition indicates an expected call of UpdateV3ApplicationDefinition.
func (mr *MockRadrpDBMockRecorder) UpdateV3ApplicationDefinition(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateV3ApplicationDefinition", reflect.TypeOf((*MockRadrpDB)(nil).UpdateV3ApplicationDefinition), arg0, arg1)
}

// UpdateV3ResourceDefinition mocks base method.
func (m *MockRadrpDB) UpdateV3ResourceDefinition(arg0 context.Context, arg1 azresources.ResourceID, arg2 RadiusResource) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateV3ResourceDefinition", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateV3ResourceDefinition indicates an expected call of UpdateV3ResourceDefinition.
func (mr *MockRadrpDBMockRecorder) UpdateV3ResourceDefinition(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateV3ResourceDefinition", reflect.TypeOf((*MockRadrpDB)(nil).UpdateV3ResourceDefinition), arg0, arg1, arg2)
}

// UpdateV3ResourceStatus mocks base method.
func (m *MockRadrpDB) UpdateV3ResourceStatus(arg0 context.Context, arg1 azresources.ResourceID, arg2 RadiusResource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateV3ResourceStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateV3ResourceStatus indicates an expected call of UpdateV3ResourceStatus.
func (mr *MockRadrpDBMockRecorder) UpdateV3ResourceStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateV3ResourceStatus", reflect.TypeOf((*MockRadrpDB)(nil).UpdateV3ResourceStatus), arg0, arg1, arg2)
}
