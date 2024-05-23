package auth

import (
	"context"

	"github.com/JesseNicholas00/BeliMang/repos/auth"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (svc *authServiceImpl) LoginUser(
	ctx context.Context,
	req LoginUserReq,
	res *LoginUserRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	user, err := svc.repo.FindUserByUsername(ctx, req.Username)

	if err != nil {
		if errors.Is(err, auth.ErrUsernameNotFound) {
			return ErrUserNotFound
		}

		return errorutil.AddCurrentContext(err)
	}

	if !user.IsAdmin && req.Role == "admin" {
		return ErrUserNotFound
	}
	if user.IsAdmin && req.Role == "user" {
		return ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)
	if err != nil {
		return ErrInvalidCredentials
	}

	token, err := svc.generateToken(user)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	*res = LoginUserRes{
		AccessToken: token,
	}

	return nil
}
