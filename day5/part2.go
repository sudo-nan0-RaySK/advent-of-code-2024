package day5

import (
	"fmt"
	"slices"
)

func SolveTaskPart2() int64 {
	ans := int64(0)
	orderingAndUpdates := ParseInputContent()
	ordering := orderingAndUpdates.GetOrdering()
	updates := orderingAndUpdates.GetUpdates()
	for _, update := range updates {
		sortedUpdates := make([]int64, len(update))
		copy(sortedUpdates, update)
		slices.SortFunc(sortedUpdates, func(a, b int64) int {
			if set, ok := ordering[a]; ok {
				if _, contains := set[b]; contains {
					return -1
				}
			}
			if set, ok := ordering[b]; ok {
				if _, contains := set[a]; contains {
					return 1
				}
			}
			// Node have no mutual ordering relation
			return 0
		})
		if !slices.Equal(sortedUpdates, update) {
			fmt.Printf("%v is valid\n", update)
			ans += sortedUpdates[len(update)/2]
		}
	}
	return ans
}
