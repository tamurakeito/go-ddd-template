// Code generated by MockGen. DO NOT EDIT.
// Source: src/domain/repository/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	model "go-ddd-template/src/domain/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHelloRepository is a mock of HelloRepository interface.
type MockHelloRepository struct {
	ctrl     *gomock.Controller
	recorder *MockHelloRepositoryMockRecorder
}

// MockHelloRepositoryMockRecorder is the mock recorder for MockHelloRepository.
type MockHelloRepositoryMockRecorder struct {
	mock *MockHelloRepository
}

// NewMockHelloRepository creates a new mock instance.
func NewMockHelloRepository(ctrl *gomock.Controller) *MockHelloRepository {
	mock := &MockHelloRepository{ctrl: ctrl}
	mock.recorder = &MockHelloRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHelloRepository) EXPECT() *MockHelloRepositoryMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockHelloRepository) Find(id int) (model.Hello, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(model.Hello)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockHelloRepositoryMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockHelloRepository)(nil).Find), id)
}

// MockAuthRepository is a mock of AuthRepository interface.
type MockAuthRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRepositoryMockRecorder
}

// MockAuthRepositoryMockRecorder is the mock recorder for MockAuthRepository.
type MockAuthRepositoryMockRecorder struct {
	mock *MockAuthRepository
}

// NewMockAuthRepository creates a new mock instance.
func NewMockAuthRepository(ctrl *gomock.Controller) *MockAuthRepository {
	mock := &MockAuthRepository{ctrl: ctrl}
	mock.recorder = &MockAuthRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthRepository) EXPECT() *MockAuthRepositoryMockRecorder {
	return m.recorder
}

// FindUserId mocks base method.
func (m *MockAuthRepository) FindUserId(userId string) (model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserId", userId)
	ret0, _ := ret[0].(model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserId indicates an expected call of FindUserId.
func (mr *MockAuthRepositoryMockRecorder) FindUserId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserId", reflect.TypeOf((*MockAuthRepository)(nil).FindUserId), userId)
}
