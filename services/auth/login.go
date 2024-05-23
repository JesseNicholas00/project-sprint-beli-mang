package auth

import (
	"context"

	"github.com/JesseNicholas00/BeliMang/repos/auth"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (svc *authServiceImpl) LoginStaff(
	ctx context.Context,
	req LoginStaffReq,
	res *LoginStaffRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	staff, err := svc.repo.FindStaffByPhone(ctx, req.PhoneNumber)

	if err != nil {
		if errors.Is(err, auth.ErrPhoneNumberNotFound) {
			return ErrUserNotFound
		}

		return errorutil.AddCurrentContext(err)
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(staff.Password),
		[]byte(req.Password),
	)
	if err != nil {
		return ErrInvalidCredentials
	}

	token, err := svc.generateToken(staff)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	*res = LoginStaffRes{
		UserId:      staff.Id,
		PhoneNumber: staff.Phone,
		Name:        staff.Name,
		AccessToken: token,
	}

	return nil
}
