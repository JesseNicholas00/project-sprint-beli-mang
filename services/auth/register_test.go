package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/KerakTelor86/GoBoiler/repos/auth"
	gomock "github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRegisterStaff(t *testing.T) {
	Convey("When registering staff", t, func() {
		mockCtrl, service, mockedRepo := NewWithMockedRepo(t)
		defer mockCtrl.Finish()

		req := RegisterStaffReq{
			PhoneNumber: "+6281234567890",
			Name:        "firstname lastname",
			Password:    "password",
		}

		repoReq := auth.Staff{
			Id:       "bread",
			Name:     req.Name,
			Phone:    req.PhoneNumber,
			Password: req.Password,
		}
		repoRes := auth.Staff{
			Id:        repoReq.Id,
			Name:      repoReq.Name,
			Phone:     repoReq.Phone,
			Password:  repoReq.Password,
			CreatedAt: "now",
			UpdatedAt: "now",
		}

		Convey("If the phone number is already registered", func() {
			mockedRepo.EXPECT().
				FindStaffByPhone(gomock.Any(), req.PhoneNumber).
				Return(repoRes, nil).
				Times(1)

			res := RegisterStaffRes{}
			err := service.RegisterStaff(context.TODO(), req, &res)
			Convey("Should return ErrPhoneNumberAlreadyRegistered", func() {
				So(
					errors.Is(err, ErrPhoneNumberAlreadyRegistered),
					ShouldBeTrue,
				)
			})
		})

		Convey("If the phone number is unique", func() {
			mockedRepo.EXPECT().
				FindStaffByPhone(gomock.Any(), req.PhoneNumber).
				Return(auth.Staff{}, auth.ErrPhoneNumberNotFound).
				Times(1)
			mockedRepo.EXPECT().
				CreateStaff(gomock.Any(), gomock.Any()).
				Do(func(_ context.Context, reqFromSvc auth.Staff) {
					So(reqFromSvc.Name, ShouldEqual, req.Name)
					So(reqFromSvc.Phone, ShouldEqual, req.PhoneNumber)
				}).
				Return(repoRes, nil).
				Times(1)

			res := RegisterStaffRes{}
			err := service.RegisterStaff(context.TODO(), req, &res)
			Convey(
				"Should return nil and write the correct result to res",
				func() {
					So(err, ShouldBeNil)
					So(res.Name, ShouldEqual, req.Name)
					So(res.PhoneNumber, ShouldEqual, req.PhoneNumber)
					So(res.UserId, ShouldEqual, repoRes.Id)
				},
			)
		})
	})
}
