package day2

import "fmt"

func SolveTaskPart1() int64 {
	ans := int64(0)
	reports := ParseInputContent()
	for _, report := range reports {
		diffArray := GetDiffArrayFromReport(report)
		if IsValidReport(diffArray) {
			fmt.Printf("diffArray %v from report %v is value\n", diffArray, report)
			ans += 1
		}
	}
	return ans
}
