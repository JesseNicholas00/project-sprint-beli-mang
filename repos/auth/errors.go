package auth

import "errors"

var ErrPhoneNumberNotFound = errors.New(
	"authRepository: no such phone number found",
)
