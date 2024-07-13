// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/luxfi/node/vms/platformvm/state (interfaces: StakerIterator)
//
// Generated by this command:
//
//	mockgen -package=state -destination=vms/platformvm/state/mock_staker_iterator.go github.com/luxfi/node/vms/platformvm/state StakerIterator
//

// Package state is a generated GoMock package.
package state

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockStakerIterator is a mock of StakerIterator interface.
type MockStakerIterator struct {
	ctrl     *gomock.Controller
	recorder *MockStakerIteratorMockRecorder
}

// MockStakerIteratorMockRecorder is the mock recorder for MockStakerIterator.
type MockStakerIteratorMockRecorder struct {
	mock *MockStakerIterator
}

// NewMockStakerIterator creates a new mock instance.
func NewMockStakerIterator(ctrl *gomock.Controller) *MockStakerIterator {
	mock := &MockStakerIterator{ctrl: ctrl}
	mock.recorder = &MockStakerIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStakerIterator) EXPECT() *MockStakerIteratorMockRecorder {
	return m.recorder
}

// Next mocks base method.
func (m *MockStakerIterator) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockStakerIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockStakerIterator)(nil).Next))
}

// Release mocks base method.
func (m *MockStakerIterator) Release() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Release")
}

// Release indicates an expected call of Release.
func (mr *MockStakerIteratorMockRecorder) Release() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Release", reflect.TypeOf((*MockStakerIterator)(nil).Release))
}

// Value mocks base method.
func (m *MockStakerIterator) Value() *Staker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].(*Staker)
	return ret0
}

// Value indicates an expected call of Value.
func (mr *MockStakerIteratorMockRecorder) Value() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockStakerIterator)(nil).Value))
}
