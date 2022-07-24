// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/interfaces.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	model "github.com/danilashushkanov/student-service/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockStudentRepository is a mock of StudentRepository interface.
type MockStudentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockStudentRepositoryMockRecorder
}

// MockStudentRepositoryMockRecorder is the mock recorder for MockStudentRepository.
type MockStudentRepositoryMockRecorder struct {
	mock *MockStudentRepository
}

// NewMockStudentRepository creates a new mock instance.
func NewMockStudentRepository(ctrl *gomock.Controller) *MockStudentRepository {
	mock := &MockStudentRepository{ctrl: ctrl}
	mock.recorder = &MockStudentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStudentRepository) EXPECT() *MockStudentRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockStudentRepository) Create(arg0 context.Context, arg1 *model.Student) (*model.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*model.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockStudentRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStudentRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockStudentRepository) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockStudentRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStudentRepository)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockStudentRepository) Get(arg0 context.Context, arg1 int64) (*model.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*model.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStudentRepositoryMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStudentRepository)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockStudentRepository) List(arg0 context.Context, arg1 *StudentListFilter) ([]*model.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*model.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockStudentRepositoryMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStudentRepository)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockStudentRepository) Update(arg0 context.Context, arg1 *model.Student) (*model.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*model.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockStudentRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStudentRepository)(nil).Update), arg0, arg1)
}

// MockTeacherRepository is a mock of TeacherRepository interface.
type MockTeacherRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTeacherRepositoryMockRecorder
}

// MockTeacherRepositoryMockRecorder is the mock recorder for MockTeacherRepository.
type MockTeacherRepositoryMockRecorder struct {
	mock *MockTeacherRepository
}

// NewMockTeacherRepository creates a new mock instance.
func NewMockTeacherRepository(ctrl *gomock.Controller) *MockTeacherRepository {
	mock := &MockTeacherRepository{ctrl: ctrl}
	mock.recorder = &MockTeacherRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeacherRepository) EXPECT() *MockTeacherRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTeacherRepository) Create(arg0 context.Context, arg1 []*model.Teacher) ([]*model.Teacher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].([]*model.Teacher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTeacherRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTeacherRepository)(nil).Create), arg0, arg1)
}

// DeleteByStudentID mocks base method.
func (m *MockTeacherRepository) DeleteByStudentID(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByStudentID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByStudentID indicates an expected call of DeleteByStudentID.
func (mr *MockTeacherRepositoryMockRecorder) DeleteByStudentID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByStudentID", reflect.TypeOf((*MockTeacherRepository)(nil).DeleteByStudentID), arg0, arg1)
}

// Get mocks base method.
func (m *MockTeacherRepository) Get(arg0 context.Context, arg1 int64) (*model.Teacher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*model.Teacher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTeacherRepositoryMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTeacherRepository)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockTeacherRepository) List(arg0 context.Context, arg1 *TeacherListFilter) ([]*model.Teacher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*model.Teacher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockTeacherRepositoryMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTeacherRepository)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockTeacherRepository) Update(arg0 context.Context, arg1 *model.Teacher) (*model.Teacher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*model.Teacher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTeacherRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTeacherRepository)(nil).Update), arg0, arg1)
}

// UpdateTeachers mocks base method.
func (m *MockTeacherRepository) UpdateTeachers(arg0 context.Context, arg1 int64, arg2 []*model.Teacher) ([]*model.Teacher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTeachers", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*model.Teacher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTeachers indicates an expected call of UpdateTeachers.
func (mr *MockTeacherRepositoryMockRecorder) UpdateTeachers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTeachers", reflect.TypeOf((*MockTeacherRepository)(nil).UpdateTeachers), arg0, arg1, arg2)
}
