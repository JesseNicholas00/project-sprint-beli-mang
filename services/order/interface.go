package order

import "context"

type OrderService interface {
	EstimateOrder(
		ctx context.Context,
		req EstimateOrderReq,
		res *EstimateOrderRes,
	) error
	OrderHistory(
		ctx context.Context,
		req OrderHistoryReq,
		res *OrderHistoryRes,
	) error
}
