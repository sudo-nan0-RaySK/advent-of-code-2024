package day3

import (
	"advent-of-code-2024/utils"
	"fmt"
	"regexp"
)

func SolveTaskPart1() int64 {
	var pattern *regexp.Regexp
	ans := int64(0)
	instructions := ParseInputContent()
	pattern = utils.Must(regexp.Compile("mul\\([0-9]{0,3},[0-9]{0,3}\\)"))
	matches := pattern.FindAllString(instructions, -1)
	for _, instruction := range matches {
		executionResult := ExecuteMulInstruction(instruction)
		fmt.Printf("%v = %d\n", instruction, executionResult)
		ans += executionResult
	}
	return ans
}
