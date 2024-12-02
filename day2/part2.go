package day2

func SolveTaskPart2() int64 {
	ans := int64(0)
	reports := ParseInputContent()
	for _, originalReport := range reports {
		report := make([]int64, len(originalReport))
		copy(report, originalReport)
		diffArray := GetDiffArrayFromReport(report)
		if IsValidReport(diffArray) {
			ans += 1
			continue
		}
		for idx := 0; idx < len(originalReport); idx++ {
			// Deleting element at index idx
			report = make([]int64, len(originalReport))
			copy(report, originalReport)
			report = append(report[:idx], report[idx+1:]...)
			diffArray := GetDiffArrayFromReport(report)
			if IsValidReport(diffArray) {
				ans += 1
				break
			}
		}
	}
	return ans
}
