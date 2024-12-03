package day3

import (
	"advent-of-code-2024/utils"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed static/input.txt
var inputContent string

func ParseInputContent() string {
	return inputContent
}

func ExecuteMulInstruction(instruction string) int64 {
	operands := strings.Split(instruction[3:], ",")
	operandARaw, operandBRaw := operands[0], operands[1]
	operandA := utils.Must(strconv.ParseInt(operandARaw[1:], 10, 64))
	operandB := utils.Must(strconv.ParseInt(operandBRaw[:len(operandBRaw)-1], 10, 64))
	return operandA * operandB
}
