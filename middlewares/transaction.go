package middlewares

import (
	"github.com/KerakTelor86/GoBoiler/utils/ctxrizz"
	"github.com/KerakTelor86/GoBoiler/utils/errorutil"
	"github.com/KerakTelor86/GoBoiler/utils/transaction"
	"github.com/labstack/echo/v4"
)

type txMiddleware struct {
	dbRizzer ctxrizz.DbContextRizzer
}

func (mw *txMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, txSession, err := mw.dbRizzer.AppendTx(c.Request().Context())
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		c.SetRequest(c.Request().WithContext(ctx))

		return transaction.RunWithAutoCommit(
			&txSession,
			func() error {
				return next(c)
			},
		)
	}
}

func NewWithTxMiddleware(dbRizzer ctxrizz.DbContextRizzer) Middleware {
	return &txMiddleware{
		dbRizzer: dbRizzer,
	}
}
