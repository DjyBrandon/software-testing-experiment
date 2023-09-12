package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJudge(t *testing.T) {

	var triangle = "3,4,5"
	Convey("基础三角形测试用例判断结果", t, func() {
		So(Judge(triangle), ShouldEqual, "Scalene")
	})
}
