// Copyright (C) 2019-2023, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/luxdefi/node/snow/engine/snowman/block (interfaces: WithVerifyContext)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	block "github.com/luxdefi/node/snow/engine/snowman/block"
	gomock "github.com/golang/mock/gomock"
)

// MockWithVerifyContext is a mock of WithVerifyContext interface.
type MockWithVerifyContext struct {
	ctrl     *gomock.Controller
	recorder *MockWithVerifyContextMockRecorder
}

// MockWithVerifyContextMockRecorder is the mock recorder for MockWithVerifyContext.
type MockWithVerifyContextMockRecorder struct {
	mock *MockWithVerifyContext
}

// NewMockWithVerifyContext creates a new mock instance.
func NewMockWithVerifyContext(ctrl *gomock.Controller) *MockWithVerifyContext {
	mock := &MockWithVerifyContext{ctrl: ctrl}
	mock.recorder = &MockWithVerifyContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWithVerifyContext) EXPECT() *MockWithVerifyContextMockRecorder {
	return m.recorder
}

// ShouldVerifyWithContext mocks base method.
func (m *MockWithVerifyContext) ShouldVerifyWithContext(arg0 context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldVerifyWithContext", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShouldVerifyWithContext indicates an expected call of ShouldVerifyWithContext.
func (mr *MockWithVerifyContextMockRecorder) ShouldVerifyWithContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldVerifyWithContext", reflect.TypeOf((*MockWithVerifyContext)(nil).ShouldVerifyWithContext), arg0)
}

// VerifyWithContext mocks base method.
func (m *MockWithVerifyContext) VerifyWithContext(arg0 context.Context, arg1 *block.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyWithContext", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyWithContext indicates an expected call of VerifyWithContext.
func (mr *MockWithVerifyContextMockRecorder) VerifyWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyWithContext", reflect.TypeOf((*MockWithVerifyContext)(nil).VerifyWithContext), arg0, arg1)
}
