// Copyright (C) 2019-2023, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/luxdefi/node/chains/atomic (interfaces: SharedMemory)

// Package atomic is a generated GoMock package.
package atomic

import (
	reflect "reflect"

	database "github.com/luxdefi/node/database"
	ids "github.com/luxdefi/node/ids"
	gomock "github.com/golang/mock/gomock"
)

// MockSharedMemory is a mock of SharedMemory interface.
type MockSharedMemory struct {
	ctrl     *gomock.Controller
	recorder *MockSharedMemoryMockRecorder
}

// MockSharedMemoryMockRecorder is the mock recorder for MockSharedMemory.
type MockSharedMemoryMockRecorder struct {
	mock *MockSharedMemory
}

// NewMockSharedMemory creates a new mock instance.
func NewMockSharedMemory(ctrl *gomock.Controller) *MockSharedMemory {
	mock := &MockSharedMemory{ctrl: ctrl}
	mock.recorder = &MockSharedMemoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSharedMemory) EXPECT() *MockSharedMemoryMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockSharedMemory) Apply(arg0 map[ids.ID]*Requests, arg1 ...database.Batch) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Apply", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Apply indicates an expected call of Apply.
func (mr *MockSharedMemoryMockRecorder) Apply(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockSharedMemory)(nil).Apply), varargs...)
}

// Get mocks base method.
func (m *MockSharedMemory) Get(arg0 ids.ID, arg1 [][]byte) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSharedMemoryMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSharedMemory)(nil).Get), arg0, arg1)
}

// Indexed mocks base method.
func (m *MockSharedMemory) Indexed(arg0 ids.ID, arg1 [][]byte, arg2, arg3 []byte, arg4 int) ([][]byte, []byte, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Indexed", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].([]byte)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Indexed indicates an expected call of Indexed.
func (mr *MockSharedMemoryMockRecorder) Indexed(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Indexed", reflect.TypeOf((*MockSharedMemory)(nil).Indexed), arg0, arg1, arg2, arg3, arg4)
}
