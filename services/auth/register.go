package auth

import (
	"context"
	"errors"

	"github.com/JesseNicholas00/BeliMang/repos/auth"
	"github.com/JesseNicholas00/BeliMang/types/role"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/transaction"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (svc *authServiceImpl) RegisterUser(
	ctx context.Context,
	req RegisterUserReq,
	res *RegisterUserRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := svc.dbRizzer.GetOrAppendTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return transaction.RunWithAutoCommit(&sess, func() error {
		_, err := svc.repo.FindUserByUsername(ctx, req.Username)

		if err == nil {
			// duplicate username
			return ErrUsernameAlreadyRegistered
		}

		if !errors.Is(err, auth.ErrUsernameNotFound) {
			// unexpected kind of error
			return errorutil.AddCurrentContext(err)
		}

		_, err = svc.repo.FindUserByEmailAndIsAdmin(ctx, req.Email, req.Role != role.Admin)

		if err == nil {
			return ErrEmailAlreadyRegistered
		}

		if !errors.Is(err, auth.ErrEmailAndIsAdminNotFound) {
			// unexpected kind of error
			return errorutil.AddCurrentContext(err)
		}

		cryptedPw, err := bcrypt.GenerateFromPassword(
			[]byte(req.Password),
			svc.bcryptCost,
		)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		repoRes, err := svc.repo.CreateUser(ctx, auth.User{
			Id:       uuid.New().String(),
			Username: req.Username,
			Email:    req.Email,
			Password: string(cryptedPw),
			IsAdmin:  role.ToBoolean(req.Role),
		})
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		token, err := svc.generateToken(repoRes)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		*res = RegisterUserRes{
			AccessToken: token,
		}
		return nil
	})
}
