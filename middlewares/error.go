package middlewares

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/KerakTelor86/GoBoiler/utils/logging"
	"github.com/labstack/echo/v4"
)

type loggingErrorHandlerMiddleware struct {
}

var errorHandlerLogger = logging.GetLogger(
	"unhandledError",
)

func (*loggingErrorHandlerMiddleware) Process(
	next echo.HandlerFunc,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}

		// most errors should return here
		if err, ok := err.(*echo.HTTPError); ok {
			return c.JSON(err.Code, err.Message)
		}

		// purposefully checked after HTTP errors for speed
		if errors.Is(err, context.Canceled) {
			return c.NoContent(http.StatusNoContent)
		}

		if errors.Is(err, context.DeadlineExceeded) {
			return c.NoContent(http.StatusGatewayTimeout)
		}

		errorHandlerLogger.Error(
			"internal server error",
			"trace",
			fmt.Errorf("\n%w", err),
		)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}
}

// # Make sure this is the first middleware in the stack!
func NewLoggingErrorHandlerMiddleware() Middleware {
	return &loggingErrorHandlerMiddleware{}
}
