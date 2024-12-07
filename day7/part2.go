package day7

import (
	"math"
	"strconv"
	"sync"
	"sync/atomic"
)

func SolveTaskPart2() int64 {
	equations := ParseInputContent()
	var wg sync.WaitGroup
	var ans atomic.Int64
	wg.Add(len(equations))
	for _, equation := range equations {
		go func() {
			defer wg.Done()
			if equationSatisfiesWithConcat(&equation) {
				ans.Add(equation.ret)
			}
		}()
	}
	wg.Wait()
	return ans.Load()
}

func equationSatisfiesWithConcat(equation *Equation) bool {
	operands, ret := equation.GetOperands(), equation.GetRet()
	var satisfies func(index int, total int64) bool
	satisfies = func(index int, total int64) bool {
		if index >= len(equation.GetOperands()) {
			return total == ret
		}
		operand := operands[index]
		digits := len(strconv.FormatInt(operand, 10))
		if index > 0 {
			return satisfies(index+1, total+operand) ||
				satisfies(index+1, total*operand) ||
				satisfies(index+1, total*(int64(math.Pow10(digits)))+operand)
		} else {
			return satisfies(index+1, operand)
		}
	}
	return satisfies(0, 0)
}
