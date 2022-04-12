// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nashmaniac/golang-application-template/adapters (interfaces: Usecases)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/nashmaniac/golang-application-template/models"
)

// MockUsecases is a mock of Usecases interface.
type MockUsecases struct {
	ctrl     *gomock.Controller
	recorder *MockUsecasesMockRecorder
}

// MockUsecasesMockRecorder is the mock recorder for MockUsecases.
type MockUsecasesMockRecorder struct {
	mock *MockUsecases
}

// NewMockUsecases creates a new mock instance.
func NewMockUsecases(ctrl *gomock.Controller) *MockUsecases {
	mock := &MockUsecases{ctrl: ctrl}
	mock.recorder = &MockUsecasesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecases) EXPECT() *MockUsecasesMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUsecases) CreateUser(arg0 context.Context, arg1, arg2 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUsecasesMockRecorder) CreateUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUsecases)(nil).CreateUser), arg0, arg1, arg2)
}

// GetHealthz mocks base method.
func (m *MockUsecases) GetHealthz(arg0 context.Context, arg1 string) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHealthz", arg0, arg1)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHealthz indicates an expected call of GetHealthz.
func (mr *MockUsecasesMockRecorder) GetHealthz(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHealthz", reflect.TypeOf((*MockUsecases)(nil).GetHealthz), arg0, arg1)
}