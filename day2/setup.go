package day2

import (
	"advent-of-code-2024/utils"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed static/input.txt
var inputContent string

func ParseInputContent() [][]int64 {
	var ret [][]int64
	reports := strings.Split(inputContent, "\n")
	ret = make([][]int64, len(reports))
	for idx, report := range reports {
		ret[idx] = parseReport(&report)
	}
	return ret
}

func parseReport(reportArray *string) []int64 {
	var ret []int64
	report := strings.Split(*reportArray, " ")
	ret = make([]int64, len(report))
	for idx, level := range report {
		ret[idx] = utils.Must(strconv.ParseInt(level, 10, 64))
	}
	return ret
}

func IsValidReport(diffArray []int64) bool {
	isDecreasing := true
	for idx, value := range diffArray {
		if idx == 0 && isDecreasing {
			isDecreasing = value < 0
		}
		if (value >= 0 && isDecreasing) || (value <= 0 && !isDecreasing) {
			return false
		}
		if (isDecreasing && (value < -3)) || (!isDecreasing && (value > 3)) {
			return false
		}
	}
	return true
}

func GetDiffArrayFromReport(report []int64) []int64 {
	diffArray := make([]int64, len(report)-1)
	for idx := 1; idx < len(report); idx = idx + 1 {
		diffArray[idx-1] = report[idx] - report[idx-1]
	}
	return diffArray
}
