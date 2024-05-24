package merchant

import "errors"

var ErrMerchantNotFound = errors.New(
	"merchantRepository: merchant not found",
)
