// Code generated by MockGen. DO NOT EDIT.
// Source: ./trie/trie.go

// Package mock_trie is a generated GoMock package.
package mock_trie

import (
	gomock "github.com/golang/mock/gomock"
	common "github.com/iotexproject/iotex-core/common"
	reflect "reflect"
)

// MockTrie is a mock of Trie interface
type MockTrie struct {
	ctrl     *gomock.Controller
	recorder *MockTrieMockRecorder
}

// MockTrieMockRecorder is the mock recorder for MockTrie
type MockTrieMockRecorder struct {
	mock *MockTrie
}

// NewMockTrie creates a new mock instance
func NewMockTrie(ctrl *gomock.Controller) *MockTrie {
	mock := &MockTrie{ctrl: ctrl}
	mock.recorder = &MockTrieMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTrie) EXPECT() *MockTrieMockRecorder {
	return m.recorder
}

// Insert mocks base method
func (m *MockTrie) Insert(key, value []byte) error {
	ret := m.ctrl.Call(m, "Insert", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockTrieMockRecorder) Insert(key, value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTrie)(nil).Insert), key, value)
}

// Get mocks base method
func (m *MockTrie) Get(key []byte) ([]byte, error) {
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTrieMockRecorder) Get(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTrie)(nil).Get), key)
}

// Update mocks base method
func (m *MockTrie) Update(key, value []byte) error {
	ret := m.ctrl.Call(m, "Update", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockTrieMockRecorder) Update(key, value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTrie)(nil).Update), key, value)
}

// Delete mocks base method
func (m *MockTrie) Delete(key []byte) error {
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockTrieMockRecorder) Delete(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTrie)(nil).Delete), key)
}

// RootHash mocks base method
func (m *MockTrie) RootHash() common.Hash32B {
	ret := m.ctrl.Call(m, "RootHash")
	ret0, _ := ret[0].(common.Hash32B)
	return ret0
}

// RootHash indicates an expected call of RootHash
func (mr *MockTrieMockRecorder) RootHash() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RootHash", reflect.TypeOf((*MockTrie)(nil).RootHash))
}
