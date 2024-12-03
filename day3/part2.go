package day3

import (
	"advent-of-code-2024/utils"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type Pair[U any, V any] struct {
	left  U
	right V
}

func SolveTaskPart2() int64 {
	var mulPattern, doPattern, dontPattern *regexp.Regexp
	ans := int64(0)
	instructions := ParseInputContent()
	mulPattern = utils.Must(regexp.Compile("mul\\([0-9]{0,3},[0-9]{0,3}\\)"))
	doPattern = utils.Must(regexp.Compile("do\\(\\)"))
	dontPattern = utils.Must(regexp.Compile("don't\\(\\)"))

	mulInstructionMatches := mulPattern.FindAllString(instructions, -1)
	mulInstructionIndexMatches := mulPattern.FindAllStringIndex(instructions, -1)
	instructionsAndIndices := zip(mulInstructionMatches, mulInstructionIndexMatches)

	doInstructionMatches := doPattern.FindAllString(instructions, -1)
	doInstructionIndexMatches := doPattern.FindAllStringIndex(instructions, -1)
	instructionsAndIndices = append(instructionsAndIndices, zip(doInstructionMatches, doInstructionIndexMatches)...)

	dontInstructionMatches := dontPattern.FindAllString(instructions, -1)
	dontInstructionIndexMatches := dontPattern.FindAllStringIndex(instructions, -1)
	instructionsAndIndices = append(instructionsAndIndices, zip(dontInstructionMatches, dontInstructionIndexMatches)...)

	sort.Slice(instructionsAndIndices, func(i, j int) bool {
		first := instructionsAndIndices[i].left
		second := instructionsAndIndices[j].left
		return first < second
	})

	mulEnabled := true

	for _, instructionAndIndex := range instructionsAndIndices {
		fmt.Println(instructionAndIndex)
		instruction := instructionAndIndex.right
		if strings.HasPrefix(instruction, "don't") {
			mulEnabled = false
			continue
		}
		if strings.HasPrefix(instruction, "do") {
			mulEnabled = true
			continue
		}
		if strings.HasPrefix(instruction, "mul") && mulEnabled {
			executionResult := ExecuteMulInstruction(instruction)
			ans += executionResult
		} else {
			fmt.Printf("Skipped %v\n", instruction)
		}
	}
	return ans
}

func zip(instructionMatches []string, instructionIndexMatches [][]int) []Pair[int, string] {
	zipped := make([]Pair[int, string], len(instructionMatches))
	for idx, instruction := range instructionMatches {
		zipped[idx] = Pair[int, string]{instructionIndexMatches[idx][0], instruction}
	}
	return zipped
}
