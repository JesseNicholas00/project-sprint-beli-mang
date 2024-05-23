//go:build integration
// +build integration

package auth_test

import (
	"context"
	"testing"

	"github.com/KerakTelor86/GoBoiler/repos/auth"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateStaff(t *testing.T) {
	Convey("When inserting new staff from parameter", t, func() {
		repo := NewWithTestDatabase(t)

		reqStaff := auth.Staff{
			Id:       "testId",
			Name:     "firstname lastname",
			Phone:    "+123456789012",
			Password: "hashedPasswordVerySecure",
		}

		resStaff, err := repo.CreateStaff(context.TODO(), reqStaff)
		Convey("Should return the created staff with the same data", func() {
			So(err, ShouldBeNil)
			So(resStaff.Id, ShouldEqual, reqStaff.Id)
			So(resStaff.Name, ShouldEqual, reqStaff.Name)
			So(resStaff.Phone, ShouldEqual, reqStaff.Phone)
			So(resStaff.Password, ShouldEqual, reqStaff.Password)
		})

		Convey("When inserting duplicate staff", func() {
			_, err := repo.CreateStaff(context.TODO(), reqStaff)
			Convey("Should return duplicate error", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
