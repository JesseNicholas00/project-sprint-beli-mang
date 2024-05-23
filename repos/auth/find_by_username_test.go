//go:build integration
// +build integration

package auth_test

import (
	"context"
	"errors"
	"testing"

	"github.com/JesseNicholas00/BeliMang/repos/auth"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFindUsername(t *testing.T) {
	Convey(
		"When database contains user members with different usernames",
		t,
		func() {
			repo := NewWithTestDatabase(t)

			dummyIds := []string{
				"id1",
				"id2",
				"id3",
			}
			dummyEmails := []string{
				"email1",
				"email2",
				"email3",
			}
			dummyUsernames := []string{
				"usernameA",
				"usernameB",
				"usernameC",
			}
			for i := range dummyIds {
				curReqUser := auth.User{
					Id:       dummyIds[i],
					Username: dummyUsernames[i],
					Email:    dummyEmails[i],
					IsAdmin:  true,
					Password: "hashedPasswordVeryScure",
				}
				_, err := repo.CreateUser(context.TODO(), curReqUser)
				So(err, ShouldBeNil)
			}

			Convey(
				"Should return the user with the requested username if one exists",
				func() {
					for _, expectedUsername := range dummyUsernames {
						resUser, err := repo.FindUserByUsername(
							context.TODO(),
							expectedUsername,
						)
						So(err, ShouldBeNil)
						So(resUser.Username, ShouldEqual, expectedUsername)
					}
				},
			)

			Convey(
				"Should return ErrPhoneNumberNotFound when username doesn't exist",
				func() {
					_, err := repo.FindUserByUsername(
						context.TODO(),
						"+123456789015",
					)
					So(
						errors.Is(err, auth.ErrUsernameNotFound),
						ShouldBeTrue,
					)
				},
			)
		},
	)
}
