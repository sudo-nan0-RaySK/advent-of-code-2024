package day12

import (
	_ "embed"
	"strings"
)

//go:embed static/input.txt
var inputContent string

var Directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func ParseInputContent() [][]rune {
	lines := strings.Split(inputContent, "\n")
	ret := make([][]rune, len(lines))
	for idx, line := range lines {
		ret[idx] = []rune(line)
	}
	return ret
}
