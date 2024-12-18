package day18

import (
	"advent-of-code-2024/utils"
	_ "embed"
	"math"
	"strconv"
	"strings"
)

//go:embed static/input.txt
var InputContent string

var Directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func isSafe(r, c int) bool {
	return r >= 0 && r < 71 && c >= 0 && c < 71
}

type GridMeta struct {
	Grid         [][]int
	BlockedBytes []string
}

func ParseInputContent() GridMeta {
	R, C := 71, 71
	grid := make([][]int, R)
	for row := 0; row < R; row++ {
		grid[row] = make([]int, C)
		for col := 0; col < C; col++ {
			grid[row][col] = math.MaxInt
		}
	}
	lines := strings.Split(InputContent, "\n")
	for idx, line := range lines {
		points := strings.Split(line, ",")
		c, r := utils.Must(strconv.Atoi(points[0])), utils.Must(strconv.Atoi(points[1]))
		grid[r][c] = min(grid[r][c], idx)
	}
	return GridMeta{grid, lines}
}
