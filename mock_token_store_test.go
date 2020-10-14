// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/se7enkings/quic-go (interfaces: TokenStore)

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTokenStore is a mock of TokenStore interface
type MockTokenStore struct {
	ctrl     *gomock.Controller
	recorder *MockTokenStoreMockRecorder
}

// MockTokenStoreMockRecorder is the mock recorder for MockTokenStore
type MockTokenStoreMockRecorder struct {
	mock *MockTokenStore
}

// NewMockTokenStore creates a new mock instance
func NewMockTokenStore(ctrl *gomock.Controller) *MockTokenStore {
	mock := &MockTokenStore{ctrl: ctrl}
	mock.recorder = &MockTokenStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTokenStore) EXPECT() *MockTokenStoreMockRecorder {
	return m.recorder
}

// Pop mocks base method
func (m *MockTokenStore) Pop(arg0 string) *ClientToken {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pop", arg0)
	ret0, _ := ret[0].(*ClientToken)
	return ret0
}

// Pop indicates an expected call of Pop
func (mr *MockTokenStoreMockRecorder) Pop(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pop", reflect.TypeOf((*MockTokenStore)(nil).Pop), arg0)
}

// Put mocks base method
func (m *MockTokenStore) Put(arg0 string, arg1 *ClientToken) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Put", arg0, arg1)
}

// Put indicates an expected call of Put
func (mr *MockTokenStoreMockRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockTokenStore)(nil).Put), arg0, arg1)
}
