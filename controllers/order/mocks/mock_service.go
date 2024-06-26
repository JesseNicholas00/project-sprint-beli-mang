// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/JesseNicholas00/BeliMang/services/order (interfaces: OrderService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	order "github.com/JesseNicholas00/BeliMang/services/order"
	gomock "github.com/golang/mock/gomock"
)

// MockOrderService is a mock of OrderService interface.
type MockOrderService struct {
	ctrl     *gomock.Controller
	recorder *MockOrderServiceMockRecorder
}

// MockOrderServiceMockRecorder is the mock recorder for MockOrderService.
type MockOrderServiceMockRecorder struct {
	mock *MockOrderService
}

// NewMockOrderService creates a new mock instance.
func NewMockOrderService(ctrl *gomock.Controller) *MockOrderService {
	mock := &MockOrderService{ctrl: ctrl}
	mock.recorder = &MockOrderServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderService) EXPECT() *MockOrderServiceMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockOrderService) CreateOrder(arg0 context.Context, arg1 order.CreateOrderReq, arg2 *order.CreateOrderRes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderServiceMockRecorder) CreateOrder(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderService)(nil).CreateOrder), arg0, arg1, arg2)
}

// EstimateOrder mocks base method.
func (m *MockOrderService) EstimateOrder(arg0 context.Context, arg1 order.EstimateOrderReq, arg2 *order.EstimateOrderRes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstimateOrder", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// EstimateOrder indicates an expected call of EstimateOrder.
func (mr *MockOrderServiceMockRecorder) EstimateOrder(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateOrder", reflect.TypeOf((*MockOrderService)(nil).EstimateOrder), arg0, arg1, arg2)
}

// OrderHistory mocks base method.
func (m *MockOrderService) OrderHistory(arg0 context.Context, arg1 order.OrderHistoryReq, arg2 *order.OrderHistoryRes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrderHistory", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// OrderHistory indicates an expected call of OrderHistory.
func (mr *MockOrderServiceMockRecorder) OrderHistory(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderHistory", reflect.TypeOf((*MockOrderService)(nil).OrderHistory), arg0, arg1, arg2)
}
