package merchant

import (
	"github.com/JesseNicholas00/BeliMang/services/merchant/mocks"
	"github.com/golang/mock/gomock"
	"testing"
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
	service = NewMerchantService(mockedRepo).(*merchantServiceImpl)
	return
}
