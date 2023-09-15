package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLogTriangle(t *testing.T) {

	Convey("Logic Coverage Test Case Result", t, func() {
		So(LogTriangle("2,2,2"), ShouldEqual, "Equilateral")
		So(LogTriangle("3,4,5"), ShouldEqual, "Scalene")
		So(LogTriangle("2,2,3"), ShouldEqual, "Isosceles")
		So(LogTriangle("1,2,3"), ShouldEqual, "Not a Triangle")
		So(LogTriangle("0,4,5"), ShouldEqual, "Not a Triangle")
		So(LogTriangle("200,4,5"), ShouldEqual, "Not a Triangle")
	})
}
