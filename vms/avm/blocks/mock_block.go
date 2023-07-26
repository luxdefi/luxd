// Copyright (C) 2019-2023, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/luxdefi/node/vms/avm/blocks (interfaces: Block)

// Package blocks is a generated GoMock package.
package blocks

import (
	reflect "reflect"
	time "time"

	codec "github.com/luxdefi/node/codec"
	ids "github.com/luxdefi/node/ids"
	snow "github.com/luxdefi/node/snow"
	txs "github.com/luxdefi/node/vms/avm/txs"
	gomock "github.com/golang/mock/gomock"
)

// MockBlock is a mock of Block interface.
type MockBlock struct {
	ctrl     *gomock.Controller
	recorder *MockBlockMockRecorder
}

// MockBlockMockRecorder is the mock recorder for MockBlock.
type MockBlockMockRecorder struct {
	mock *MockBlock
}

// NewMockBlock creates a new mock instance.
func NewMockBlock(ctrl *gomock.Controller) *MockBlock {
	mock := &MockBlock{ctrl: ctrl}
	mock.recorder = &MockBlockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlock) EXPECT() *MockBlockMockRecorder {
	return m.recorder
}

// Bytes mocks base method.
func (m *MockBlock) Bytes() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bytes")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Bytes indicates an expected call of Bytes.
func (mr *MockBlockMockRecorder) Bytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bytes", reflect.TypeOf((*MockBlock)(nil).Bytes))
}

// Height mocks base method.
func (m *MockBlock) Height() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Height")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Height indicates an expected call of Height.
func (mr *MockBlockMockRecorder) Height() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Height", reflect.TypeOf((*MockBlock)(nil).Height))
}

// ID mocks base method.
func (m *MockBlock) ID() ids.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(ids.ID)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockBlockMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockBlock)(nil).ID))
}

// InitCtx mocks base method.
func (m *MockBlock) InitCtx(arg0 *snow.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitCtx", arg0)
}

// InitCtx indicates an expected call of InitCtx.
func (mr *MockBlockMockRecorder) InitCtx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitCtx", reflect.TypeOf((*MockBlock)(nil).InitCtx), arg0)
}

// MerkleRoot mocks base method.
func (m *MockBlock) MerkleRoot() ids.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MerkleRoot")
	ret0, _ := ret[0].(ids.ID)
	return ret0
}

// MerkleRoot indicates an expected call of MerkleRoot.
func (mr *MockBlockMockRecorder) MerkleRoot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MerkleRoot", reflect.TypeOf((*MockBlock)(nil).MerkleRoot))
}

// Parent mocks base method.
func (m *MockBlock) Parent() ids.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parent")
	ret0, _ := ret[0].(ids.ID)
	return ret0
}

// Parent indicates an expected call of Parent.
func (mr *MockBlockMockRecorder) Parent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parent", reflect.TypeOf((*MockBlock)(nil).Parent))
}

// Timestamp mocks base method.
func (m *MockBlock) Timestamp() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timestamp")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Timestamp indicates an expected call of Timestamp.
func (mr *MockBlockMockRecorder) Timestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timestamp", reflect.TypeOf((*MockBlock)(nil).Timestamp))
}

// Txs mocks base method.
func (m *MockBlock) Txs() []*txs.Tx {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Txs")
	ret0, _ := ret[0].([]*txs.Tx)
	return ret0
}

// Txs indicates an expected call of Txs.
func (mr *MockBlockMockRecorder) Txs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Txs", reflect.TypeOf((*MockBlock)(nil).Txs))
}

// initialize mocks base method.
func (m *MockBlock) initialize(arg0 []byte, arg1 codec.Manager) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "initialize", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// initialize indicates an expected call of initialize.
func (mr *MockBlockMockRecorder) initialize(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "initialize", reflect.TypeOf((*MockBlock)(nil).initialize), arg0, arg1)
}
