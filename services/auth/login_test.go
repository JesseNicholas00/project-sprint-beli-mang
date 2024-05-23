package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/KerakTelor86/GoBoiler/repos/auth"
	gomock "github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginStaff(t *testing.T) {
	Convey("When logging in as staff", t, func() {
		mockCtrl, service, mockedRepo := NewWithMockedRepo(t)
		defer mockCtrl.Finish()

		req := LoginStaffReq{
			PhoneNumber: "+6281234567890",
			Password:    "password",
		}
		reqWrong := LoginStaffReq{
			PhoneNumber: req.PhoneNumber,
			Password:    "epic bruh moment",
		}

		cryptedPw, err := bcrypt.GenerateFromPassword(
			[]byte(req.Password),
			service.bcryptCost,
		)
		So(err, ShouldBeNil)

		repoRes := auth.Staff{
			Id:        "bread",
			Name:      "john bread",
			Phone:     req.PhoneNumber,
			Password:  string(cryptedPw),
			CreatedAt: "now",
			UpdatedAt: "now",
		}

		Convey("If the phone number is not registered", func() {
			mockedRepo.EXPECT().
				FindStaffByPhone(gomock.Any(), req.PhoneNumber).
				Return(auth.Staff{}, auth.ErrPhoneNumberNotFound).
				Times(1)

			res := LoginStaffRes{}
			err := service.LoginStaff(context.TODO(), req, &res)
			Convey("Should return ErrUserNotFound", func() {
				So(
					errors.Is(err, ErrUserNotFound),
					ShouldBeTrue,
				)
			})
		})

		Convey("If the phone number is registered", func() {
			mockedRepo.EXPECT().
				FindStaffByPhone(gomock.Any(), req.PhoneNumber).
				Return(repoRes, nil).
				Times(1)

			Convey(
				"And the password is incorrect",
				func() {
					res := LoginStaffRes{}
					err := service.LoginStaff(context.TODO(), reqWrong, &res)

					Convey("Should return ErrInvalidCredentials", func() {
						So(errors.Is(err, ErrInvalidCredentials), ShouldBeTrue)
					})
				},
			)

			Convey(
				"And the password is correct",
				func() {
					res := LoginStaffRes{}
					err := service.LoginStaff(context.TODO(), req, &res)

					Convey(
						"Should return nil and write the correct result to res",
						func() {
							So(err, ShouldBeNil)
							So(res.Name, ShouldEqual, repoRes.Name)
							So(res.PhoneNumber, ShouldEqual, repoRes.Phone)
							So(res.UserId, ShouldEqual, repoRes.Id)
						},
					)
				},
			)
		})
	})
}
