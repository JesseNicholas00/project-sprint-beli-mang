package order

import (
	"github.com/JesseNicholas00/BeliMang/controllers/order/mocks"
	"github.com/JesseNicholas00/BeliMang/middlewares"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

//go:generate mockgen -destination mocks/mock_service.go -package mocks github.com/JesseNicholas00/BeliMang/services/order OrderService

func NewControllerWithMockedService(t *testing.T) (
	mockCtrl *gomock.Controller,
	controller *orderController,
	mockedService *mocks.MockOrderService,
) {
	mockCtrl = gomock.NewController(t)
	mockedService = mocks.NewMockOrderService(mockCtrl)
	controller = NewOrderController(
		mockedService,
		middlewares.NewNoopMiddleware(),
	).(*orderController)
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
