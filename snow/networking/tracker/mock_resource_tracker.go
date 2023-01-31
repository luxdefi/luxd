// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/luxdefi/node/snow/networking/tracker (interfaces: Tracker)

// Package tracker is a generated GoMock package.
package tracker

import (
	reflect "reflect"
	time "time"

	ids "github.com/luxdefi/node/ids"
	gomock "github.com/golang/mock/gomock"
)

// MockTracker is a mock of Tracker interface.
type MockTracker struct {
	ctrl     *gomock.Controller
	recorder *MockTrackerMockRecorder
}

// MockTrackerMockRecorder is the mock recorder for MockTracker.
type MockTrackerMockRecorder struct {
	mock *MockTracker
}

// NewMockTracker creates a new mock instance.
func NewMockTracker(ctrl *gomock.Controller) *MockTracker {
	mock := &MockTracker{ctrl: ctrl}
	mock.recorder = &MockTrackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTracker) EXPECT() *MockTrackerMockRecorder {
	return m.recorder
}

// TimeUntilUsage mocks base method.
func (m *MockTracker) TimeUntilUsage(arg0 ids.NodeID, arg1 time.Time, arg2 float64) time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TimeUntilUsage", arg0, arg1, arg2)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// TimeUntilUsage indicates an expected call of TimeUntilUsage.
func (mr *MockTrackerMockRecorder) TimeUntilUsage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeUntilUsage", reflect.TypeOf((*MockTracker)(nil).TimeUntilUsage), arg0, arg1, arg2)
}

// TotalUsage mocks base method.
func (m *MockTracker) TotalUsage() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TotalUsage")
	ret0, _ := ret[0].(float64)
	return ret0
}

// TotalUsage indicates an expected call of TotalUsage.
func (mr *MockTrackerMockRecorder) TotalUsage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TotalUsage", reflect.TypeOf((*MockTracker)(nil).TotalUsage))
}

// Usage mocks base method.
func (m *MockTracker) Usage(arg0 ids.NodeID, arg1 time.Time) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Usage", arg0, arg1)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Usage indicates an expected call of Usage.
func (mr *MockTrackerMockRecorder) Usage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Usage", reflect.TypeOf((*MockTracker)(nil).Usage), arg0, arg1)
}
