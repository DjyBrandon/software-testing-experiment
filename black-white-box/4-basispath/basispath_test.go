package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBasTriangle(t *testing.T) {

	Convey("边界值分析测试用例判断结果", t, func() {
		So(BasTriangle("3,4,5"), ShouldEqual, "Scalene")
		So(BasTriangle("3,3,3"), ShouldEqual, "Equilateral")
		So(BasTriangle("3,4,5"), ShouldEqual, "Scalene1")
	})
}
