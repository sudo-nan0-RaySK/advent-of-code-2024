package day1

func summarizeList(list []int64) map[int64]int64 {
	summary := make(map[int64]int64)
	for _, number := range list {
		if _, ok := summary[number]; ok {
			summary[number] += 1
		} else {
			summary[number] = 1
		}
	}
	return summary
}

func SolveTaskPart2() int64 {
	ans := int64(0)
	listA, listB := ParseInputContent()
	listBSummary := summarizeList(listB)

	for _, num := range listA {
		ans += num * listBSummary[num]
	}

	return ans
}
