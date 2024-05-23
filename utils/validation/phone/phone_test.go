package phone

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestValidatePhoneNumberImpl(t *testing.T) {
	Convey("When phone number length is invalid", t, func() {
		Convey("Should return false", func() {
			So(validatePhoneNumberImpl("127836"), ShouldBeFalse)
			So(
				validatePhoneNumberImpl("1289312983671293812867123"),
				ShouldBeFalse,
			)
		})
	})

	Convey("When phone number does not start with +62", t, func() {
		Convey("Should return false", func() {
			So(validatePhoneNumberImpl("-21678391236"), ShouldBeFalse)
		})
	})

	Convey("When phone number is valid", t, func() {
		Convey("Should return true", func() {
			So(validatePhoneNumberImpl("+6281234567890"), ShouldBeTrue)
		})
	})
}
