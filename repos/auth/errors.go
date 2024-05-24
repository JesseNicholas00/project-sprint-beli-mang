package auth

import "errors"

var ErrUsernameNotFound = errors.New(
	"authRepository: no such username found",
)

var ErrEmailAndIsAdminNotFound = errors.New(
	"authRepository: no such email and is_admin found",
)
