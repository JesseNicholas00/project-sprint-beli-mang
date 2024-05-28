package auth

import (
	"errors"
	"net/http"

	"github.com/JesseNicholas00/BeliMang/services/auth"
	"github.com/JesseNicholas00/BeliMang/types/role"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/request"
	"github.com/labstack/echo/v4"
)

func (ctrl *authController) registerUser(c echo.Context, role role.Role) error {
	var req auth.RegisterUserReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}
	req.Role = role

	var res auth.RegisterUserRes
	if err := ctrl.service.RegisterUser(
		c.Request().Context(),
		req,
		&res,
	); err != nil {
		switch {
		case errors.Is(err, auth.ErrUsernameAlreadyRegistered):
			return echo.NewHTTPError(http.StatusConflict, echo.Map{
				"message": "user already exists",
			})

		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	return c.JSON(http.StatusCreated, res)
}
