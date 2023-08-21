// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/radius-project/radius/pkg/corerp/handlers (interfaces: ResourceHandler)

// Package handlers is a generated GoMock package.
package handlers

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockResourceHandler is a mock of ResourceHandler interface.
type MockResourceHandler struct {
	ctrl     *gomock.Controller
	recorder *MockResourceHandlerMockRecorder
}

// MockResourceHandlerMockRecorder is the mock recorder for MockResourceHandler.
type MockResourceHandlerMockRecorder struct {
	mock *MockResourceHandler
}

// NewMockResourceHandler creates a new mock instance.
func NewMockResourceHandler(ctrl *gomock.Controller) *MockResourceHandler {
	mock := &MockResourceHandler{ctrl: ctrl}
	mock.recorder = &MockResourceHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceHandler) EXPECT() *MockResourceHandlerMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockResourceHandler) Delete(arg0 context.Context, arg1 *DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockResourceHandlerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockResourceHandler)(nil).Delete), arg0, arg1)
}

// Put mocks base method.
func (m *MockResourceHandler) Put(arg0 context.Context, arg1 *PutOptions) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockResourceHandlerMockRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockResourceHandler)(nil).Put), arg0, arg1)
}
