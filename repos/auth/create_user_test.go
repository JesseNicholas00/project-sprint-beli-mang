//go:build integration
// +build integration

package auth_test

import (
	"context"
	"testing"

	"github.com/JesseNicholas00/BeliMang/repos/auth"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateUser(t *testing.T) {
	Convey("When inserting new user from parameter", t, func() {
		repo := NewWithTestDatabase(t)

		reqUser := auth.User{
			Id:       "testId",
			Username: "username",
			Password: "password",
			IsAdmin:  true,
		}

		resUser, err := repo.CreateUser(context.TODO(), reqUser)
		Convey("Should return the created user with the same data", func() {
			So(err, ShouldBeNil)
			So(resUser.Id, ShouldEqual, reqUser.Id)
			So(resUser.Username, ShouldEqual, reqUser.Username)
			So(resUser.IsAdmin, ShouldEqual, reqUser.IsAdmin)
			So(resUser.Password, ShouldEqual, reqUser.Password)
		})

		Convey("When inserting duplicate user", func() {
			_, err := repo.CreateUser(context.TODO(), reqUser)
			Convey("Should return duplicate error", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
