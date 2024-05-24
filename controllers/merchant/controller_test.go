package merchant

import (
	"github.com/JesseNicholas00/BeliMang/controllers/merchant/mocks"
	"github.com/JesseNicholas00/BeliMang/middlewares"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

//go:generate mockgen -destination mocks/mock_service.go -package mocks github.com/JesseNicholas00/BeliMang/services/merchant MerchantService

func NewControllerWithMockedService(t *testing.T) (
	mockCtrl *gomock.Controller,
	controller *merchantController,
	mockedService *mocks.MockMerchantService,
) {
	mockCtrl = gomock.NewController(t)
	mockedService = mocks.NewMockMerchantService(mockCtrl)
	controller = NewMerchantController(
		mockedService,
		middlewares.NewNoopMiddleware(),
	).(*merchantController)
	return
}

func TestRegister(t *testing.T) {
	mockCtrl, controller, _ := NewControllerWithMockedService(t)
	defer mockCtrl.Finish()

	Convey("When registering methods with an echo instance", t, func() {
		e := echo.New()
		err := controller.Register(e)
		Convey("Should not return error", func() {
			So(err, ShouldBeNil)
		})
	})
}
