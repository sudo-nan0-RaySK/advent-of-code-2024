package day6

import (
	_ "embed"
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

func ParseInputContent() [][]rune {
	lines := strings.Split(inputContent, "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func GetDirection(grid [][]rune, r, c int) Direction {
	if grid[r][c] == '^' {
		return UP
	} else if grid[r][c] == 'V' {
		return DOWN
	} else if grid[r][c] == '>' {
		return RIGHT
	} else if grid[r][c] == '<' {
		return LEFT
	}
	panic("No guard found")
}

func FindGuardPosition(grid [][]rune) (row, col int) {
	for r, row := range grid {
		for c, cell := range row {
			if cell == 'X' {
				panic("Reused array!")
			}
			if cell == '^' || cell == 'v' || cell == '>' || cell == '<' {
				return r, c
			}
		}
	}
	panic("Guard not found!")
}
