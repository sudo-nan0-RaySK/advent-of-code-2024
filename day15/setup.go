package day15

import (
	_ "embed"
	"fmt"
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
	grid  [][]rune
	R, C  int
	start Point
	moves []rune
}

func (g *GridMeta) Grid() [][]rune {
	return g.grid
}

func (g *GridMeta) Size() Point {
	return Point{g.R, g.C}
}

func (g *GridMeta) Start() Point {
	return g.start
}

func (g *GridMeta) Moves() []rune {
	return g.moves
}

func ParseInputContent() *GridMeta {
	parts := strings.Split(inputContent, "\n\n")
	gridLines := strings.Split(parts[0], "\n")
	grid := make([][]rune, len(gridLines))
	var startPt Point
	for row, line := range gridLines {
		grid[row] = []rune(line)
		for col := 0; col < len(line); col++ {
			if grid[row][col] == '@' {
				startPt = Point{row, col}
			}
		}
	}
	moves := make([]rune, 0)
	moveLines := strings.Split(parts[1], "\n")
	for _, line := range moveLines {
		moves = append(moves, []rune(line)...)
	}
	return &GridMeta{grid: grid, R: len(gridLines), C: len(gridLines[0]), start: startPt, moves: moves}
}

func RenderState(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
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
