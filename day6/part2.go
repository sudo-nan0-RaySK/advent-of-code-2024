package day6

func SolveTaskPart2() int64 {
	ans := int64(0)
	inputGrid := ParseInputContent()
	R, C := len(inputGrid), len(inputGrid[0])
	gr, gc := FindGuardPosition(inputGrid)
	for row := 0; row < R; row++ {
		for col := 0; col < C; col++ {
			grid := copy2d(inputGrid)
			if row == gr && col == gc {
				continue
			}
			if causesLoop(grid, gr, gc, row, col) {
				ans += 1
			}
		}
	}
	return ans
}

func copy2d[T any](src [][]T) [][]T {
	R, C := len(src), len(src[0])
	duplicate := make([][]T, R)
	for idx, row := range src {
		duplicate[idx] = make([]T, C)
		copy(duplicate[idx], row)
	}
	return duplicate
}

type State struct {
	row       int
	col       int
	direction Direction
}

func causesLoop(grid [][]rune, r int, c int, mr, mc int) bool {
	R, C := len(grid), len(grid[0])
	isSafe := func(row, col int) bool {
		return row >= 0 && row < R && col >= 0 && col < C
	}
	stateSet := make(map[State]bool)
	direction := GetDirection(grid, r, c)
	grid[mr][mc] = '#'
	gr, gc := r, c
	for isSafe(gr, gc) {
		metObstruction := grid[gr][gc] == '#'
		stateSet[State{gr, gc, direction}] = true
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
		if contains, ok := stateSet[State{gr, gc, direction}]; ok && contains {
			return true
		}
	}
	return false
}
