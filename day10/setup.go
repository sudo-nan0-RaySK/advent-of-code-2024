package day10

import (
	"advent-of-code-2024/utils"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed static/input.txt
var inputContent string

type Point struct {
	row, col int
}

func (p *Point) Row() int {
	return p.row
}

func (p *Point) Col() int {
	return p.col
}

type GridMeta struct {
	grid          [][]int
	R, C          int
	zeroPositions []Point
}

func (g *GridMeta) Grid() [][]int {
	return g.grid
}

func (g *GridMeta) ZeroPositions() []Point {
	return g.zeroPositions
}

func (g *GridMeta) Size() Point {
	return Point{g.R, g.C}
}

func ParseInputContent() *GridMeta {
	lines := strings.Split(inputContent, "\n")
	R, C := len(lines), len(lines[0])
	grid := make([][]int, R)
	zeroPositions := make([]Point, 0)
	for row, line := range lines {
		grid[row] = make([]int, C)
		for col, char := range line {
			grid[row][col] = int(utils.Must(strconv.ParseInt(string(char), 10, 64)))
			if grid[row][col] == 0 {
				zeroPositions = append(zeroPositions, Point{row, col})
			}
		}
	}
	return &GridMeta{grid, R, C, zeroPositions}
}
