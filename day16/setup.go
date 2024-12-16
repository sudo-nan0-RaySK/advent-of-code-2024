package day16

import (
	_ "embed"
	godUtils "github.com/emirpasic/gods/utils"
	"strings"
)

//go:embed static/input.txt
var inputContent string

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

type GridMeta struct {
	Grid           [][]rune
	R, C           int
	StartX, StartY int
	EndX, EndY     int
}

type Point struct {
	row, col int
}

type Step struct {
	row, col  int
	score     int64
	direction Direction
}

func StepComparator(a, b Step) int {
	return godUtils.Int64Comparator(a.score, b.score)
}

func ParseInputContent() GridMeta {
	lines := strings.Split(inputContent, "\n")
	grid := make([][]rune, len(lines))
	for idx, line := range lines {
		grid[idx] = []rune(line)
	}
	var startX, startY, endX, endY int
	R, C := len(grid), len(grid[0])
	for row := 0; row < R; row++ {
		for col := 0; col < C; col++ {
			if grid[row][col] == 'S' {
				startX, startY = row, col
			}
			if grid[row][col] == 'E' {
				endX, endY = row, col
			}
		}
	}
	return GridMeta{grid, R, C, startX, startY, endX, endY}
}
