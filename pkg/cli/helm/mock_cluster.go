// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/project-radius/radius/pkg/cli/helm (interfaces: Interface)

// Package helm is a generated GoMock package.
package helm

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// CheckRadiusInstall mocks base method.
func (m *MockInterface) CheckRadiusInstall(arg0 string) (InstallState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRadiusInstall", arg0)
	ret0, _ := ret[0].(InstallState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckRadiusInstall indicates an expected call of CheckRadiusInstall.
func (mr *MockInterfaceMockRecorder) CheckRadiusInstall(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRadiusInstall", reflect.TypeOf((*MockInterface)(nil).CheckRadiusInstall), arg0)
}

// InstallRadius mocks base method.
func (m *MockInterface) InstallRadius(arg0 context.Context, arg1 ClusterOptions, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallRadius", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstallRadius indicates an expected call of InstallRadius.
func (mr *MockInterfaceMockRecorder) InstallRadius(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallRadius", reflect.TypeOf((*MockInterface)(nil).InstallRadius), arg0, arg1, arg2)
}

// UninstallRadius mocks base method.
func (m *MockInterface) UninstallRadius(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UninstallRadius", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UninstallRadius indicates an expected call of UninstallRadius.
func (mr *MockInterfaceMockRecorder) UninstallRadius(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallRadius", reflect.TypeOf((*MockInterface)(nil).UninstallRadius), arg0, arg1)
}
