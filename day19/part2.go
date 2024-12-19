package day19

func SolveTaskPart2() int {
	patternInfo := ParseInputContent()
	ans := 0
	availablePatterns := patternInfo.AvailableDesigns
	patternsToCheck := patternInfo.DesignsToCheck
	for _, pattern := range patternsToCheck {
		ans += NumberOfWays(pattern, availablePatterns)
	}
	return ans
}

func NumberOfWays(patternToMake string, availablePatterns []string) int {
	var memo = make(map[int]int)
	var dp func(startPos int) int
	dp = func(startPos int) int {
		if startPos >= len(patternToMake) {
			return 1
		}

		if val, ok := memo[startPos]; ok {
			return val
		}

		ans := 0
		for _, pattern := range availablePatterns {
			if len(patternToMake) < startPos+len(pattern) {
				continue
			}
			if pattern == patternToMake[startPos:startPos+len(pattern)] {
				ans += dp(startPos + len(pattern))
			}
		}
		memo[startPos] = ans
		return memo[startPos]
	}
	return dp(0)
}
