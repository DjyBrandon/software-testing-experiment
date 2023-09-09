package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var triangle string
	var a, b, c int

	fmt.Println("Enter 3 integers which are sides of a triangle (e.g.: 3,4,5)")
	fmt.Scanln(&triangle)

	strArr := strings.Split(triangle, ",")
	a, _ = strconv.Atoi(strArr[0])
	fmt.Println("Side A is", a)
	b, _ = strconv.Atoi(strArr[1])
	fmt.Println("Side B is:", b)
	c, _ = strconv.Atoi(strArr[2])
	fmt.Println("Side C is:", c)

	if a <= 0 || a >= 200 || b <= 0 || b >= 200 || c <= 0 || c >= 200 {
		fmt.Println("Not a Triangle")
	} else {
		if (a < b+c) && (b < a+c) && (c < a+b) {
			if a == b && b == c {
				fmt.Println("Equilateral")
			} else if a != b && a != c && b != c {
				fmt.Println("Scalene")
			} else {
				fmt.Println("Isosceles")
			}
		} else {
			fmt.Println("Not a Triangle")
		}
	}

}
