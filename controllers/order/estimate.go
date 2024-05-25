package order

import (
	"errors"
	"github.com/JesseNicholas00/BeliMang/services/auth"
	"github.com/JesseNicholas00/BeliMang/services/order"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/request"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ctrl *orderController) estimateOrder(c echo.Context) error {
	var req order.EstimateOrderReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	startPointCnt := 0
	for _, merchantOrder := range req.Orders {
		if *merchantOrder.IsStartingPoint {
			startPointCnt++
		}
	}

	if startPointCnt != 1 {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": "must have exactly one starting point",
		})
	}

	req.UserId = c.Get("session").(auth.GetSessionFromTokenRes).UserId

	var res order.EstimateOrderRes
	err := ctrl.service.EstimateOrder(c.Request().Context(), req, &res)

	if err != nil {
		switch {
		case errors.Is(err, order.ErrTooFar):
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"message": "minimum distance traveled is too far",
			})
		case errors.Is(err, order.ErrMerchantNotFound):
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"message": "merchant not found",
			})
		case errors.Is(err, order.ErrItemNotFound):
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"message": "item not found",
			})

		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	return c.JSON(http.StatusOK, res)
}
