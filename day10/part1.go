package day10

func SolveTaskPart1() int64 {
	ans := int64(0)
	gridInfo := ParseInputContent()
	gridSize := gridInfo.Size()
	R, C := gridSize.Row(), gridSize.Col()
	grid := gridInfo.Grid()
	zeroLocations := gridInfo.ZeroPositions()
	for _, location := range zeroLocations {
		ans += findScore(grid, location, R, C)
	}
	return ans
}

func findScore(grid [][]int, location Point, R int, C int) int64 {
	directions := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	isSafe := func(location Point) bool {
		return location.Row() >= 0 && location.Row() < R && location.Col() >= 0 && location.Col() < C
	}
	unique9s := make(map[Point]bool)
	var dp func(pt Point, last int) int64
	dp = func(pt Point, last int) int64 {
		if last == 9 && !unique9s[pt] {
			unique9s[pt] = true
			return 1
		}
		ans := int64(0)
		r, c := pt.Row(), pt.Col()
		for _, direction := range directions {
			dr, dc := direction[0], direction[1]
			nextPt := Point{r + dr, c + dc}
			if isSafe(nextPt) && grid[nextPt.Row()][nextPt.Col()] == last+1 {
				ans += dp(nextPt, last+1)
			}
		}
		return ans
	}
	return dp(location, 0)
}
