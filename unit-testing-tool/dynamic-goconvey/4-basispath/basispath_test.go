package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBasTriangle(t *testing.T) {

	Convey("Basis Path Test Case Result", t, func() {
		So(BasTriangle("2,2,2"), ShouldEqual, "Equilateral")
		So(BasTriangle("3,4,5"), ShouldEqual, "Scalene")
		So(BasTriangle("2,2,3"), ShouldEqual, "Isosceles")
		So(BasTriangle("1,2,3"), ShouldEqual, "Not a Triangle")
		So(BasTriangle("0,4,5"), ShouldEqual, "Not a Triangle")
		So(BasTriangle("200,4,5"), ShouldEqual, "Not a Triangle")
	})
}
