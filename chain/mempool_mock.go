// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: ./chain/mempool.go

package chain

import (
	reflect "reflect"

	ids "github.com/ava-labs/avalanchego/ids"
	set "github.com/ava-labs/avalanchego/utils/set"
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
func (m *MockMempool) Add(arg0 *Transaction) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockMempoolMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockMempool)(nil).Add), arg0)
}

// Len mocks base method.
func (m *MockMempool) Len() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

// Len indicates an expected call of Len.
func (mr *MockMempoolMockRecorder) Len() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Len", reflect.TypeOf((*MockMempool)(nil).Len))
}

// NewTxs mocks base method.
func (m *MockMempool) NewTxs(arg0 uint64) []*Transaction {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTxs", arg0)
	ret0, _ := ret[0].([]*Transaction)
	return ret0
}

// NewTxs indicates an expected call of NewTxs.
func (mr *MockMempoolMockRecorder) NewTxs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTxs", reflect.TypeOf((*MockMempool)(nil).NewTxs), arg0)
}

// PopMax mocks base method.
func (m *MockMempool) PopMax() (*Transaction, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PopMax")
	ret0, _ := ret[0].(*Transaction)
	ret1, _ := ret[1].(uint64)
	return ret0, ret1
}

// PopMax indicates an expected call of PopMax.
func (mr *MockMempoolMockRecorder) PopMax() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopMax", reflect.TypeOf((*MockMempool)(nil).PopMax))
}

// Prune mocks base method.
func (m *MockMempool) Prune(arg0 set.Set[ids.ID]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Prune", arg0)
}

// Prune indicates an expected call of Prune.
func (mr *MockMempoolMockRecorder) Prune(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prune", reflect.TypeOf((*MockMempool)(nil).Prune), arg0)
}
