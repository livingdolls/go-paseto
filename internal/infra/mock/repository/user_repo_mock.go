// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/livingdolls/go-paseto/internal/core/port/repository (interfaces: UserPortRepository)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	gomock "github.com/golang/mock/gomock"
	dto "github.com/livingdolls/go-paseto/internal/core/dto"
	request "github.com/livingdolls/go-paseto/internal/core/model/request"
	response "github.com/livingdolls/go-paseto/internal/core/model/response"
	reflect "reflect"
)

// MockUserPortRepository is a mock of UserPortRepository interface
type MockUserPortRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserPortRepositoryMockRecorder
}

// MockUserPortRepositoryMockRecorder is the mock recorder for MockUserPortRepository
type MockUserPortRepositoryMockRecorder struct {
	mock *MockUserPortRepository
}

// NewMockUserPortRepository creates a new mock instance
func NewMockUserPortRepository(ctrl *gomock.Controller) *MockUserPortRepository {
	mock := &MockUserPortRepository{ctrl: ctrl}
	mock.recorder = &MockUserPortRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserPortRepository) EXPECT() *MockUserPortRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method
func (m *MockUserPortRepository) CreateUser(arg0 *dto.UserDTO) (*dto.UserDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(*dto.UserDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockUserPortRepositoryMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserPortRepository)(nil).CreateUser), arg0)
}

// GetListUser mocks base method
func (m *MockUserPortRepository) GetListUser() (*[]response.RegisterUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListUser")
	ret0, _ := ret[0].(*[]response.RegisterUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListUser indicates an expected call of GetListUser
func (mr *MockUserPortRepositoryMockRecorder) GetListUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListUser", reflect.TypeOf((*MockUserPortRepository)(nil).GetListUser))
}

// GetUserById mocks base method
func (m *MockUserPortRepository) GetUserById(arg0 *request.GetUserByIdRequest) (response.GetUserByIdResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0)
	ret0, _ := ret[0].(response.GetUserByIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById
func (mr *MockUserPortRepositoryMockRecorder) GetUserById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserPortRepository)(nil).GetUserById), arg0)
}