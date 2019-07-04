// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/int128/kubelogin/usecases (interfaces: Login,LoginAndExec,Authentication)

// Package mock_usecases is a generated GoMock package.
package mock_usecases

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	usecases "github.com/int128/kubelogin/usecases"
	reflect "reflect"
)

// MockLogin is a mock of Login interface
type MockLogin struct {
	ctrl     *gomock.Controller
	recorder *MockLoginMockRecorder
}

// MockLoginMockRecorder is the mock recorder for MockLogin
type MockLoginMockRecorder struct {
	mock *MockLogin
}

// NewMockLogin creates a new mock instance
func NewMockLogin(ctrl *gomock.Controller) *MockLogin {
	mock := &MockLogin{ctrl: ctrl}
	mock.recorder = &MockLoginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogin) EXPECT() *MockLoginMockRecorder {
	return m.recorder
}

// Do mocks base method
func (m *MockLogin) Do(arg0 context.Context, arg1 usecases.LoginIn) error {
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Do indicates an expected call of Do
func (mr *MockLoginMockRecorder) Do(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockLogin)(nil).Do), arg0, arg1)
}

// MockLoginAndExec is a mock of LoginAndExec interface
type MockLoginAndExec struct {
	ctrl     *gomock.Controller
	recorder *MockLoginAndExecMockRecorder
}

// MockLoginAndExecMockRecorder is the mock recorder for MockLoginAndExec
type MockLoginAndExecMockRecorder struct {
	mock *MockLoginAndExec
}

// NewMockLoginAndExec creates a new mock instance
func NewMockLoginAndExec(ctrl *gomock.Controller) *MockLoginAndExec {
	mock := &MockLoginAndExec{ctrl: ctrl}
	mock.recorder = &MockLoginAndExecMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLoginAndExec) EXPECT() *MockLoginAndExecMockRecorder {
	return m.recorder
}

// Do mocks base method
func (m *MockLoginAndExec) Do(arg0 context.Context, arg1 usecases.LoginAndExecIn) (*usecases.LoginAndExecOut, error) {
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].(*usecases.LoginAndExecOut)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do
func (mr *MockLoginAndExecMockRecorder) Do(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockLoginAndExec)(nil).Do), arg0, arg1)
}

// MockAuthentication is a mock of Authentication interface
type MockAuthentication struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationMockRecorder
}

// MockAuthenticationMockRecorder is the mock recorder for MockAuthentication
type MockAuthenticationMockRecorder struct {
	mock *MockAuthentication
}

// NewMockAuthentication creates a new mock instance
func NewMockAuthentication(ctrl *gomock.Controller) *MockAuthentication {
	mock := &MockAuthentication{ctrl: ctrl}
	mock.recorder = &MockAuthenticationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthentication) EXPECT() *MockAuthenticationMockRecorder {
	return m.recorder
}

// Do mocks base method
func (m *MockAuthentication) Do(arg0 context.Context, arg1 usecases.AuthenticationIn) (*usecases.AuthenticationOut, error) {
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].(*usecases.AuthenticationOut)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do
func (mr *MockAuthenticationMockRecorder) Do(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockAuthentication)(nil).Do), arg0, arg1)
}
