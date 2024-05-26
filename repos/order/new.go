package order

import (
	"github.com/JesseNicholas00/BeliMang/utils/ctxrizz"
)

type orderRepositoryImpl struct {
	dbRizzer   ctxrizz.DbContextRizzer
	statements statements
}

func NewOrderRepository(dbRizzer ctxrizz.DbContextRizzer) OrderRepository {
	return &orderRepositoryImpl{
		dbRizzer:   dbRizzer,
		statements: prepareStatements(),
	}
}
