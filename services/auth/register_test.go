package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/JesseNicholas00/BeliMang/repos/auth"
	gomock "github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRegisterUser(t *testing.T) {
	Convey("When registering user", t, func() {
		mockCtrl, service, mockedRepo := NewWithMockedRepo(t)
		defer mockCtrl.Finish()

		req := RegisterUserReq{
			Username: "+6281234567890",
			Password: "password",
			Email:    "jn@gmail.com",
		}

		repoReq := auth.User{
			Id:       "bread",
			Username: "+6281234567890",
			Password: "password",
			Email:    "jn@gmail.com",
			IsAdmin:  true,
		}
		repoRes := auth.User{
			Id:       repoReq.Id,
			Username: repoReq.Username,
			Password: repoReq.Password,
			Email:    "jn@gmail.com",
			IsAdmin:  true,
		}

		Convey("If the username is already registered", func() {
			mockedRepo.EXPECT().
				FindUserByUsername(gomock.Any(), req.Username).
				Return(repoRes, nil).
				Times(1)

			res := RegisterUserRes{}
			err := service.RegisterUser(context.TODO(), req, &res)
			Convey("Should return ErrPhoneNumberAlreadyRegistered", func() {
				So(
					errors.Is(err, ErrUsernameAlreadyRegistered),
					ShouldBeTrue,
				)
			})
		})

		Convey("If the username is unique", func() {
			mockedRepo.EXPECT().
				FindUserByUsername(gomock.Any(), req.Username).
				Return(auth.User{}, auth.ErrUsernameNotFound).
				Times(1)
			mockedRepo.EXPECT().
				CreateUser(gomock.Any(), gomock.Any()).
				Do(func(_ context.Context, reqFromSvc auth.User) {
					So(reqFromSvc.Username, ShouldEqual, req.Username)
					So(reqFromSvc.Email, ShouldEqual, req.Email)
				}).
				Return(repoRes, nil).
				Times(1)

			res := RegisterUserRes{}
			err := service.RegisterUser(context.TODO(), req, &res)
			Convey(
				"Should return nil and write the correct result to res",
				func() {
					So(err, ShouldBeNil)
					So(res.AccessToken, ShouldNotBeNil)
				},
			)
		})
	})
}
