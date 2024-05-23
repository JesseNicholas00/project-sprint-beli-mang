package auth

import (
	"errors"
	"net/http"

	"github.com/KerakTelor86/GoBoiler/services/auth"
	"github.com/KerakTelor86/GoBoiler/utils/errorutil"
	"github.com/KerakTelor86/GoBoiler/utils/request"
	"github.com/labstack/echo/v4"
)

func (ctrl *authController) loginStaff(c echo.Context) error {
	var req auth.LoginStaffReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	var res auth.LoginStaffRes
	if err := ctrl.service.LoginStaff(
		c.Request().Context(),
		req,
		&res,
	); err != nil {
		switch {
		case errors.Is(err, auth.ErrUserNotFound):
			return echo.NewHTTPError(http.StatusNotFound, echo.Map{
				"message": "user not found",
			})

		case errors.Is(err, auth.ErrInvalidCredentials):
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"message": "wrong password",
			})

		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User logged in successfully",
		"data":    res,
	})
}
