// Copyright (C) 2019-2023, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/luxdefi/node/vms/avm/txs/mempool (interfaces: Mempool)

// Package mempool is a generated GoMock package.
package mempool

import (
	reflect "reflect"

	ids "github.com/luxdefi/node/ids"
	txs "github.com/luxdefi/node/vms/avm/txs"
	gomock "github.com/golang/mock/gomock"
)

// MockMempool is a mock of Mempool interface.
type MockMempool struct {
	ctrl     *gomock.Controller
	recorder *MockMempoolMockRecorder
}

// MockMempoolMockRecorder is the mock recorder for MockMempool.
type MockMempoolMockRecorder struct {
	mock *MockMempool
}

// NewMockMempool creates a new mock instance.
func NewMockMempool(ctrl *gomock.Controller) *MockMempool {
	mock := &MockMempool{ctrl: ctrl}
	mock.recorder = &MockMempoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMempool) EXPECT() *MockMempoolMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockMempool) Add(arg0 *txs.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockMempoolMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockMempool)(nil).Add), arg0)
}

// Get mocks base method.
func (m *MockMempool) Get(arg0 ids.ID) *txs.Tx {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*txs.Tx)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockMempoolMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMempool)(nil).Get), arg0)
}

// GetDropReason mocks base method.
func (m *MockMempool) GetDropReason(arg0 ids.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDropReason", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetDropReason indicates an expected call of GetDropReason.
func (mr *MockMempoolMockRecorder) GetDropReason(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDropReason", reflect.TypeOf((*MockMempool)(nil).GetDropReason), arg0)
}

// Has mocks base method.
func (m *MockMempool) Has(arg0 ids.ID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockMempoolMockRecorder) Has(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockMempool)(nil).Has), arg0)
}

// MarkDropped mocks base method.
func (m *MockMempool) MarkDropped(arg0 ids.ID, arg1 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MarkDropped", arg0, arg1)
}

// MarkDropped indicates an expected call of MarkDropped.
func (mr *MockMempoolMockRecorder) MarkDropped(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkDropped", reflect.TypeOf((*MockMempool)(nil).MarkDropped), arg0, arg1)
}

// Peek mocks base method.
func (m *MockMempool) Peek(arg0 int) *txs.Tx {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Peek", arg0)
	ret0, _ := ret[0].(*txs.Tx)
	return ret0
}

// Peek indicates an expected call of Peek.
func (mr *MockMempoolMockRecorder) Peek(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Peek", reflect.TypeOf((*MockMempool)(nil).Peek), arg0)
}

// Remove mocks base method.
func (m *MockMempool) Remove(arg0 []*txs.Tx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", arg0)
}

// Remove indicates an expected call of Remove.
func (mr *MockMempoolMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockMempool)(nil).Remove), arg0)
}

// RequestBuildBlock mocks base method.
func (m *MockMempool) RequestBuildBlock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RequestBuildBlock")
}

// RequestBuildBlock indicates an expected call of RequestBuildBlock.
func (mr *MockMempoolMockRecorder) RequestBuildBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestBuildBlock", reflect.TypeOf((*MockMempool)(nil).RequestBuildBlock))
}
