package day13

import "math"

func SolveTaskPart1() int64 {
	ans := int64(0)
	machineInfos := ParseInputContent()
	for _, machineInfo := range machineInfos {
		minTokens := getMinTokens(&machineInfo)
		if minTokens >= math.MaxInt64 {
			continue
		}
		ans += minTokens
	}
	return ans
}

type MemoKey struct {
	A, B int64
}

func getMinTokens(m *MachineInfo) int64 {
	memo := make(map[MemoKey]int64)
	prizeLocation := m.PrizeLocation()
	var dp func(int64, int64) int64
	btnA, btnB := m.ButtonADelta(), m.ButtonBDelta()
	dp = func(a int64, b int64) int64 {
		key := MemoKey{A: a, B: b}
		if a > 100 || b > 100 {
			return math.MaxInt64
		}

		if val, ok := memo[key]; ok {
			return val
		}

		horizontal := a*int64(btnA.X()) + b*int64(btnB.X())
		vertical := a*int64(btnA.Y()) + b*int64(btnB.Y())

		if horizontal == int64(prizeLocation.X()) && vertical == int64(prizeLocation.Y()) {
			memo[key] = 3*a + b
			return memo[key]
		}

		if horizontal > int64(prizeLocation.X()) || vertical > int64(prizeLocation.Y()) {
			memo[key] = math.MaxInt64
			return memo[key]
		}

		memo[key] = min(dp(a+1, b), dp(a, b+1))
		return memo[key]
	}
	return dp(0, 0)
}
