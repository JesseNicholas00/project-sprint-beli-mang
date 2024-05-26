// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/JesseNicholas00/BeliMang/repos/merchant (interfaces: MerchantRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	merchant "github.com/JesseNicholas00/BeliMang/repos/merchant"
	gomock "github.com/golang/mock/gomock"
)

// MockMerchantRepository is a mock of MerchantRepository interface.
type MockMerchantRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMerchantRepositoryMockRecorder
}

// MockMerchantRepositoryMockRecorder is the mock recorder for MockMerchantRepository.
type MockMerchantRepositoryMockRecorder struct {
	mock *MockMerchantRepository
}

// NewMockMerchantRepository creates a new mock instance.
func NewMockMerchantRepository(ctrl *gomock.Controller) *MockMerchantRepository {
	mock := &MockMerchantRepository{ctrl: ctrl}
	mock.recorder = &MockMerchantRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMerchantRepository) EXPECT() *MockMerchantRepositoryMockRecorder {
	return m.recorder
}

// CreateMerchant mocks base method.
func (m *MockMerchantRepository) CreateMerchant(arg0 context.Context, arg1 merchant.Merchant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMerchant", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMerchant indicates an expected call of CreateMerchant.
func (mr *MockMerchantRepositoryMockRecorder) CreateMerchant(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMerchant", reflect.TypeOf((*MockMerchantRepository)(nil).CreateMerchant), arg0, arg1)
}

// CreateMerchantItem mocks base method.
func (m *MockMerchantRepository) CreateMerchantItem(arg0 context.Context, arg1 merchant.MerchantItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMerchantItem", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMerchantItem indicates an expected call of CreateMerchantItem.
func (mr *MockMerchantRepositoryMockRecorder) CreateMerchantItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMerchantItem", reflect.TypeOf((*MockMerchantRepository)(nil).CreateMerchantItem), arg0, arg1)
}

// FindMerchantById mocks base method.
func (m *MockMerchantRepository) FindMerchantById(arg0 context.Context, arg1 string) (merchant.Merchant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMerchantById", arg0, arg1)
	ret0, _ := ret[0].(merchant.Merchant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMerchantById indicates an expected call of FindMerchantById.
func (mr *MockMerchantRepositoryMockRecorder) FindMerchantById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMerchantById", reflect.TypeOf((*MockMerchantRepository)(nil).FindMerchantById), arg0, arg1)
}

// ListItemsByIds mocks base method.
func (m *MockMerchantRepository) ListItemsByIds(arg0 context.Context, arg1 []string) ([]merchant.MerchantItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListItemsByIds", arg0, arg1)
	ret0, _ := ret[0].([]merchant.MerchantItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListItemsByIds indicates an expected call of ListItemsByIds.
func (mr *MockMerchantRepositoryMockRecorder) ListItemsByIds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListItemsByIds", reflect.TypeOf((*MockMerchantRepository)(nil).ListItemsByIds), arg0, arg1)
}

// ListMerchantsByIds mocks base method.
func (m *MockMerchantRepository) ListMerchantsByIds(arg0 context.Context, arg1 []string) ([]merchant.Merchant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMerchantsByIds", arg0, arg1)
	ret0, _ := ret[0].([]merchant.Merchant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMerchantsByIds indicates an expected call of ListMerchantsByIds.
func (mr *MockMerchantRepositoryMockRecorder) ListMerchantsByIds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMerchantsByIds", reflect.TypeOf((*MockMerchantRepository)(nil).ListMerchantsByIds), arg0, arg1)
}