package merchant

import (
	"github.com/JesseNicholas00/BeliMang/controllers"
	"github.com/JesseNicholas00/BeliMang/middlewares"
	"github.com/JesseNicholas00/BeliMang/services/merchant"
	"github.com/labstack/echo/v4"
)

type merchantController struct {
	authMw  middlewares.Middleware
	service merchant.MerchantService
}

func (mc *merchantController) Register(server *echo.Echo) error {
	server.POST("/admin/merchants", mc.createMerchant, mc.authMw.Process)
	return nil
}

func NewMerchantController(
	service merchant.MerchantService,
	authMw middlewares.Middleware,
) controllers.Controller {
	return &merchantController{
		service: service,
		authMw:  authMw,
	}
}