package day7

import (
	"sync"
	"sync/atomic"
)

func SolveTaskPart1() int64 {
	equations := ParseInputContent()
	var wg sync.WaitGroup
	var ans atomic.Int64
	wg.Add(len(equations))
	for _, equation := range equations {
		go func() {
			defer wg.Done()
			if equationSatisfies(&equation) {
				ans.Add(equation.ret)
			}
		}()
	}
	wg.Wait()
	return ans.Load()
}

func equationSatisfies(equation *Equation) bool {
	operands, ret := equation.GetOperands(), equation.GetRet()
	var satisfies func(index int, total int64) bool
	satisfies = func(index int, total int64) bool {
		if index >= len(equation.GetOperands()) {
			return total == ret
		}
		operand := operands[index]
		if index > 0 {
			return satisfies(index+1, total+operand) || satisfies(index+1, total*operand)
		} else {
			return satisfies(index+1, operand)
		}
	}
	return satisfies(0, 0)
}
