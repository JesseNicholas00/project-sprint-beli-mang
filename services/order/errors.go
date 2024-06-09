package order

import "errors"

var (
	ErrTooFar = errors.New(
		"orderService: total distance too far (max 3km)",
	)
	ErrMerchantNotFound = errors.New("orderService: no such merchant found")
	ErrItemNotFound     = errors.New("orderService: no such item found")

	ErrEstimateNotFound = errors.New("orderService: no such estimate found")
)
