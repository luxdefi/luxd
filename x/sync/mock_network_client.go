// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/luxfi/node/x/sync (interfaces: NetworkClient)
//
// Generated by this command:
//
//	mockgen -package=sync -destination=x/sync/mock_network_client.go github.com/luxfi/node/x/sync NetworkClient
//

// Package sync is a generated GoMock package.
package sync

import (
	context "context"
	reflect "reflect"

	ids "github.com/luxfi/node/ids"
	version "github.com/luxfi/node/version"
	gomock "go.uber.org/mock/gomock"
)

// MockNetworkClient is a mock of NetworkClient interface.
type MockNetworkClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkClientMockRecorder
}

// MockNetworkClientMockRecorder is the mock recorder for MockNetworkClient.
type MockNetworkClientMockRecorder struct {
	mock *MockNetworkClient
}

// NewMockNetworkClient creates a new mock instance.
func NewMockNetworkClient(ctrl *gomock.Controller) *MockNetworkClient {
	mock := &MockNetworkClient{ctrl: ctrl}
	mock.recorder = &MockNetworkClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkClient) EXPECT() *MockNetworkClientMockRecorder {
	return m.recorder
}

// AppRequestFailed mocks base method.
func (m *MockNetworkClient) AppRequestFailed(arg0 context.Context, arg1 ids.NodeID, arg2 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppRequestFailed", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppRequestFailed indicates an expected call of AppRequestFailed.
func (mr *MockNetworkClientMockRecorder) AppRequestFailed(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppRequestFailed", reflect.TypeOf((*MockNetworkClient)(nil).AppRequestFailed), arg0, arg1, arg2)
}

// AppResponse mocks base method.
func (m *MockNetworkClient) AppResponse(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppResponse", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppResponse indicates an expected call of AppResponse.
func (mr *MockNetworkClientMockRecorder) AppResponse(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppResponse", reflect.TypeOf((*MockNetworkClient)(nil).AppResponse), arg0, arg1, arg2, arg3)
}

// Connected mocks base method.
func (m *MockNetworkClient) Connected(arg0 context.Context, arg1 ids.NodeID, arg2 *version.Application) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connected", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connected indicates an expected call of Connected.
func (mr *MockNetworkClientMockRecorder) Connected(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connected", reflect.TypeOf((*MockNetworkClient)(nil).Connected), arg0, arg1, arg2)
}

// Disconnected mocks base method.
func (m *MockNetworkClient) Disconnected(arg0 context.Context, arg1 ids.NodeID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disconnected", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Disconnected indicates an expected call of Disconnected.
func (mr *MockNetworkClientMockRecorder) Disconnected(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnected", reflect.TypeOf((*MockNetworkClient)(nil).Disconnected), arg0, arg1)
}

// Request mocks base method.
func (m *MockNetworkClient) Request(arg0 context.Context, arg1 ids.NodeID, arg2 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Request indicates an expected call of Request.
func (mr *MockNetworkClientMockRecorder) Request(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockNetworkClient)(nil).Request), arg0, arg1, arg2)
}

// RequestAny mocks base method.
func (m *MockNetworkClient) RequestAny(arg0 context.Context, arg1 []byte) (ids.NodeID, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestAny", arg0, arg1)
	ret0, _ := ret[0].(ids.NodeID)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RequestAny indicates an expected call of RequestAny.
func (mr *MockNetworkClientMockRecorder) RequestAny(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestAny", reflect.TypeOf((*MockNetworkClient)(nil).RequestAny), arg0, arg1)
}
