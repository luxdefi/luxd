<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// Copyright (C) 2022, Lux Partners Limited. All rights reserved.
=======
// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
>>>>>>> a653cdfcf (Update imports)
=======
// Copyright (C) 2019-2022, Lux Partners Limited. All rights reserved.
>>>>>>> d7a7925ff (Update various imports)
=======
// Copyright (C) 2022, Lux Partners Limited. All rights reserved.
>>>>>>> c5eafdb72 (Update LICENSE)
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/luxdefi/node/vms/proposervm/tree (interfaces: Tree)

// Package tree is a generated GoMock package.
package tree

import (
<<<<<<< HEAD
	context "context"
=======
>>>>>>> a653cdfcf (Update imports)
	reflect "reflect"

	snowman "github.com/luxdefi/node/snow/consensus/snowman"
	gomock "github.com/golang/mock/gomock"
)

// MockTree is a mock of Tree interface.
type MockTree struct {
	ctrl     *gomock.Controller
	recorder *MockTreeMockRecorder
}

// MockTreeMockRecorder is the mock recorder for MockTree.
type MockTreeMockRecorder struct {
	mock *MockTree
}

// NewMockTree creates a new mock instance.
func NewMockTree(ctrl *gomock.Controller) *MockTree {
	mock := &MockTree{ctrl: ctrl}
	mock.recorder = &MockTreeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTree) EXPECT() *MockTreeMockRecorder {
	return m.recorder
}

// Accept mocks base method.
<<<<<<< HEAD
func (m *MockTree) Accept(arg0 context.Context, arg1 snowman.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", arg0, arg1)
=======
func (m *MockTree) Accept(arg0 snowman.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", arg0)
>>>>>>> a653cdfcf (Update imports)
	ret0, _ := ret[0].(error)
	return ret0
}

// Accept indicates an expected call of Accept.
<<<<<<< HEAD
func (mr *MockTreeMockRecorder) Accept(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockTree)(nil).Accept), arg0, arg1)
=======
func (mr *MockTreeMockRecorder) Accept(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockTree)(nil).Accept), arg0)
>>>>>>> a653cdfcf (Update imports)
}

// Add mocks base method.
func (m *MockTree) Add(arg0 snowman.Block) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", arg0)
}

// Add indicates an expected call of Add.
func (mr *MockTreeMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTree)(nil).Add), arg0)
}

// Get mocks base method.
func (m *MockTree) Get(arg0 snowman.Block) (snowman.Block, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(snowman.Block)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTreeMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTree)(nil).Get), arg0)
}
