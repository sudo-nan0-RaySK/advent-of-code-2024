package day18

import (
	"fmt"
	"github.com/gammazero/deque"
	"math"
)

type Step struct {
	stepsTaken int
	row, col   int
}

type Point struct {
	row, col int
}

func SolveTaskPart1() int {
	ans := math.MaxInt
	grid := ParseInputContent().Grid
	for _, row := range grid {
		for _, cell := range row {
			if cell == math.MaxInt {
				fmt.Printf("MX\t")
			} else {
				fmt.Printf("%d\t", cell)
			}
		}
		fmt.Println()
	}
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
			if isSafe(r, c) && !visited[r][c] && (grid[r][c] >= 1024) {
				visited[r][c] = true
				parent[Point{r, c}] = Point{row, col}
				q.PushBack(Step{1 + stepsTaken, r, c})
			}
		}
	}
	curr := Point{R - 1, C - 1}
	grid[curr.row][curr.col] = 0
	for curr != parent[curr] {
		curr = parent[curr]
		grid[curr.row][curr.col] = -1
	}
	fmt.Println("---")
	for _, row := range grid {
		for _, cell := range row {
			if cell == math.MaxInt {
				fmt.Printf("MX\t")
			} else if cell == -1 {
				fmt.Printf("*\t")
			} else {
				fmt.Printf("%d\t", cell)
			}
		}
		fmt.Println()
	}
	return ans
}
