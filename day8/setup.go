package day8

import (
	_ "embed"
	"strings"
)

//go:embed static/input.txt
var inputContent string

func ParseInputContent() [][]rune {
	lines := strings.Split(inputContent, "\n")
	ret := make([][]rune, len(lines))
	for i, line := range lines {
		ret[i] = []rune(line)
	}
	return ret
}

type Point struct {
	row int
	col int
}

func (p *Point) GetRow() int {
	return p.row
}

func (p *Point) GetCol() int {
	return p.col
}
