package auth

import (
	"github.com/KerakTelor86/GoBoiler/controllers"
	"github.com/KerakTelor86/GoBoiler/services/auth"
	"github.com/labstack/echo/v4"
)

type authController struct {
	service auth.AuthService
}

func (s *authController) Register(server *echo.Echo) error {
	server.POST("/v1/staff/register", s.registerStaff)
	server.POST("/v1/staff/login", s.loginStaff)
	return nil
}

func NewAuthController(
	service auth.AuthService,
) controllers.Controller {
	return &authController{
		service: service,
	}
}
