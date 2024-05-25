package order

import "errors"

var (
	ErrOrderNotFound = errors.New(
		"orderRepository: no such order found",
	)
	ErrEstimateNotFound = errors.New(
		"orderRepository: no such estimate found",
	)
)
