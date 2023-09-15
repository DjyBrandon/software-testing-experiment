package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBouTriangle(t *testing.T) {

	Convey("Boundary Value Analysis Test Case Result", t, func() {
		So(BouTriangle("100,100,0"), ShouldEqual, "Not a Triangle")
		So(BouTriangle("100,100,1"), ShouldEqual, "Isosceles")
		So(BouTriangle("100,100,100"), ShouldEqual, "Equilateral")
		So(BouTriangle("100,100,199"), ShouldEqual, "Isosceles")
		So(BouTriangle("100,100,200"), ShouldEqual, "Not a Triangle")
		So(BouTriangle("100,0,100"), ShouldEqual, "Not a Triangle")
		So(BouTriangle("100,1,100"), ShouldEqual, "Isosceles")
		So(BouTriangle("100,199,100"), ShouldEqual, "Isosceles")
		So(BouTriangle("100,200,100"), ShouldEqual, "Not a Triangle")
		So(BouTriangle("0,100,100"), ShouldEqual, "Not a Triangle")
		So(BouTriangle("1,100,100"), ShouldEqual, "Isosceles")
		So(BouTriangle("199,100,100"), ShouldEqual, "Isosceles")
		So(BouTriangle("200,100,100"), ShouldEqual, "Not a Triangle")
		So(BouTriangle("100,0,0"), ShouldEqual, "Not a Triangle")
		So(BouTriangle("100,0,1"), ShouldEqual, "Not a Triangle")
		So(BouTriangle("100,0,199"), ShouldEqual, "Not a Triangle")
		So(BouTriangle("100,0,200"), ShouldEqual, "Not a Triangle")
		So(BouTriangle("100,1,0"), ShouldEqual, "Not a Triangle")
		So(BouTriangle("100,1,1"), ShouldEqual, "Not a Triangle")
		So(BouTriangle("100,1,199"), ShouldEqual, "Not a Triangle")
		So(BouTriangle("100,1,200"), ShouldEqual, "Not a Triangle")
		So(BouTriangle("100,199,0"), ShouldEqual, "Not a Triangle")
		So(BouTriangle("100,199,1"), ShouldEqual, "Not a Triangle")
		So(BouTriangle("100,199,199"), ShouldEqual, "Isosceles")
		So(BouTriangle("100,199,200"), ShouldEqual, "Not a Triangle")
		So(BouTriangle("100,200,0"), ShouldEqual, "Not a Triangle")
		So(BouTriangle("100,200,1"), ShouldEqual, "Not a Triangle")
		So(BouTriangle("100,200,199"), ShouldEqual, "Not a Triangle")
		So(BouTriangle("100,200,200"), ShouldEqual, "Not a Triangle")

	})
}
