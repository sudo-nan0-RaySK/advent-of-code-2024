package day4

func SolveTaskPart2() int64 {
	ans := int64(0)
	grid := ParseInputContent()
	R, C := len(grid), len(grid[0])

	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if grid[r][c] == 'A' && isValidMidPoint(grid, r, c, R, C) {
				ans += 1
			}
		}
	}
	return ans
}

func isValidMidPoint(grid [][]rune, r, c, R, C int) bool {
	if !isInBounds(r, c, R, C) {
		return false
	}
	return formsMasWord([]rune{grid[r-1][c-1], grid[r+1][c+1]}) && formsMasWord([]rune{grid[r-1][c+1], grid[r+1][c-1]})
}

func formsMasWord(runes []rune) bool {
	if (runes[0] == 'M' && runes[1] == 'S') || (runes[0] == 'S' && runes[1] == 'M') {
		return true
	}
	return false
}

func isInBounds(r, c, R, C int) bool {
	return IsSafe(r-1, c-1, R, C) &&
		IsSafe(r-1, c+1, R, C) &&
		IsSafe(r+1, c-1, R, C) &&
		IsSafe(r+1, c+1, R, C)
}
