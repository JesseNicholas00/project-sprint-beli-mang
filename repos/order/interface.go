package order

import "context"

type OrderRepository interface {
	CreateEstimate(ctx context.Context, est Estimate) error
}
