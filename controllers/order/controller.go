package order

import (
	"github.com/JesseNicholas00/BeliMang/controllers"
	"github.com/JesseNicholas00/BeliMang/middlewares"
	"github.com/JesseNicholas00/BeliMang/services/order"
	"github.com/labstack/echo/v4"
)

type orderController struct {
	authMw  middlewares.Middleware
	service order.OrderService
}

func (ctrl *orderController) Register(server *echo.Echo) error {
	server.POST("/users/estimate", ctrl.estimateOrder, ctrl.authMw.Process)
	server.GET("/users/orders", ctrl.orderHistory, ctrl.authMw.Process)
	server.POST("/users/orders", ctrl.createOrder, ctrl.authMw.Process)
	return nil
}

func NewOrderController(
	service order.OrderService,
	authMw middlewares.Middleware,
) controllers.Controller {
	return &orderController{
		service: service,
		authMw:  authMw,
	}
}
