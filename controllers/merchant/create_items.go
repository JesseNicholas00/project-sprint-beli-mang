package merchant

import (
	"errors"
	"net/http"

	mi "github.com/JesseNicholas00/BeliMang/services/merchantitem"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/request"
	"github.com/labstack/echo/v4"
)

func (mc *merchantController) createMerchantItems(c echo.Context) error {
	var req mi.CreateMerchantItemReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	var res mi.CreateMerchantItemRes
	err := mc.miService.CreateMerchantItem(c.Request().Context(), req, &res)
	if err != nil {
		if errors.Is(err, mi.ErrMerchantNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return errorutil.AddCurrentContext(err)
	}

	return c.JSON(http.StatusCreated, res)
}
