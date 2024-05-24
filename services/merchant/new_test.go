package merchant

import (
	"testing"

	"github.com/JesseNicholas00/BeliMang/services/merchant/mocks"
	"github.com/JesseNicholas00/BeliMang/utils/ctxrizz"
	"github.com/golang/mock/gomock"
)

//go:generate mockgen -destination mocks/mock_repo.go -package mocks github.com/JesseNicholas00/BeliMang/repos/merchant MerchantRepository

func NewWithMockedRepo(
	t *testing.T,
) (
	mockCtrl *gomock.Controller,
	service *merchantServiceImpl,
	mockedRepo *mocks.MockMerchantRepository,
) {
	mockCtrl = gomock.NewController(t)
	mockedRepo = mocks.NewMockMerchantRepository(mockCtrl)
	noopRizzer := ctxrizz.NewDbContextNoopRizzer()
	service = NewMerchantService(mockedRepo, noopRizzer).(*merchantServiceImpl)
	return
}
