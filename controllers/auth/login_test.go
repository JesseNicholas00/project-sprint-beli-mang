package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JesseNicholas00/BeliMang/services/auth"
	"github.com/JesseNicholas00/BeliMang/types/role"
	"github.com/JesseNicholas00/BeliMang/utils/helper"
	"github.com/JesseNicholas00/BeliMang/utils/unittesting"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLoginValid(t *testing.T) {
	Convey("When given a valid login request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		username := "username"
		password := "password"
		accessToken := "token"

		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodPost,
			"/user/login",
			rec,
			unittesting.WithJsonPayload(map[string]interface{}{
				"username": username,
				"password": password,
			}),
		)

		Convey("Should forward the request to the service layer", func() {
			expectedReq := auth.LoginUserReq{
				Username: username,
				Password: password,
				Role:     role.User,
			}
			expectedRes := auth.LoginUserRes{
				AccessToken: accessToken,
			}

			service.
				EXPECT().
				LoginUser(gomock.Any(), expectedReq, gomock.Any()).
				Do(func(_ context.Context, _ auth.LoginUserReq, res *auth.LoginUserRes) {
					*res = expectedRes
				}).
				Return(nil).
				Times(1)

			unittesting.CallControllerWithRole(ctx, controller.loginUser, role.User)

			Convey(
				"Should return the expected response with HTTP 200",
				func() {
					So(rec.Code, ShouldEqual, http.StatusOK)

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

func TestLoginInvalid(t *testing.T) {
	Convey("When given an invalid login request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		Convey("On bad request", func() {
			// username length too short
			username := "user"
			password := "password"

			rec := httptest.NewRecorder()
			ctx := unittesting.CreateEchoContextFromRequest(
				http.MethodPost,
				"/user/login",
				rec,
				unittesting.WithJsonPayload(map[string]interface{}{
					"username": username,
					"password": password,
				}),
			)

			Convey("Should return HTTP code 400", func() {
				unittesting.CallControllerWithRole(ctx, controller.loginUser, role.User)
				So(rec.Code, ShouldEqual, http.StatusBadRequest)
			})
		})

		username := "username"
		password := "password"
		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodPost,
			"/user/login",
			rec,
			unittesting.WithJsonPayload(map[string]interface{}{
				"username": username,
				"password": password,
			}),
		)

		expectedReq := auth.LoginUserReq{
			Username: username,
			Password: password,
			Role:     role.User,
		}

		Convey("On user not found", func() {
			service.
				EXPECT().
				LoginUser(gomock.Any(), expectedReq, gomock.Any()).
				Return(auth.ErrUserNotFound).
				Times(1)

			Convey(
				"Should return HTTP code 400",
				func() {
					unittesting.CallControllerWithRole(ctx, controller.loginUser, role.User)
					So(rec.Code, ShouldEqual, http.StatusBadRequest)
				},
			)
		})

		Convey("On invalid credentials", func() {
			service.
				EXPECT().
				LoginUser(gomock.Any(), expectedReq, gomock.Any()).
				Return(auth.ErrInvalidCredentials).
				Times(1)

			Convey(
				"Should return HTTP code 400",
				func() {
					unittesting.CallControllerWithRole(ctx, controller.loginUser, role.User)
					So(rec.Code, ShouldEqual, http.StatusBadRequest)
				},
			)
		})
	})
}
