package day17

import (
	"math"
	"slices"
	"strings"
)

const ProgramStr = "2,4,1,3,7,5,0,3,4,3,1,5,5,5,3,0"

func SolveTaskPart2() int64 {
	programArr := strings.Split(ProgramStr, ",")
	var dfs func(currVal int64, iteration int)
	ans := int64(math.MaxInt64)
	dfs = func(currVal int64, iteration int) {
		for bitTriple := int64(0); bitTriple < 8; bitTriple++ {
			nextNum := (currVal << 3) + bitTriple
			execution := executeProgram(nextNum)
			if !slices.Equal(execution, programArr[len(programArr)-iteration-1:]) {
				continue
			}
			if iteration == len(programArr)-1 {
				ans = min(ans, nextNum)
				return
			}
			dfs(nextNum, iteration+1)
		}
	}
	dfs(0, 0)
	return ans
}

func executeProgram(aReg int64) []string {
	program := ParseInputContent()
	regState := program.RegisterState
	regState.A = aReg
	codeSegment := program.Instructions
	for int(regState.IP) < len(codeSegment) {
		codeSegment[int(regState.IP)].Execute()
	}
	outVal := regState.PinOut
	return strings.Split(outVal[:len(outVal)-1], ",")
}
