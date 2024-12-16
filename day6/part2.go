package day6

import "sync"

func SolveTaskPart2() int64 {
	ans := int64(0)
	inputGrid := ParseInputContent()
	R, C := len(inputGrid), len(inputGrid[0])
	grid := Copy2d(inputGrid)
	gr, gc := FindGuardPosition(grid)
	var wg sync.WaitGroup
	wg.Add((R * C) - 1)
	var wlock sync.Mutex
	for row := 0; row < R; row++ {
		for col := 0; col < C; col++ {
			if row == gr && col == gc {
				continue
			}
			go func() {
				defer wg.Done()
				if causesLoop(grid, gr, gc, row, col) {
					defer wlock.Unlock()
					wlock.Lock()
					ans += 1
				}
			}()
		}
	}
	wg.Wait()
	return ans
}

func Copy2d[T any](src [][]T) [][]T {
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
	gr, gc := r, c
	for isSafe(gr, gc) {
		metObstruction := grid[gr][gc] == '#' || (gr == mr && gc == mc)
		stateSet[State{gr, gc, direction}] = true
		if direction == UP {
			if metObstruction {
				gr += 1
				direction = RIGHT
			} else {
				gr -= 1
			}
		} else if direction == DOWN {
			if metObstruction {
				gr -= 1
				direction = LEFT
			} else {
				gr += 1
			}
		} else if direction == LEFT {
			if metObstruction {
				gc += 1
				direction = UP
			} else {
				gc -= 1
			}
		} else if direction == RIGHT {
			if metObstruction {
				gc -= 1
				direction = DOWN
			} else {
				gc += 1
			}
		}
		if contains, ok := stateSet[State{gr, gc, direction}]; ok && contains {
			return true
		}
	}
	return false
}
