// Code generated by MockGen. DO NOT EDIT.
// Source: internal/controllers/consumer_transactions/repository/repository.go
//
// Generated by this command:
//
//	mockgen -package mock -destination internal/controllers/consumer_transactions/repository/mock/mock.go -source=internal/controllers/consumer_transactions/repository/repository.go
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	repository "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/repository"
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
func (m *MockRepository) CountList(ctx context.Context, consumerID int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountList", ctx, consumerID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountList indicates an expected call of CountList.
func (mr *MockRepositoryMockRecorder) CountList(ctx, consumerID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountList", reflect.TypeOf((*MockRepository)(nil).CountList), ctx, consumerID)
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

// CreateAffiliatedDealer mocks base method.
func (m *MockRepository) CreateAffiliatedDealer(ctx context.Context, arg *repository.CreateAffiliatedDealerParams) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAffiliatedDealer", ctx, arg)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAffiliatedDealer indicates an expected call of CreateAffiliatedDealer.
func (mr *MockRepositoryMockRecorder) CreateAffiliatedDealer(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAffiliatedDealer", reflect.TypeOf((*MockRepository)(nil).CreateAffiliatedDealer), ctx, arg)
}

// CreateConsumers mocks base method.
func (m *MockRepository) CreateConsumers(ctx context.Context, arg *repository.CreateConsumersParams) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateConsumers", ctx, arg)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateConsumers indicates an expected call of CreateConsumers.
func (mr *MockRepositoryMockRecorder) CreateConsumers(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConsumers", reflect.TypeOf((*MockRepository)(nil).CreateConsumers), ctx, arg)
}

// CreateTx mocks base method.
func (m *MockRepository) CreateTx(ctx context.Context, arg *repository.CreateTxParams) (repository.CreateTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTx", ctx, arg)
	ret0, _ := ret[0].(repository.CreateTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTx indicates an expected call of CreateTx.
func (mr *MockRepositoryMockRecorder) CreateTx(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTx", reflect.TypeOf((*MockRepository)(nil).CreateTx), ctx, arg)
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

// DelIdempotencyCreate mocks base method.
func (m *MockRepository) DelIdempotencyCreate(ctx context.Context, key string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DelIdempotencyCreate", ctx, key)
}

// DelIdempotencyCreate indicates an expected call of DelIdempotencyCreate.
func (mr *MockRepositoryMockRecorder) DelIdempotencyCreate(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelIdempotencyCreate", reflect.TypeOf((*MockRepository)(nil).DelIdempotencyCreate), ctx, key)
}

// DelIdempotencyUpdate mocks base method.
func (m *MockRepository) DelIdempotencyUpdate(ctx context.Context, key string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DelIdempotencyUpdate", ctx, key)
}

// DelIdempotencyUpdate indicates an expected call of DelIdempotencyUpdate.
func (mr *MockRepositoryMockRecorder) DelIdempotencyUpdate(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelIdempotencyUpdate", reflect.TypeOf((*MockRepository)(nil).DelIdempotencyUpdate), ctx, key)
}

// FindByID mocks base method.
func (m *MockRepository) FindByID(ctx context.Context, id int64) (*repository.ConsumerTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*repository.ConsumerTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockRepositoryMockRecorder) FindByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockRepository)(nil).FindByID), ctx, id)
}

// Get mocks base method.
func (m *MockRepository) Get(ctx context.Context, key string) (*repository.ConsumerTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(*repository.ConsumerTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRepositoryMockRecorder) Get(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), ctx, key)
}

// GetIdempotencyCreate mocks base method.
func (m *MockRepository) GetIdempotencyCreate(ctx context.Context, key string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIdempotencyCreate", ctx, key)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIdempotencyCreate indicates an expected call of GetIdempotencyCreate.
func (mr *MockRepositoryMockRecorder) GetIdempotencyCreate(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIdempotencyCreate", reflect.TypeOf((*MockRepository)(nil).GetIdempotencyCreate), ctx, key)
}

// GetIdempotencyUpdate mocks base method.
func (m *MockRepository) GetIdempotencyUpdate(ctx context.Context, key string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIdempotencyUpdate", ctx, key)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIdempotencyUpdate indicates an expected call of GetIdempotencyUpdate.
func (mr *MockRepositoryMockRecorder) GetIdempotencyUpdate(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIdempotencyUpdate", reflect.TypeOf((*MockRepository)(nil).GetIdempotencyUpdate), ctx, key)
}

// List mocks base method.
func (m *MockRepository) List(ctx context.Context, arg *repository.ListParams) ([]*repository.ConsumerTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, arg)
	ret0, _ := ret[0].([]*repository.ConsumerTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRepositoryMockRecorder) List(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List), ctx, arg)
}

// Put mocks base method.
func (m *MockRepository) Put(ctx context.Context, key string, arg *repository.ConsumerTransaction) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Put", ctx, key, arg)
}

// Put indicates an expected call of Put.
func (mr *MockRepositoryMockRecorder) Put(ctx, key, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockRepository)(nil).Put), ctx, key, arg)
}

// PutIdempotencyCreate mocks base method.
func (m *MockRepository) PutIdempotencyCreate(ctx context.Context, key string, arg int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutIdempotencyCreate", ctx, key, arg)
}

// PutIdempotencyCreate indicates an expected call of PutIdempotencyCreate.
func (mr *MockRepositoryMockRecorder) PutIdempotencyCreate(ctx, key, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutIdempotencyCreate", reflect.TypeOf((*MockRepository)(nil).PutIdempotencyCreate), ctx, key, arg)
}

// PutIdempotencyUpdate mocks base method.
func (m *MockRepository) PutIdempotencyUpdate(ctx context.Context, key string, arg int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutIdempotencyUpdate", ctx, key, arg)
}

// PutIdempotencyUpdate indicates an expected call of PutIdempotencyUpdate.
func (mr *MockRepositoryMockRecorder) PutIdempotencyUpdate(ctx, key, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutIdempotencyUpdate", reflect.TypeOf((*MockRepository)(nil).PutIdempotencyUpdate), ctx, key, arg)
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

// UpdateTx mocks base method.
func (m *MockRepository) UpdateTx(ctx context.Context, arg *repository.UpdateTxParams) (repository.UpdateTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTx", ctx, arg)
	ret0, _ := ret[0].(repository.UpdateTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTx indicates an expected call of UpdateTx.
func (mr *MockRepositoryMockRecorder) UpdateTx(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTx", reflect.TypeOf((*MockRepository)(nil).UpdateTx), ctx, arg)
}