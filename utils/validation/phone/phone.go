package phone

import (
	"strings"

	"github.com/KerakTelor86/GoBoiler/utils/helper"
	"github.com/go-playground/validator/v10"
)

func validatePhoneNumberImpl(phoneNum string) bool {
	return helper.IsBetween(len(phoneNum), 10, 15) &&
		strings.HasPrefix(phoneNum, "+62")
}

func ValidatePhoneNumber(fl validator.FieldLevel) bool {
	return validatePhoneNumberImpl(fl.Field().String())
}
