package day6

import "fmt"

func SolveTaskPart1() int64 {
	parsedGrid := ParseInputContent()
	grid := make([][]rune, len(parsedGrid))
	copy(grid, parsedGrid)
	r, c := FindGuardPosition(grid)
	markGuardPath(grid, r, c)
	printGrid(grid)
	return countVisitedCells(grid)
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func countVisitedCells(grid [][]rune) int64 {
	ans := int64(0)
	for _, row := range grid {
		for _, cell := range row {
			if cell == 'X' {
				ans += 1
			}
		}
	}
	return ans
}

func markGuardPath(grid [][]rune, r int, c int) {
	R, C := len(grid), len(grid[0])
	isSafe := func(row, col int) bool {
		return row >= 0 && row < R && col >= 0 && col < C
	}
	direction := GetDirection(grid, r, c)
	gr, gc := r, c
	for isSafe(gr, gc) {
		metObstruction := grid[gr][gc] == '#'
		if direction == UP {
			if metObstruction {
				gr += 1
				direction = RIGHT
			} else {
				grid[gr][gc] = 'X'
				gr -= 1
			}
		} else if direction == DOWN {
			if metObstruction {
				gr -= 1
				direction = LEFT
			} else {
				grid[gr][gc] = 'X'
				gr += 1
			}
		} else if direction == LEFT {
			if metObstruction {
				gc += 1
				direction = UP
			} else {
				grid[gr][gc] = 'X'
				gc -= 1
			}
		} else if direction == RIGHT {
			if metObstruction {
				gc -= 1
				direction = DOWN
			} else {
				grid[gr][gc] = 'X'
				gc += 1
			}
		}
	}
}
