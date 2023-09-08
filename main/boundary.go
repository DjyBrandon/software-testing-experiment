package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
)

func main() {
	bouList("./data/boundary-value-analysis.xlsx")
}

func bouList(excelPath string) []string {

	xlsx, err := excelize.OpenFile(excelPath)
	if err != nil {
		fmt.Println("open file err:", err)
	}
	sheetName := xlsx.GetSheetName(xlsx.GetActiveSheetIndex())
	rows, _ := xlsx.GetRows(sheetName)
	out := make([]string, 0)
	for _, row := range rows {
		var result = bouTriangle(row[0])
		if row[1] == result {
			fmt.Println(row[0] + "  " + row[1] + "  ✅")
		} else {
			fmt.Println(row[0] + "  " + row[1] + "  ❎")
		}
	}

	return out

}

func bouTriangle(colCel string) string {

	var a, b, c int
	var result string

	strArr := strings.Split(colCel, ",")
	a, _ = strconv.Atoi(strArr[0])
	b, _ = strconv.Atoi(strArr[1])
	c, _ = strconv.Atoi(strArr[2])

	if (a < b+c) && (b < a+c) && (c < a+b) {
		if a == b && b == c {
			result := "Equilateral"
			return result
		} else if a != b && a != c && b != c {
			result := "Scalene"
			return result
		} else {
			result := "Isosceles"
			return result
		}
	} else {
		result := "Not a Triangle"
		return result
	}

	return result
}
