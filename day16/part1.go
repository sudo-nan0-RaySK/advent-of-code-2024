package day16

import (
	"github.com/emirpasic/gods/v2/trees/binaryheap"
	"math"
)

func SolveTaskPart1() int64 {
	gridInfo := ParseInputContent()
	R, C := gridInfo.R, gridInfo.C
	grid := gridInfo.Grid
	isSafe := func(r, c int) bool {
		return (r >= 0 && r < R && c >= 0 && c < C) && grid[r][c] != '#'
	}
	startX, startY := gridInfo.StartX, gridInfo.StartY
	endX, endY := gridInfo.EndX, gridInfo.EndY
	h := binaryheap.NewWith(StepComparator)
	h.Push(Step{startX, startY, 0, RIGHT})
	distance := make([][]int64, R)
	for r := 0; r < R; r++ {
		distance[r] = make([]int64, C)
		for c := 0; c < C; c++ {
			distance[r][c] = math.MaxInt64
		}
	}
	for !h.Empty() {
		step, _ := h.Pop()
		r, c := step.row, step.col
		if r == endX && c == endY {
			break
		}
		direction := step.direction
		currScore := step.score

		if direction != UP { // Going Down
			var addUp int64
			if direction != DOWN {
				addUp = 1001
			} else {
				addUp = 1
			}

			row, col := r+1, c
			if isSafe(row, col) && distance[row][col] > currScore+addUp {
				distance[row][col] = currScore + addUp
				h.Push(Step{row, col, distance[row][col], DOWN})
			}
		}

		if direction != LEFT { // Going Right
			var addUp int64
			if direction != RIGHT {
				addUp = 1001
			} else {
				addUp = 1
			}

			row, col := r, c+1
			if isSafe(row, col) && distance[row][col] > currScore+addUp {
				distance[row][col] = currScore + addUp
				h.Push(Step{row, col, distance[row][col], RIGHT})
			}
		}

		if direction != RIGHT { // Going Left
			var addUp int64
			if direction != LEFT {
				addUp = 1001
			} else {
				addUp = 1
			}

			row, col := r, c-1
			if isSafe(row, col) && distance[row][col] > currScore+addUp {
				distance[row][col] = currScore + addUp
				h.Push(Step{row, col, distance[row][col], LEFT})
			}
		}

		if direction != DOWN { // Going Up
			var addUp int64
			if direction != UP {
				addUp = 1001
			} else {
				addUp = 1
			}

			row, col := r-1, c
			if isSafe(row, col) && distance[row][col] > currScore+addUp {
				distance[row][col] = currScore + addUp
				h.Push(Step{row, col, distance[row][col], UP})
			}
		}
	}
	return distance[endX][endY]
}
