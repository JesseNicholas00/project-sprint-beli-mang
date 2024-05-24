package image

import (
	"github.com/JesseNicholas00/BeliMang/controllers"
	"github.com/JesseNicholas00/BeliMang/middlewares"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/labstack/echo/v4"
)

type imageController struct {
	service *manager.Uploader
	bucket  string
	authMw  middlewares.Middleware
}

func (ctrl *imageController) Register(server *echo.Echo) error {
	g := server.Group("/v1/image")

	g.POST("", ctrl.uploadImage, ctrl.authMw.Process)

	return nil
}

func NewImageController(
	service *manager.Uploader,
	bucket string,
	authMw middlewares.Middleware,
) controllers.Controller {
	return &imageController{
		service: service,
		bucket:  bucket,
		authMw:  authMw,
	}
}
