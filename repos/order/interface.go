package order

import "context"

type OrderRepository interface {
	CreateEstimate(ctx context.Context, est Estimate) error
	FindEstimateById(ctx context.Context, id string) (Estimate, error)
	CreateOrder(ctx context.Context, order Order) error
	ListOrderSummary(ctx context.Context, filter OrderSummaryListFilter) ([]OrderSummaryView, error)
}
