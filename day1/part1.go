package day1

import (
	_ "embed"
	"math"
	"slices"
)

func SolveTaskPart1() int64 {
	var ans int64
	listA, listB := ParseInputContent()
	slices.Sort(listA)
	slices.Sort(listB)

	for idx := 0; idx < len(listA); idx++ {
		ans += int64(math.Abs(float64(listA[idx] - listB[idx])))
	}
	return ans
}
