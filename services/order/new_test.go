package order

import (
	"testing"

	"github.com/JesseNicholas00/BeliMang/services/order/mocks"
	"github.com/JesseNicholas00/BeliMang/utils/ctxrizz"
	"github.com/golang/mock/gomock"
)

//go:generate mockgen -destination mocks/mock_order_repo.go -package mocks github.com/JesseNicholas00/BeliMang/repos/order OrderRepository
//go:generate mockgen -destination mocks/mock_merchant_repo.go -package mocks github.com/JesseNicholas00/BeliMang/repos/merchant MerchantRepository

func NewWithMockedRepo(
	t *testing.T,
) (
	mockCtrl *gomock.Controller,
	service *orderServiceImpl,
	mockedRepo *mocks.MockOrderRepository,
	mockedMerchantRepo *mocks.MockMerchantRepository,
) {
	mockCtrl = gomock.NewController(t)
	mockedRepo = mocks.NewMockOrderRepository(mockCtrl)
	mockedMerchantRepo = mocks.NewMockMerchantRepository(mockCtrl)
	noopRizzer := ctxrizz.NewDbContextNoopRizzer()
	service = NewOrderService(
		mockedRepo,
		mockedMerchantRepo,
		noopRizzer,
	).(*orderServiceImpl)
	return
}
