// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/JesseNicholas00/BeliMang/services/merchant (interfaces: MerchantService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	merchant "github.com/JesseNicholas00/BeliMang/services/merchant"
	gomock "github.com/golang/mock/gomock"
)

// MockMerchantService is a mock of MerchantService interface.
type MockMerchantService struct {
	ctrl     *gomock.Controller
	recorder *MockMerchantServiceMockRecorder
}

// MockMerchantServiceMockRecorder is the mock recorder for MockMerchantService.
type MockMerchantServiceMockRecorder struct {
	mock *MockMerchantService
}

// NewMockMerchantService creates a new mock instance.
func NewMockMerchantService(ctrl *gomock.Controller) *MockMerchantService {
	mock := &MockMerchantService{ctrl: ctrl}
	mock.recorder = &MockMerchantServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMerchantService) EXPECT() *MockMerchantServiceMockRecorder {
	return m.recorder
}

// AdminListMerchant mocks base method.
func (m *MockMerchantService) AdminListMerchant(arg0 context.Context, arg1 merchant.AdminListMerchantReq, arg2 *merchant.AdminListMerchantRes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdminListMerchant", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AdminListMerchant indicates an expected call of AdminListMerchant.
func (mr *MockMerchantServiceMockRecorder) AdminListMerchant(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminListMerchant", reflect.TypeOf((*MockMerchantService)(nil).AdminListMerchant), arg0, arg1, arg2)
}

// CreateMerchant mocks base method.
func (m *MockMerchantService) CreateMerchant(arg0 context.Context, arg1 merchant.CreateMerchantReq, arg2 *merchant.CreateMerchantRes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMerchant", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMerchant indicates an expected call of CreateMerchant.
func (mr *MockMerchantServiceMockRecorder) CreateMerchant(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMerchant", reflect.TypeOf((*MockMerchantService)(nil).CreateMerchant), arg0, arg1, arg2)
}

// CreateMerchantItem mocks base method.
func (m *MockMerchantService) CreateMerchantItem(arg0 context.Context, arg1 merchant.CreateMerchantItemReq, arg2 *merchant.CreateMerchantItemRes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMerchantItem", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMerchantItem indicates an expected call of CreateMerchantItem.
func (mr *MockMerchantServiceMockRecorder) CreateMerchantItem(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMerchantItem", reflect.TypeOf((*MockMerchantService)(nil).CreateMerchantItem), arg0, arg1, arg2)
}

// FindMerchantByFilter mocks base method.
func (m *MockMerchantService) FindMerchantByFilter(arg0 context.Context, arg1 merchant.FindMerchantReq, arg2 *merchant.FindMerchantRes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMerchantByFilter", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindMerchantByFilter indicates an expected call of FindMerchantByFilter.
func (mr *MockMerchantServiceMockRecorder) FindMerchantByFilter(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMerchantByFilter", reflect.TypeOf((*MockMerchantService)(nil).FindMerchantByFilter), arg0, arg1, arg2)
}

// FindMerchantItemList mocks base method.
func (m *MockMerchantService) FindMerchantItemList(arg0 context.Context, arg1 merchant.MerchantItemListReq, arg2 *merchant.MerchantItemListRes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMerchantItemList", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindMerchantItemList indicates an expected call of FindMerchantItemList.
func (mr *MockMerchantServiceMockRecorder) FindMerchantItemList(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMerchantItemList", reflect.TypeOf((*MockMerchantService)(nil).FindMerchantItemList), arg0, arg1, arg2)
}
