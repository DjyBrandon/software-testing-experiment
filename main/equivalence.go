package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
)

func main() {
	equList("./data/equivalence-partitioning.xlsx")
}

func equList(excelPath string) []string {

	xlsx, err := excelize.OpenFile(excelPath)
	if err != nil {
		fmt.Println("open file err:", err)
	}
	sheetName := xlsx.GetSheetName(xlsx.GetActiveSheetIndex())
	rows, _ := xlsx.GetRows(sheetName)
	out := make([]string, 0)
	for _, row := range rows {
		var result = equTriangle(row[0])
		if row[1] == result {
			fmt.Println(row[0] + "  " + row[1] + "  ✅")
		} else {
			fmt.Println(row[0] + "  " + row[1] + "  ❎")
		}
	}

	return out

}

func equTriangle(colCel string) string {

	var a, b, c int
	var result string

	strArr := strings.Split(colCel, ",")
	a, _ = strconv.Atoi(strArr[0])
	b, _ = strconv.Atoi(strArr[1])
	c, _ = strconv.Atoi(strArr[2])

	if a <= 0 || a >= 200 || b <= 0 || b >= 200 || c <= 0 || c >= 200 {
		result := "Not a Triangle"
		return result
	} else {
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
	}

	return result
}
