package order_test

import (
	"testing"

	"github.com/JesseNicholas00/BeliMang/repos/order"
	"github.com/JesseNicholas00/BeliMang/utils/ctxrizz"
	"github.com/JesseNicholas00/BeliMang/utils/unittesting"
)

func NewWithTestDatabase(t *testing.T) order.OrderRepository {
	db := unittesting.SetupTestDatabase("../../migrations", t)
	return order.NewOrderRepository(ctxrizz.NewDbContextRizzer(db))
}
