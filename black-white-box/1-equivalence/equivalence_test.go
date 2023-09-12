package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEquTriangle(t *testing.T) {

	Convey("等价类划分测试用例判断结果", t, func() {
		So(EquTriangle("3,4,5"), ShouldEqual, "Scalene")
		So(EquTriangle("3,3,3"), ShouldEqual, "Equilateral")
		So(EquTriangle("3,4,5"), ShouldEqual, "Scalene1")
	})
}
