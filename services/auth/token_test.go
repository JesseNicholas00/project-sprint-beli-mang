package auth

import (
	"testing"

	"github.com/KerakTelor86/GoBoiler/repos/auth"
	"github.com/golang-jwt/jwt/v4"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateToken(t *testing.T) {
	Convey("When generating token from staff", t, func() {
		mockCtrl, service, _ := NewWithMockedRepo(t)
		defer mockCtrl.Finish()

		staff := auth.Staff{
			Id:    "bread",
			Name:  "firstname lastname",
			Phone: "+621234567890",
		}

		token, err := service.generateToken(staff)
		Convey(
			"Should return a token containing staff data without errors",
			func() {
				So(err, ShouldBeNil)

				claims := jwtClaims{}
				_, err := jwt.ParseWithClaims(
					token,
					&claims,
					func(t *jwt.Token) (interface{}, error) {
						return service.jwtSecret, nil
					},
				)

				So(err, ShouldBeNil)
				So(claims.Data.Name, ShouldEqual, staff.Name)
				So(claims.Data.PhoneNumber, ShouldEqual, staff.Phone)
				So(claims.Data.UserId, ShouldEqual, staff.Id)
			},
		)
	})
}
