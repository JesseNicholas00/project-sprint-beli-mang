package auth

import "context"

type AuthService interface {
	RegisterStaff(
		ctx context.Context,
		req RegisterStaffReq,
		res *RegisterStaffRes,
	) error
	LoginStaff(
		ctx context.Context,
		req LoginStaffReq,
		res *LoginStaffRes,
	) error
	GetSessionFromToken(
		ctx context.Context,
		req GetSessionFromTokenReq,
		res *GetSessionFromTokenRes,
	) error
}
