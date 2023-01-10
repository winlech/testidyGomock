// Code generated by MockGen. DO NOT EDIT.
// Source: main.go

// Package main is a generated GoMock package.
package main

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockChecker is a mock of Checker interface.
type MockChecker struct {
	ctrl     *gomock.Controller
	recorder *MockCheckerMockRecorder
}

// MockCheckerMockRecorder is the mock recorder for MockChecker.
type MockCheckerMockRecorder struct {
	mock *MockChecker
}

// NewMockChecker creates a new mock instance.
func NewMockChecker(ctrl *gomock.Controller) *MockChecker {
	mock := &MockChecker{ctrl: ctrl}
	mock.recorder = &MockCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChecker) EXPECT() *MockCheckerMockRecorder {
	return m.recorder
}

// isDividedBy3 mocks base method.
func (m *MockChecker) isDividedBy3(nbr int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "isDividedBy3", nbr)
	ret0, _ := ret[0].(bool)
	return ret0
}

// isDividedBy3 indicates an expected call of isDividedBy3.
func (mr *MockCheckerMockRecorder) isDividedBy3(nbr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isDividedBy3", reflect.TypeOf((*MockChecker)(nil).isDividedBy3), nbr)
}

// isDividedBy3And5 mocks base method.
func (m *MockChecker) isDividedBy3And5(nbr int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "isDividedBy3And5", nbr)
	ret0, _ := ret[0].(bool)
	return ret0
}

// isDividedBy3And5 indicates an expected call of isDividedBy3And5.
func (mr *MockCheckerMockRecorder) isDividedBy3And5(nbr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isDividedBy3And5", reflect.TypeOf((*MockChecker)(nil).isDividedBy3And5), nbr)
}

// isDividedBy5 mocks base method.
func (m *MockChecker) isDividedBy5(nbr int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "isDividedBy5", nbr)
	ret0, _ := ret[0].(bool)
	return ret0
}

// isDividedBy5 indicates an expected call of isDividedBy5.
func (mr *MockCheckerMockRecorder) isDividedBy5(nbr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isDividedBy5", reflect.TypeOf((*MockChecker)(nil).isDividedBy5), nbr)
}