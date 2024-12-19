package day19

import (
	"fmt"
	"github.com/gammazero/deque"
)

func SolveTaskPart1() int {
	patternInfo := ParseInputContent()
	availablePatterns := patternInfo.AvailableDesigns
	fmt.Printf("len(availablePatterns): %d\n", len(availablePatterns))
	patternsToTest := patternInfo.DesignsToCheck
	ans := 0
	for _, pattern := range patternsToTest {
		if CanMakePattern(availablePatterns, pattern) {
			ans += 1
		}
	}
	return ans
}

func CanMakePattern(availablePatterns []string, patternToTest string) bool {
	q := new(deque.Deque[int])
	q.PushBack(0)
	matched := make(map[int]bool)
	for q.Len() > 0 {
		startPos := q.PopFront()
		if startPos >= len(patternToTest) {
			return true
		}
		for _, pattern := range availablePatterns {
			if startPos+len(pattern) > len(patternToTest) {
				continue
			}
			if !matched[startPos+len(pattern)] && patternToTest[startPos:startPos+len(pattern)] == pattern {
				matched[startPos+len(pattern)] = true
				q.PushBack(startPos + len(pattern))
			}
		}
	}
	return false
}
