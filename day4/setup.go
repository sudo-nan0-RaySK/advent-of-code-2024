package day4

import (
	_ "embed"
	"strings"
)

//go:embed static/input.txt
var inputContent string

func ParseInputContent() [][]rune {
	lines := strings.Split(inputContent, "\n")
	puzzle := make([][]rune, len(lines))
	for idx, line := range lines {
		puzzle[idx] = []rune(line)
	}
	return puzzle
}

func IsSafe(r, c, R, C int) bool {
	return r >= 0 && r < R && c >= 0 && c < C
}
