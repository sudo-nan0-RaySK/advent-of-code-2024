package day4

import "slices"

func SolveTaskPart1() int64 {
	grid := ParseInputContent()
	return findMatches(grid, "XMAS") + findMatches(grid, "SAMX")
}

func findMatches(grid [][]rune, pattern string) int64 {
	return searchHorizontal(grid, pattern) +
		searchVertical(grid, pattern) +
		searchForwardDiagonal(grid, pattern) +
		searchBackwardDiagonal(grid, pattern)
}

func searchHorizontal(grid [][]rune, pattern string) int64 {
	ans := int64(0)
	h, w := len(grid), len(grid[0])
	patterRune := []rune(pattern)
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if IsSafe(r, c+3, h, w) && slices.Equal(grid[r][c:c+4], patterRune) {
				ans += 1
			}
		}
	}
	return ans
}

func searchVertical(grid [][]rune, pattern string) int64 {
	ans := int64(0)
	h, w := len(grid), len(grid[0])
	patterRune := []rune(pattern)
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if IsSafe(r+3, c, h, w) && slices.Equal(getVerticalSlice(grid, r, c), patterRune) {
				ans += 1
			}
		}
	}
	return ans
}

func searchForwardDiagonal(grid [][]rune, pattern string) int64 {
	ans := int64(0)
	h, w := len(grid), len(grid[0])
	patterRune := []rune(pattern)
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if IsSafe(r+3, c+3, h, w) && slices.Equal(getForwardDiagonalSlice(grid, r, c), patterRune) {
				ans += 1
			}
		}
	}
	return ans
}

func searchBackwardDiagonal(grid [][]rune, pattern string) int64 {
	ans := int64(0)
	h, w := len(grid), len(grid[0])
	patterRune := []rune(pattern)
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if IsSafe(r+3, c-3, h, w) && slices.Equal(getBackwardDiagonalSlice(grid, r, c), patterRune) {
				ans += 1
			}
		}
	}
	return ans
}

func getForwardDiagonalSlice(grid [][]rune, r int, c int) []rune {
	fDiagonalSlice := make([]rune, 4)
	for idx := 0; idx < 4; idx++ {
		fDiagonalSlice[idx] = grid[r+idx][c+idx]
	}
	return fDiagonalSlice
}

func getBackwardDiagonalSlice(grid [][]rune, r int, c int) []rune {
	bDiagonalSlice := make([]rune, 4)
	for idx := 0; idx < 4; idx++ {
		bDiagonalSlice[idx] = grid[r+idx][c-idx]
	}
	return bDiagonalSlice
}

func getVerticalSlice(grid [][]rune, r, c int) []rune {
	vSlice := make([]rune, 4)
	for idx := 0; idx < 4; idx++ {
		vSlice[idx] = grid[r+idx][c]
	}
	return vSlice
}
