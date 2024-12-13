package day13

import "fmt"

const scaleAddUp = 10000000000000

func SolveTaskPart2() int64 {
	ans := int64(0)
	machineInfos := ParseInputContent()
	for idx, machineInfo := range machineInfos {
		btnADelta := machineInfo.ButtonADelta()
		btnBDelta := machineInfo.ButtonBDelta()
		prizeLoc := machineInfo.PrizeLocation()
		xa, xb := int64(btnADelta.X()), int64(btnBDelta.X())
		ya, yb := int64(btnADelta.Y()), int64(btnBDelta.Y())
		x, y := scaleAddUp+int64(prizeLoc.X()), scaleAddUp+int64(prizeLoc.Y())

		if ((xa*y - x*ya) % (xa*yb - xb*ya)) != 0 {
			continue
		}

		b := (xa*y - x*ya) / (xa*yb - xb*ya)

		if ((y - yb*b) % ya) != 0 {
			continue
		}

		a := (y - yb*b) / ya

		if a < 0 || b < 0 {
			continue
		}

		fmt.Printf("Using %d A presses and %d B presses for #%d \n", a, b, idx)

		ans += a*3 + b
	}
	return ans
}
