package day18

import (
	"github.com/gammazero/deque"
	"math"
)

func check(grid [][]int, threshold int) bool {
	ans := math.MaxInt
	R, C := len(grid), len(grid[0])
	q := new(deque.Deque[Step])
	visited := make([][]bool, R)
	parent := make(map[Point]Point)
	for row := 0; row < R; row++ {
		visited[row] = make([]bool, C)
		for col := 0; col < C; col++ {
			visited[row][col] = false
		}
	}
	q.PushBack(Step{0, 0, 0})
	visited[0][0] = true
	parent[Point{0, 0}] = Point{0, 0}
	for q.Len() > 0 {
		step := q.PopFront()
		stepsTaken, row, col := step.stepsTaken, step.row, step.col
		if row == R-1 && col == C-1 {
			ans = stepsTaken
			break
		}
		for _, d := range Directions {
			dr, dc := d[0], d[1]
			r, c := row+dr, col+dc
			if isSafe(r, c) && !visited[r][c] && (grid[r][c] >= threshold) {
				visited[r][c] = true
				parent[Point{r, c}] = Point{row, col}
				q.PushBack(Step{1 + stepsTaken, r, c})
			}
		}
	}
	return ans != math.MaxInt
}

func SolveTaskPart2() string {
	meta := ParseInputContent()
	grid, blockedBytes := meta.Grid, meta.BlockedBytes
	var lb, ub int
	for lb, ub = 0, len(blockedBytes)-1; lb <= ub; {
		mid := (lb + ub) >> 1
		if !check(grid, mid) {
			ub = mid - 1
		} else {
			lb = mid + 1
		}
	}
	return blockedBytes[lb-1]
}
