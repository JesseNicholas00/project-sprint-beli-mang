package merchant

import (
	"github.com/JesseNicholas00/BeliMang/controllers"
	"github.com/JesseNicholas00/BeliMang/middlewares"
	"github.com/JesseNicholas00/BeliMang/services/merchant"
	mi "github.com/JesseNicholas00/BeliMang/services/merchantitem"
	"github.com/labstack/echo/v4"
)

type merchantController struct {
	authMw    middlewares.Middleware
	service   merchant.MerchantService
	miService mi.MerchantItemService
}

func (mc *merchantController) Register(server *echo.Echo) error {
	server.POST("/admin/merchants", mc.createMerchant, mc.authMw.Process)
	server.GET("/admin/merchants", mc.adminList, mc.authMw.Process)
	server.POST("/admin/merchants/:merchantId/items", mc.createMerchantItems, mc.authMw.Process)
	return nil
}

func NewMerchantController(
	service merchant.MerchantService,
	miService mi.MerchantItemService,
	authMw middlewares.Middleware,
) controllers.Controller {
	return &merchantController{
		service:   service,
		miService: miService,
		authMw:    authMw,
	}
}
