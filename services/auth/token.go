package auth

import (
	"time"

	"github.com/KerakTelor86/GoBoiler/repos/auth"
	"github.com/golang-jwt/jwt/v4"
)

func (svc *authServiceImpl) generateToken(
	staff auth.Staff,
) (res string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
		},
		Data: jwtSubClaims{
			UserId:      staff.Id,
			PhoneNumber: staff.Phone,
			Name:        staff.Name,
		},
	})
	res, err = token.SignedString(svc.jwtSecret)
	return
}
