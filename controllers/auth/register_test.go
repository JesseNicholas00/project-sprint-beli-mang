package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KerakTelor86/GoBoiler/services/auth"
	"github.com/KerakTelor86/GoBoiler/utils/helper"
	"github.com/KerakTelor86/GoBoiler/utils/unittesting"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRegisterValid(t *testing.T) {
	Convey("When given a valid register request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		userId := "dummyId"
		name := "namadepan namabelakang"
		phoneNumber := "+1-2468123123123"
		password := "password"
		accessToken := "token"

		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodPost,
			"/v1/staff/register",
			rec,
			unittesting.WithJsonPayload(map[string]interface{}{
				"phoneNumber": phoneNumber,
				"name":        name,
				"password":    password,
			}),
		)

		Convey("Should forward the request to the service layer", func() {
			expectedReq := auth.RegisterStaffReq{
				PhoneNumber: phoneNumber,
				Name:        name,
				Password:    password,
			}
			expectedRes := auth.RegisterStaffRes{
				UserId:      userId,
				PhoneNumber: phoneNumber,
				Name:        name,
				AccessToken: accessToken,
			}

			service.
				EXPECT().
				RegisterStaff(gomock.Any(), expectedReq, gomock.Any()).
				Do(func(_ context.Context, _ auth.RegisterStaffReq, res *auth.RegisterStaffRes) {
					*res = expectedRes
				}).
				Return(nil).
				Times(1)

			unittesting.CallController(ctx, controller.registerStaff)

			Convey(
				"Should return the expected response with HTTP 201",
				func() {
					So(rec.Code, ShouldEqual, http.StatusCreated)

					expectedBody := helper.MustMarshalJson(
						map[string]interface{}{
							"message": "User registered successfully",
							"data":    expectedRes,
						},
					)

					So(
						rec.Body.String(),
						ShouldEqualJSON,
						string(expectedBody),
					)
				},
			)
		})
	})
}

func TestRegisterInvalid(t *testing.T) {
	Convey("When given an invalid register request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		name := "firstname lastname"
		phoneNumber := "+1-2468123123123"
		password := "password"

		Convey("On invalid request", func() {
			phoneNumber := "+1-2468123123123"
			password := "password"

			rec := httptest.NewRecorder()
			ctx := unittesting.CreateEchoContextFromRequest(
				http.MethodPost,
				"/v1/staff/register",
				rec,
				unittesting.WithJsonPayload(map[string]interface{}{
					// no name
					"phoneNumber": phoneNumber,
					"password":    password,
				}),
			)

			Convey("Should return HTTP code 400", func() {
				unittesting.CallController(ctx, controller.registerStaff)
				So(rec.Code, ShouldEqual, http.StatusBadRequest)
			})
		})

		Convey("On duplicate phone number", func() {
			rec := httptest.NewRecorder()
			ctx := unittesting.CreateEchoContextFromRequest(
				http.MethodPost,
				"/v1/staff/register",
				rec,
				unittesting.WithJsonPayload(map[string]interface{}{
					"name":        name,
					"phoneNumber": phoneNumber,
					"password":    password,
				}),
			)

			Convey("Should return HTTP code 409", func() {
				expectedReq := auth.RegisterStaffReq{
					PhoneNumber: phoneNumber,
					Name:        name,
					Password:    password,
				}

				service.EXPECT().
					RegisterStaff(gomock.Any(), expectedReq, gomock.Any()).
					Return(auth.ErrPhoneNumberAlreadyRegistered).
					Times(1)

				unittesting.CallController(ctx, controller.registerStaff)
				So(rec.Code, ShouldEqual, http.StatusConflict)
			})
		})
	})
}
