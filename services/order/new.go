package order

import (
	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	"github.com/JesseNicholas00/BeliMang/repos/order"
	"github.com/JesseNicholas00/BeliMang/utils/ctxrizz"
)

type orderServiceImpl struct {
	repo         order.OrderRepository
	merchantRepo merchant.MerchantRepository
	dbRizzer     ctxrizz.DbContextRizzer
}

func NewOrderService(
	repo order.OrderRepository,
	merchantRepo merchant.MerchantRepository,
	dbRizzer ctxrizz.DbContextRizzer,
) OrderService {
	return &orderServiceImpl{
		repo:         repo,
		merchantRepo: merchantRepo,
		dbRizzer:     dbRizzer,
	}
}
