package auth

import (
	"context"
	"errors"

	"github.com/KerakTelor86/GoBoiler/utils/errorutil"
	"github.com/golang-jwt/jwt/v4"
)

func (svc *authServiceImpl) GetSessionFromToken(
	ctx context.Context,
	req GetSessionFromTokenReq,
	res *GetSessionFromTokenRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	claims := jwtClaims{}
	_, err := jwt.ParseWithClaims(
		req.AccessToken,
		&claims,
		func(t *jwt.Token) (interface{}, error) {
			return svc.jwtSecret, nil
		},
	)

	if err != nil {
		switch {

		case errors.Is(err, jwt.ErrTokenMalformed) ||
			errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return ErrTokenInvalid

		case errors.Is(err, jwt.ErrTokenExpired) ||
			errors.Is(err, jwt.ErrTokenNotValidYet):
			return ErrTokenExpired

		default:
			return errorutil.AddCurrentContext(err)
		}
	}

	*res = GetSessionFromTokenRes{
		UserId:      claims.Data.UserId,
		Name:        claims.Data.Name,
		PhoneNumber: claims.Data.PhoneNumber,
	}

	return nil
}
