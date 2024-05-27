package merchant

import "errors"

var ErrMerchantNotFound = errors.New(
	"merchantService: merchant not found",
)

var ErrLatLangNotValid = errors.New(
	"merchantService: lat lang is not valid",
)
