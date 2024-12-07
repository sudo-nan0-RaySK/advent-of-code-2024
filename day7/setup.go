package day7

import (
	"advent-of-code-2024/utils"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed static/input.txt
var inputContent string

type Equation struct {
	operands []int64
	ret      int64
}

func (e *Equation) GetOperands() []int64 {
	return e.operands
}

func (e *Equation) GetRet() int64 {
	return e.ret
}

func ParseInputContent() []Equation {
	lines := strings.Split(inputContent, "\n")
	equations := make([]Equation, len(lines))
	for idx, line := range lines {
		equations[idx] = parseEquation(line)
	}
	return equations
}

func parseEquation(line string) Equation {
	parts := strings.Split(line, ":")
	ret := utils.Must(strconv.ParseInt(parts[0], 10, 64))
	operandsRaw := strings.Split(strings.Trim(parts[1], " "), " ")
	operands := make([]int64, len(operandsRaw))
	for idx, operandRaw := range operandsRaw {
		operand := utils.Must(strconv.ParseInt(operandRaw, 10, 64))
		operands[idx] = operand
	}
	return Equation{operands: operands, ret: ret}
}
