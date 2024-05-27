package merchant

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/JesseNicholas00/BeliMang/services/merchant"
	"github.com/JesseNicholas00/BeliMang/types/location"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/helper"
	"github.com/JesseNicholas00/BeliMang/utils/request"
	"github.com/labstack/echo/v4"
)

func (mc *merchantController) findByFilters(c echo.Context) error {
	var req merchant.FindMerchantReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	if req.Limit == nil {
		req.Limit = helper.ToPointer(5)
	}

	fanum_tax := strings.Split(req.LatLongSeparatedByCommaIdkWhyItsLikeThatButOk, ",")

	if len(fanum_tax) == 2 {
		fanum, err_fanum := strconv.ParseFloat(fanum_tax[0], 64)
		if err_fanum != nil {
			return err_fanum
		}
		tax, err_tax := strconv.ParseFloat(fanum_tax[1], 64)
		if err_tax != nil {
			return err_tax
		}
		req.Location = location.GyattLocation{
			Latitude:  &fanum,
			Longitude: &tax,
		}
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "lat and long is not valid, please provide the gyatt long lat ok",
		})
	}

	var res merchant.FindMerchantRes
	err := mc.service.FindMerchantByFilter(c.Request().Context(), req, &res)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return c.JSON(http.StatusOK, res)
}
