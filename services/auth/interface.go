package auth

import "context"

type AuthService interface {
	RegisterUser(
		ctx context.Context,
		req RegisterUserReq,
		res *RegisterUserRes,
	) error
	LoginUser(
		ctx context.Context,
		req LoginUserReq,
		res *LoginUserRes,
	) error
	GetSessionFromToken(
		ctx context.Context,
		req GetSessionFromTokenReq,
		res *GetSessionFromTokenRes,
	) error
}
