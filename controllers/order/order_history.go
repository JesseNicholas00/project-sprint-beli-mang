package order

import (
	"net/http"

	"github.com/JesseNicholas00/BeliMang/services/auth"
	"github.com/JesseNicholas00/BeliMang/services/order"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/helper"
	"github.com/JesseNicholas00/BeliMang/utils/request"
	"github.com/labstack/echo/v4"
)

func (ctrl *orderController) orderHistory(c echo.Context) error {
	var req order.OrderHistoryReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	if req.Limit == nil {
		req.Limit = helper.ToPointer(5)
	}
	if req.Offset == nil {
		req.Offset = helper.ToPointer(0)
	}

	req.UserId = c.Get("session").(auth.GetSessionFromTokenRes).UserId

	var res order.OrderHistoryRes
	err := ctrl.service.OrderHistory(c.Request().Context(), req, &res)

	if err != nil {
		switch {
		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	return c.JSON(http.StatusOK, res.Result)
}
