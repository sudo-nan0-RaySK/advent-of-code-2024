package day11

import (
	"advent-of-code-2024/utils"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed static/input.txt
var inputContent string

func ParseInputContent() []int64 {
	nums := strings.Split(inputContent, " ")
	stones := make([]int64, len(nums))
	for idx, num := range nums {
		stones[idx] = utils.Must(strconv.ParseInt(num, 10, 64))
	}
	return stones
}
