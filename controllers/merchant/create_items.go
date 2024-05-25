package merchant

import (
	"errors"
	"net/http"

	"github.com/JesseNicholas00/BeliMang/services/merchant"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/request"
	"github.com/labstack/echo/v4"
)

func (mc *merchantController) createMerchantItems(c echo.Context) error {
	var req merchant.CreateMerchantItemReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	var res merchant.CreateMerchantItemRes
	err := mc.service.CreateMerchantItem(c.Request().Context(), req, &res)
	if err != nil {
		if errors.Is(err, merchant.ErrMerchantNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return errorutil.AddCurrentContext(err)
	}

	return c.JSON(http.StatusCreated, res)
}
