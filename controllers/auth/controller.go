package auth

import (
	"github.com/JesseNicholas00/BeliMang/controllers"
	"github.com/JesseNicholas00/BeliMang/services/auth"
	"github.com/JesseNicholas00/BeliMang/types/role"
	"github.com/labstack/echo/v4"
)

type authController struct {
	service auth.AuthService
}

func (s *authController) Register(server *echo.Echo) error {
	server.POST("/admin/register", func(c echo.Context) error {
		return s.registerUser(c, role.Admin)
	})
	server.POST("/admin/login", func(c echo.Context) error {
		return s.loginUser(c, role.Admin)
	})
	server.POST("/user/register", func(c echo.Context) error {
		return s.registerUser(c, role.User)
	})
	server.POST("/user/login", func(c echo.Context) error {
		return s.loginUser(c, role.User)
	})
	return nil
}

func NewAuthController(
	service auth.AuthService,
) controllers.Controller {
	return &authController{
		service: service,
	}
}
