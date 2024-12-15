package day14

import "fmt"

const (
	Height = 103
	Width  = 101
)

func SolveTaskPart1() int64 {
	botInfos := ParseInputContent()
	occupied := make(map[Point]int64)
	for _, botInfo := range botInfos {
		initial := botInfo.Initial()
		velocity := botInfo.Velocity()
		x, y := initial.X(), initial.Y()
		dx, dy := velocity.X(), velocity.Y()
		for turn := 0; turn < 7687; turn++ {
			x = Mod(x+dx, Width)
			y = Mod(y+dy, Height)
			//fmt.Printf("%d,%d\n", y, x)
		}
		occupied[Point{x, y}] += 1
	}
	ans := countQuadrant(0, 0, occupied)
	ans *= countQuadrant(0, (Width/2)+1, occupied)
	ans *= countQuadrant((Height/2)+1, 0, occupied)
	ans *= countQuadrant((Height/2)+1, (Width/2)+1, occupied)
	for row := 0; row < Height; row++ {
		for col := 0; col < Width; col++ {
			if val, ok := occupied[Point{int64(col), int64(row)}]; ok {
				fmt.Print(val)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	return ans
}

func countQuadrant(row int64, col int64, occupied map[Point]int64) int64 {
	ans := int64(0)
	for r := row; r < row+(Height/2); r++ {
		for c := col; c < col+(Width/2); c++ {
			if val, ok := occupied[Point{c, r}]; ok {
				ans += val
			}
		}
	}
	return ans
}
