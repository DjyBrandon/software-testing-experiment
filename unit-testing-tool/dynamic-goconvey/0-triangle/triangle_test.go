package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJudge(t *testing.T) {

	var triangle = "3,4,5"
	Convey("Common Program Triangle Test Case Judge Result\n", t, func() {
		So(Judge(triangle), ShouldEqual, "Scalene")
	})
}
