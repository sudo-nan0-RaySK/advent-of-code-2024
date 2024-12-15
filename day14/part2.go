package day14

import (
	"math"
)

func SolveTaskPart2() int64 {
	botInfos := ParseInputContent()
	lowestEntropy, ans := math.MaxFloat64, int64(0)
	for epoch := int64(0); epoch < 10403; epoch++ {
		entropy := 0.0
		for _, botInfo := range botInfos {
			initial := botInfo.Initial()
			velocity := botInfo.Velocity()
			x, y := initial.X(), initial.Y()
			dx, dy := velocity.X(), velocity.Y()
			x = Mod(x+epoch*dx, Width)
			y = Mod(y+epoch*dy, Height)
			entropy += math.Sqrt(math.Pow(float64(x-(Width/2)), 2) + math.Pow(float64(y-(Height/2)), 2))
		}
		if entropy < lowestEntropy {
			lowestEntropy = entropy
			ans = epoch
		}
	}
	return ans
}
