package merchantitem

import "errors"

var ErrMerchantNotFound = errors.New(
	"merchantService: merchant not found",
)
