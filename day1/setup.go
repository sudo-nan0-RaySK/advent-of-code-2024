package day1

import (
	"advent-of-code-2024/utils"
	_ "embed"
	"errors"
	"strconv"
	"strings"
)

//go:embed static/input.txt
var inputContent string

func ParseInputContent() (listA, listB []int64) {
	lines := strings.Split(inputContent, "\n")
	listA = make([]int64, len(lines))
	listB = make([]int64, len(lines))
	for idx, line := range lines {
		values := strings.Split(line, "   ")
		if len(values) != 2 {
			panic(errors.New("invalid line: " + line))
		}
		listA[idx] = utils.Must(strconv.ParseInt(values[0], 10, 64))
		listB[idx] = utils.Must(strconv.ParseInt(values[1], 10, 64))
	}
	return listA, listB
}
