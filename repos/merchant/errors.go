package merchant

import "errors"

var ErrMerchantNotFound = errors.New(
	"merchantRepository: merchant not found",
)

var ErrLocationNotGiven = errors.New(
	"merchantRepository: lat and long must be passed",
)
