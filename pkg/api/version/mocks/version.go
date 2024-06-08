// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/diginfra/client-go/pkg/api/version (interfaces: ServiceClient)

// Package mock_version is a generated GoMock package.
package mock_version

import (
	context "context"
	reflect "reflect"

	version "github.com/diginfra/client-go/pkg/api/version"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockServiceClient is a mock of ServiceClient interface.
type MockServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockServiceClientMockRecorder
}

// MockServiceClientMockRecorder is the mock recorder for MockServiceClient.
type MockServiceClientMockRecorder struct {
	mock *MockServiceClient
}

// NewMockServiceClient creates a new mock instance.
func NewMockServiceClient(ctrl *gomock.Controller) *MockServiceClient {
	mock := &MockServiceClient{ctrl: ctrl}
	mock.recorder = &MockServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceClient) EXPECT() *MockServiceClientMockRecorder {
	return m.recorder
}

// Version mocks base method.
func (m *MockServiceClient) Version(arg0 context.Context, arg1 *version.Request, arg2 ...grpc.CallOption) (*version.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Version", varargs...)
	ret0, _ := ret[0].(*version.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockServiceClientMockRecorder) Version(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockServiceClient)(nil).Version), varargs...)
}