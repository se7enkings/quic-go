// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/se7enkings/quic-go (interfaces: Connection)

// Package quic is a generated GoMock package.
package quic

import (
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockConnection is a mock of Connection interface
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockConnection) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockConnectionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnection)(nil).Close))
}

// LocalAddr mocks base method
func (m *MockConnection) LocalAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr
func (mr *MockConnectionMockRecorder) LocalAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*MockConnection)(nil).LocalAddr))
}

// Read mocks base method
func (m *MockConnection) Read(arg0 []byte) (int, net.Addr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(net.Addr)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Read indicates an expected call of Read
func (mr *MockConnectionMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockConnection)(nil).Read), arg0)
}

// RemoteAddr mocks base method
func (m *MockConnection) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr
func (mr *MockConnectionMockRecorder) RemoteAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockConnection)(nil).RemoteAddr))
}

// SetCurrentRemoteAddr mocks base method
func (m *MockConnection) SetCurrentRemoteAddr(arg0 net.Addr) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCurrentRemoteAddr", arg0)
}

// SetCurrentRemoteAddr indicates an expected call of SetCurrentRemoteAddr
func (mr *MockConnectionMockRecorder) SetCurrentRemoteAddr(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCurrentRemoteAddr", reflect.TypeOf((*MockConnection)(nil).SetCurrentRemoteAddr), arg0)
}

// Write mocks base method
func (m *MockConnection) Write(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write
func (mr *MockConnectionMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockConnection)(nil).Write), arg0)
}
