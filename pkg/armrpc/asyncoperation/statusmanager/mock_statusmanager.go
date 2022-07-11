// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/project-radius/radius/pkg/armrpc/asyncoperation/statusmanager (interfaces: StatusManager)

// Package statusmanager is a generated GoMock package.
package statusmanager

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	servicecontext "github.com/project-radius/radius/pkg/armrpc/servicecontext"
	armerrors "github.com/project-radius/radius/pkg/radrp/armerrors"
	resources "github.com/project-radius/radius/pkg/ucp/resources"
)

// MockStatusManager is a mock of StatusManager interface.
type MockStatusManager struct {
	ctrl     *gomock.Controller
	recorder *MockStatusManagerMockRecorder
}

// MockStatusManagerMockRecorder is the mock recorder for MockStatusManager.
type MockStatusManagerMockRecorder struct {
	mock *MockStatusManager
}

// NewMockStatusManager creates a new mock instance.
func NewMockStatusManager(ctrl *gomock.Controller) *MockStatusManager {
	mock := &MockStatusManager{ctrl: ctrl}
	mock.recorder = &MockStatusManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatusManager) EXPECT() *MockStatusManagerMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockStatusManager) Delete(arg0 context.Context, arg1 resources.ID, arg2 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockStatusManagerMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStatusManager)(nil).Delete), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockStatusManager) Get(arg0 context.Context, arg1 resources.ID, arg2 uuid.UUID) (*Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStatusManagerMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStatusManager)(nil).Get), arg0, arg1, arg2)
}

// QueueAsyncOperation mocks base method.
func (m *MockStatusManager) QueueAsyncOperation(arg0 context.Context, arg1 *servicecontext.ARMRequestContext, arg2 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueAsyncOperation", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueueAsyncOperation indicates an expected call of QueueAsyncOperation.
func (mr *MockStatusManagerMockRecorder) QueueAsyncOperation(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueAsyncOperation", reflect.TypeOf((*MockStatusManager)(nil).QueueAsyncOperation), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockStatusManager) Update(arg0 context.Context, arg1 resources.ID, arg2 uuid.UUID, arg3 v1.ProvisioningState, arg4 *time.Time, arg5 *armerrors.ErrorDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockStatusManagerMockRecorder) Update(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStatusManager)(nil).Update), arg0, arg1, arg2, arg3, arg4, arg5)
}
