package convey_test

import (
	"testing"

	"test-goconvey"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStringSliceEqual(t *testing.T) {
	Convey("TestStringSliceEqual", t, func() {
		Convey("should return true when a != nil  && b != nil", func() {
			a := []string{"hello", "goconvey"}
			b := []string{"hello", "goconvey"}
			So(convey.StringSliceEqual(a, b), ShouldBeTrue)
		})

		Convey("should return true when a ＝= nil  && b ＝= nil", func() {
			So(convey.StringSliceEqual(nil, nil), ShouldBeTrue)
		})

		Convey("should return false when a ＝= nil  && b != nil", func() {
			a := []string(nil)
			b := []string{}
			So(convey.StringSliceEqual(a, b), ShouldBeFalse)
		})

		Convey("should return false when a != nil  && b != nil", func() {
			a := []string{"hello", "world"}
			b := []string{"hello", "goconvey"}
			So(convey.StringSliceEqual(a, b), ShouldBeFalse)
		})

		Convey("should return false when len(a) != len(b)", func() {
			a := []string{"hello"}
			b := []string{"hello", "goconvey"}
			So(convey.StringSliceEqual(a, b), ShouldBeFalse)
		})
	})
}
