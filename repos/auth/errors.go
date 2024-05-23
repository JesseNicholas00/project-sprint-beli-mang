package auth

import "errors"

var ErrUsernameNotFound = errors.New(
	"authRepository: no such username found",
)
