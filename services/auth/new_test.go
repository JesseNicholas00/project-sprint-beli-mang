package auth

import (
	"testing"

	"github.com/JesseNicholas00/BeliMang/services/auth/mocks"
	"github.com/JesseNicholas00/BeliMang/utils/ctxrizz"
	gomock "github.com/golang/mock/gomock"
)

//go:generate mockgen -destination mocks/mock_repo.go -package mocks github.com/JesseNicholas00/BeliMang/repos/auth AuthRepository

func NewWithMockedRepo(
	t *testing.T,
) (
	mockCtrl *gomock.Controller,
	service *authServiceImpl,
	mockedRepo *mocks.MockAuthRepository,
) {
	mockCtrl = gomock.NewController(t)
	mockedRepo = mocks.NewMockAuthRepository(mockCtrl)
	mockedDbRizzer := ctxrizz.NewDbContextNoopRizzer()
	service = NewAuthService(mockedRepo, mockedDbRizzer, "testKey", 8).(*authServiceImpl)
	return
}
