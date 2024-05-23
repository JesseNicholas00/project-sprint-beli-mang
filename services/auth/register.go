package auth

import (
	"context"
	"errors"

	"github.com/KerakTelor86/GoBoiler/repos/auth"
	"github.com/KerakTelor86/GoBoiler/utils/errorutil"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (svc *authServiceImpl) RegisterStaff(
	ctx context.Context,
	req RegisterStaffReq,
	res *RegisterStaffRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	_, err := svc.repo.FindStaffByPhone(ctx, req.PhoneNumber)

	if err == nil {
		// duplicate phone number
		return ErrPhoneNumberAlreadyRegistered
	}

	if !errors.Is(err, auth.ErrPhoneNumberNotFound) {
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

	repoRes, err := svc.repo.CreateStaff(ctx, auth.Staff{
		Id:       uuid.New().String(),
		Name:     req.Name,
		Phone:    req.PhoneNumber,
		Password: string(cryptedPw),
	})
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	token, err := svc.generateToken(repoRes)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	*res = RegisterStaffRes{
		UserId:      repoRes.Id,
		PhoneNumber: repoRes.Phone,
		Name:        repoRes.Name,
		AccessToken: token,
	}

	return nil
}
