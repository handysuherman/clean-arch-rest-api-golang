// Code generated by MockGen. DO NOT EDIT.
// Source: internal/controllers/affiliated_dealers/repository/repository.go
//
// Generated by this command:
//
//	mockgen -package mock -destination internal/controllers/affiliated_dealers/repository/mock/mock.go -source=internal/controllers/affiliated_dealers/repository/repository.go
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	repository "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/repository"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
	isgomock struct{}
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CountList mocks base method.
func (m *MockRepository) CountList(ctx context.Context, affiliatedDealerName string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountList", ctx, affiliatedDealerName)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountList indicates an expected call of CountList.
func (mr *MockRepositoryMockRecorder) CountList(ctx, affiliatedDealerName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountList", reflect.TypeOf((*MockRepository)(nil).CountList), ctx, affiliatedDealerName)
}

// Create mocks base method.
func (m *MockRepository) Create(ctx context.Context, arg *repository.CreateParams) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, arg)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), ctx, arg)
}

// Del mocks base method.
func (m *MockRepository) Del(ctx context.Context, key string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Del", ctx, key)
}

// Del indicates an expected call of Del.
func (mr *MockRepositoryMockRecorder) Del(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockRepository)(nil).Del), ctx, key)
}

// FindByID mocks base method.
func (m *MockRepository) FindByID(ctx context.Context, id int64) (*repository.AffiliatedDealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*repository.AffiliatedDealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockRepositoryMockRecorder) FindByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockRepository)(nil).FindByID), ctx, id)
}

// Get mocks base method.
func (m *MockRepository) Get(ctx context.Context, key string) (*repository.AffiliatedDealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(*repository.AffiliatedDealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRepositoryMockRecorder) Get(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), ctx, key)
}

// List mocks base method.
func (m *MockRepository) List(ctx context.Context, arg *repository.ListParams) ([]*repository.AffiliatedDealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, arg)
	ret0, _ := ret[0].([]*repository.AffiliatedDealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRepositoryMockRecorder) List(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List), ctx, arg)
}

// Put mocks base method.
func (m *MockRepository) Put(ctx context.Context, key string, arg *repository.AffiliatedDealer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Put", ctx, key, arg)
}

// Put indicates an expected call of Put.
func (mr *MockRepositoryMockRecorder) Put(ctx, key, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockRepository)(nil).Put), ctx, key, arg)
}

// Update mocks base method.
func (m *MockRepository) Update(ctx context.Context, arg *repository.UpdateParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), ctx, arg)
}
