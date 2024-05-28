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

func (ctrl *authController) loginUser(c echo.Context, role role.Role) error {
	var req auth.LoginUserReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}
	req.Role = role

	var res auth.LoginUserRes
	if err := ctrl.service.LoginUser(
		c.Request().Context(),
		req,
		&res,
	); err != nil {
		switch {
		case errors.Is(err, auth.ErrUserNotFound):
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
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

	return c.JSON(http.StatusOK, res)
}
