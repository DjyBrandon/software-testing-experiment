package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEquTriangle(t *testing.T) {

	Convey("Equivalence Partitioning Test Case Result", t, func() {

		Convey("3,4,5", func() {
			So(EquTriangle("3,4,5"), ShouldEqual, "Scalene")
		})
		Convey("0,4,5", func() {
			So(EquTriangle("0,4,5"), ShouldEqual, "Not a Triangle")
		})
		Convey("200,4,5", func() {
			So(EquTriangle("200,4,5"), ShouldEqual, "Not a Triangle")
		})
		Convey("3,0,5", func() {
			So(EquTriangle("3,0,5"), ShouldEqual, "Not a Triangle")
		})
		Convey("3,200,5", func() {
			So(EquTriangle("3,200,5"), ShouldEqual, "Not a Triangle")
		})
		Convey("3,4,0", func() {
			So(EquTriangle("3,4,0"), ShouldEqual, "Not a Triangle")
		})
		Convey("3,4,200", func() {
			So(EquTriangle("3,4,200"), ShouldEqual, "Not a Triangle")
		})
		Convey("1,2,3", func() {
			So(EquTriangle("1,2,3"), ShouldEqual, "Not a Triangle")
		})
		Convey("3,1,2", func() {
			So(EquTriangle("3,1,2"), ShouldEqual, "Not a Triangle")
		})
		Convey("1,3,2", func() {
			So(EquTriangle("1,3,2"), ShouldEqual, "Not a Triangle")
		})
		Convey("3,3,3", func() {
			So(EquTriangle("3,3,3"), ShouldEqual, "Equilateral")
		})
		Convey("2,2,3", func() {
			So(EquTriangle("2,2,3"), ShouldEqual, "Isosceles")
		})
		Convey("3,2,2", func() {
			So(EquTriangle("3,2,2"), ShouldEqual, "Isosceles")
		})
		Convey("2,3,2", func() {
			So(EquTriangle("2,3,2"), ShouldEqual, "Isosceles")
		})
	})
}
