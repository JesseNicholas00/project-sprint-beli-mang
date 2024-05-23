package auth

import (
	"github.com/JesseNicholas00/BeliMang/controllers"
	"github.com/JesseNicholas00/BeliMang/services/auth"
	"github.com/labstack/echo/v4"
)

type authController struct {
	service auth.AuthService
}

func (s *authController) Register(server *echo.Echo) error {
	server.POST("/:role/register", s.registerUser)
	server.POST("/:role/login", s.loginUser)
	return nil
}

func NewAuthController(
	service auth.AuthService,
) controllers.Controller {
	return &authController{
		service: service,
	}
}
