package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/JesseNicholas00/BeliMang/repos/auth"
	gomock "github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginUser(t *testing.T) {
	Convey("When logging in as user", t, func() {
		mockCtrl, service, mockedRepo := NewWithMockedRepo(t)
		defer mockCtrl.Finish()

		req := LoginUserReq{
			Username: "+6281234567890",
			Password: "password",
		}
		reqWrong := LoginUserReq{
			Username: req.Username,
			Password: "epic bruh moment",
		}

		cryptedPw, err := bcrypt.GenerateFromPassword(
			[]byte(req.Password),
			service.bcryptCost,
		)
		So(err, ShouldBeNil)

		repoRes := auth.User{
			Id:       "bread",
			Username: "jessenicholas",
			Password: string(cryptedPw),
			IsAdmin:  true,
		}

		Convey("If the username is not registered", func() {
			mockedRepo.EXPECT().
				FindUserByUsername(gomock.Any(), req.Username).
				Return(auth.User{}, auth.ErrUsernameNotFound).
				Times(1)

			res := LoginUserRes{}
			err := service.LoginUser(context.TODO(), req, &res)
			Convey("Should return ErrUserNotFound", func() {
				So(
					errors.Is(err, ErrUserNotFound),
					ShouldBeTrue,
				)
			})
		})

		Convey("If the username is registered", func() {
			mockedRepo.EXPECT().
				FindUserByUsername(gomock.Any(), req.Username).
				Return(repoRes, nil).
				Times(1)

			Convey(
				"And the password is incorrect",
				func() {
					res := LoginUserRes{}
					err := service.LoginUser(context.TODO(), reqWrong, &res)

					Convey("Should return ErrInvalidCredentials", func() {
						So(errors.Is(err, ErrInvalidCredentials), ShouldBeTrue)
					})
				},
			)

			Convey(
				"And the password is correct",
				func() {
					res := LoginUserRes{}
					err := service.LoginUser(context.TODO(), req, &res)

					Convey(
						"Should return nil and write the correct result to res",
						func() {
							So(err, ShouldBeNil)
							So(res.AccessToken, ShouldNotBeNil)
						},
					)
				},
			)
		})
	})
}
