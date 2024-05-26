package merchant

import (
	"errors"
	"net/http"

	"github.com/JesseNicholas00/BeliMang/services/merchant"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/request"
	"github.com/labstack/echo/v4"
)

func (mc *merchantController) findByFilters(c echo.Context) error {
	var req merchant.FindMerchantReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	var res merchant.FindMerchantRes
	err := mc.service.FindMerchantByFilter(c.Request().Context(), req, &res)
	if err != nil {
		switch {
		case errors.Is(err, merchant.ErrLatLangNotValid):
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"message": "lat and long is not valid, please provide the gyatt long lat ok",
			})

		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	return c.JSON(http.StatusOK, res)
}
