//go:build integration
// +build integration

package auth_test

import (
	"context"
	"errors"
	"testing"

	"github.com/KerakTelor86/GoBoiler/repos/auth"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFindStaffByPhone(t *testing.T) {
	Convey(
		"When database contains staff members with different phone numbers",
		t,
		func() {
			repo := NewWithTestDatabase(t)

			dummyIds := []string{
				"id1",
				"id2",
				"id3",
			}
			dummyPhones := []string{
				"+123456789012",
				"+123456789013",
				"+123456789014",
			}
			for i := range dummyIds {
				curReqStaff := auth.Staff{
					Id:       dummyIds[i],
					Name:     "firstname lastname",
					Phone:    dummyPhones[i],
					Password: "hashedPasswordVeryScure",
				}
				_, err := repo.CreateStaff(context.TODO(), curReqStaff)
				So(err, ShouldBeNil)
			}

			Convey(
				"Should return the staff with the requested phone number if one exists",
				func() {
					for _, expectedPhone := range dummyPhones {
						resStaff, err := repo.FindStaffByPhone(
							context.TODO(),
							expectedPhone,
						)
						So(err, ShouldBeNil)
						So(resStaff.Phone, ShouldEqual, expectedPhone)
					}
				},
			)

			Convey(
				"Should return ErrPhoneNumberNotFound when phone number doesn't exist",
				func() {
					_, err := repo.FindStaffByPhone(
						context.TODO(),
						"+123456789015",
					)
					So(
						errors.Is(err, auth.ErrPhoneNumberNotFound),
						ShouldBeTrue,
					)
				},
			)
		},
	)
}
