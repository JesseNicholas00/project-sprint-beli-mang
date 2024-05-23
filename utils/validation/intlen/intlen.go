package intlen

import (
	"strconv"

	"github.com/JesseNicholas00/BeliMang/utils/helper"
	"github.com/go-playground/validator/v10"
)

func validateIntLenImpl(x int64, len int) bool {
	return helper.HasLen(x, len)
}

func ValidateIntLen(fl validator.FieldLevel) bool {
	len, err := strconv.Atoi(fl.Param())
	if err != nil {
		return false
	}

	return validateIntLenImpl(fl.Field().Int(), len)
}
