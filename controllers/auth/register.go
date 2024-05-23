package auth

import (
	"errors"
	"net/http"

	"github.com/KerakTelor86/GoBoiler/services/auth"
	"github.com/KerakTelor86/GoBoiler/utils/errorutil"
	"github.com/KerakTelor86/GoBoiler/utils/request"
	"github.com/labstack/echo/v4"
)

func (ctrl *authController) registerStaff(c echo.Context) error {
	var req auth.RegisterStaffReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	var res auth.RegisterStaffRes
	if err := ctrl.service.RegisterStaff(
		c.Request().Context(),
		req,
		&res,
	); err != nil {
		switch {
		case errors.Is(err, auth.ErrPhoneNumberAlreadyRegistered):
			return echo.NewHTTPError(http.StatusConflict, echo.Map{
				"message": "user already exists",
			})

		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "User registered successfully",
		"data":    res,
	})
}
