package helper_test

import (
	"testing"

	"github.com/KerakTelor86/GoBoiler/utils/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHasLen(t *testing.T) {
	Convey("When given a number of the correct length", t, func() {
		number := int64(123456789)
		Convey("Should return true", func() {
			So(helper.HasLen(number, 9), ShouldBeTrue)
		})
	})
	Convey("When given a number of an incorrect length", t, func() {
		number := int64(123456789)
		Convey("Should return false", func() {
			So(helper.HasLen(number, 8), ShouldBeFalse)
			So(helper.HasLen(number, 10), ShouldBeFalse)
		})
	})
}

func TestGetSubDigit(t *testing.T) {
	Convey("When getting subdigits of a number", t, func() {
		len := 9
		number := int64(123456789)
		Convey("Should return the expected values", func() {
			So(helper.GetSubDigit(number, len, 1, 3), ShouldEqual, 123)
			So(helper.GetSubDigit(number, len, 4, 8), ShouldEqual, 45678)
			So(helper.GetSubDigit(number, len, 9, 9), ShouldEqual, 9)
		})
	})
}

func TestGetLen(t *testing.T) {
	Convey("When getting the length of a number", t, func() {
		Convey("Should return the expected values", func() {
			So(helper.GetLen(123), ShouldEqual, 3)
			So(helper.GetLen(12345678), ShouldEqual, 8)
			So(helper.GetLen(1), ShouldEqual, 1)
			So(helper.GetLen(0), ShouldEqual, 0)
		})
	})
}
