package auth

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/JesseNicholas00/BeliMang/repos/auth"
	"github.com/golang-jwt/jwt/v4"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetSessionFromToken(t *testing.T) {
	Convey("When getting session from token", t, func() {
		mockCtrl, service, _ := NewWithMockedRepo(t)
		defer mockCtrl.Finish()

		user := auth.User{
			Id:       "bread",
			Username: "firstname lastname",
			IsAdmin:  true,
		}

		Convey("And the token is valid", func() {
			validToken, err := service.generateToken(user)
			So(err, ShouldBeNil)

			Convey("Should return the correct user data", func() {
				req := GetSessionFromTokenReq{
					AccessToken: validToken,
				}
				res := GetSessionFromTokenRes{}
				err := service.GetSessionFromToken(context.TODO(), req, &res)

				So(err, ShouldBeNil)
				So(res.UserId, ShouldEqual, user.Id)
				So(res.IsAdmin, ShouldEqual, user.IsAdmin)
			})
		})

		Convey("And the token is invalid", func() {
			token := "not even a token lol"

			Convey("Should return ErrTokenInvalid", func() {
				req := GetSessionFromTokenReq{
					AccessToken: token,
				}
				res := GetSessionFromTokenRes{}
				err := service.GetSessionFromToken(context.TODO(), req, &res)

				So(errors.Is(err, ErrTokenInvalid), ShouldBeTrue)
			})
		})

		Convey("And the token is expired", func() {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(
						time.Now().Add(-8 * time.Hour),
					),
				},
				Data: jwtSubClaims{
					UserId:  user.Id,
					IsAdmin: true,
				},
			})
			res, err := token.SignedString(service.jwtSecret)
			So(err, ShouldBeNil)

			Convey("Should return ErrTokenExpired", func() {
				req := GetSessionFromTokenReq{
					AccessToken: res,
				}
				res := GetSessionFromTokenRes{}
				err := service.GetSessionFromToken(context.TODO(), req, &res)

				So(errors.Is(err, ErrTokenExpired), ShouldBeTrue)
			})
		})
	})
}
