package order

import (
	"errors"
	"net/http"

	"github.com/JesseNicholas00/BeliMang/services/auth"
	"github.com/JesseNicholas00/BeliMang/services/order"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/request"
	"github.com/labstack/echo/v4"
)

func (ctrl *orderController) createOrder(c echo.Context) error {
	var req order.CreateOrderReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	req.UserId = c.Get("session").(auth.GetSessionFromTokenRes).UserId

	var res order.CreateOrderRes
	err := ctrl.service.CreateOrder(c.Request().Context(), req, &res)

	if err != nil {
		switch {
		case errors.Is(err, order.ErrEstimateNotFound):
			return echo.NewHTTPError(http.StatusNotFound, echo.Map{
				"message": "calculatedEstimateId not found",
			})

		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	return c.JSON(http.StatusCreated, res)
}
