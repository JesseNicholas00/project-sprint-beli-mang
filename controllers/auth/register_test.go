package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JesseNicholas00/BeliMang/services/auth"
	"github.com/JesseNicholas00/BeliMang/utils/helper"
	"github.com/JesseNicholas00/BeliMang/utils/unittesting"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRegisterValid(t *testing.T) {
	Convey("When given a valid register request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		username := "username"
		email := "jn@gmail.com"
		password := "password"
		accessToken := "token"
		role := "admin"

		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodPost,
			"/admin/register",
			rec,
			unittesting.WithJsonPayload(map[string]interface{}{
				"username": username,
				"email":    email,
				"password": password,
			}),
			unittesting.WithPathParams(map[string]string{
				"role": "admin",
			}),
		)

		Convey("Should forward the request to the service layer", func() {
			expectedReq := auth.RegisterUserReq{
				Username: username,
				Email:    email,
				Password: password,
				Role:     role,
			}
			expectedRes := auth.RegisterUserRes{
				AccessToken: accessToken,
			}

			service.
				EXPECT().
				RegisterUser(gomock.Any(), expectedReq, gomock.Any()).
				Do(func(_ context.Context, _ auth.RegisterUserReq, res *auth.RegisterUserRes) {
					*res = expectedRes
				}).
				Return(nil).
				Times(1)

			unittesting.CallController(ctx, controller.registerUser)

			Convey(
				"Should return the expected response with HTTP 201",
				func() {
					So(rec.Code, ShouldEqual, http.StatusCreated)

					expectedBody := helper.MustMarshalJson(expectedRes)

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

		username := "username"
		email := "jn@gmail.com"
		password := "password"
		role := "admin"

		Convey("On invalid request", func() {
			phoneNumber := "+1-2468123123123"
			password := "password"

			rec := httptest.NewRecorder()
			ctx := unittesting.CreateEchoContextFromRequest(
				http.MethodPost,
				"/admin/register",
				rec,
				unittesting.WithJsonPayload(map[string]interface{}{
					// no name
					"phoneNumber": phoneNumber,
					"password":    password,
				}),
			)

			Convey("Should return HTTP code 400", func() {
				unittesting.CallController(ctx, controller.registerUser)
				So(rec.Code, ShouldEqual, http.StatusBadRequest)
			})
		})

		Convey("On duplicate username", func() {
			rec := httptest.NewRecorder()
			ctx := unittesting.CreateEchoContextFromRequest(
				http.MethodPost,
				"/admin/register",
				rec,
				unittesting.WithJsonPayload(map[string]interface{}{
					"username": username,
					"email":    email,
					"password": password,
				}),
				unittesting.WithPathParams(map[string]string{
					"role": "admin",
				}),
			)

			Convey("Should return HTTP code 409", func() {
				expectedReq := auth.RegisterUserReq{
					Username: username,
					Email:    email,
					Password: password,
					Role:     role,
				}

				service.EXPECT().
					RegisterUser(gomock.Any(), expectedReq, gomock.Any()).
					Return(auth.ErrUsernameAlreadyRegistered).
					Times(1)

				unittesting.CallController(ctx, controller.registerUser)
				So(rec.Code, ShouldEqual, http.StatusConflict)
			})
		})
	})
}
