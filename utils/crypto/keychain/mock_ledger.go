// Copyright (C) 2019-2023, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/luxdefi/node/utils/crypto/keychain (interfaces: Ledger)

// Package keychain is a generated GoMock package.
package keychain

import (
	reflect "reflect"

	ids "github.com/luxdefi/node/ids"
	version "github.com/luxdefi/node/version"
	gomock "github.com/golang/mock/gomock"
)

// MockLedger is a mock of Ledger interface.
type MockLedger struct {
	ctrl     *gomock.Controller
	recorder *MockLedgerMockRecorder
}

// MockLedgerMockRecorder is the mock recorder for MockLedger.
type MockLedgerMockRecorder struct {
	mock *MockLedger
}

// NewMockLedger creates a new mock instance.
func NewMockLedger(ctrl *gomock.Controller) *MockLedger {
	mock := &MockLedger{ctrl: ctrl}
	mock.recorder = &MockLedgerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLedger) EXPECT() *MockLedgerMockRecorder {
	return m.recorder
}

// Address mocks base method.
func (m *MockLedger) Address(arg0 string, arg1 uint32) (ids.ShortID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Address", arg0, arg1)
	ret0, _ := ret[0].(ids.ShortID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Address indicates an expected call of Address.
func (mr *MockLedgerMockRecorder) Address(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockLedger)(nil).Address), arg0, arg1)
}

// Addresses mocks base method.
func (m *MockLedger) Addresses(arg0 []uint32) ([]ids.ShortID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addresses", arg0)
	ret0, _ := ret[0].([]ids.ShortID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Addresses indicates an expected call of Addresses.
func (mr *MockLedgerMockRecorder) Addresses(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addresses", reflect.TypeOf((*MockLedger)(nil).Addresses), arg0)
}

// Disconnect mocks base method.
func (m *MockLedger) Disconnect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disconnect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Disconnect indicates an expected call of Disconnect.
func (mr *MockLedgerMockRecorder) Disconnect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockLedger)(nil).Disconnect))
}

// Sign mocks base method.
func (m *MockLedger) Sign(arg0 []byte, arg1 []uint32) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", arg0, arg1)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign.
func (mr *MockLedgerMockRecorder) Sign(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockLedger)(nil).Sign), arg0, arg1)
}

// SignHash mocks base method.
func (m *MockLedger) SignHash(arg0 []byte, arg1 []uint32) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignHash", arg0, arg1)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignHash indicates an expected call of SignHash.
func (mr *MockLedgerMockRecorder) SignHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignHash", reflect.TypeOf((*MockLedger)(nil).SignHash), arg0, arg1)
}

// Version mocks base method.
func (m *MockLedger) Version() (*version.Semantic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(*version.Semantic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockLedgerMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockLedger)(nil).Version))
}
