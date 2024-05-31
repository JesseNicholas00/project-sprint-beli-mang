package unittesting

import (
	"github.com/JesseNicholas00/BeliMang/types/role"
	"github.com/labstack/echo/v4"
)

func CallController(
	ctx echo.Context,
	ctrl func(echo.Context) error,
) {
	err := ctrl(ctx)
	ctx.Echo().HTTPErrorHandler(err, ctx)
}

func CallControllerWithRole(
	ctx echo.Context,
	ctrl func(echo.Context, role.Role) error,
	role role.Role,
) {
	err := ctrl(ctx, role)
	ctx.Echo().HTTPErrorHandler(err, ctx)
}
