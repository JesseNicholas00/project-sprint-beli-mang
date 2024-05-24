package merchant

import (
	"github.com/JesseNicholas00/BeliMang/services/merchant"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/request"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (mc *merchantController) createMerchant(c echo.Context) error {
	var req merchant.CreateMerchantReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	var res merchant.CreateMerchantRes
	err := mc.service.CreateMerchant(c.Request().Context(), req, &res)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return c.JSON(http.StatusCreated, res)
}
