package day15

func SolveTaskPart1() int64 {
	ans := int64(0)
	gridMeta := ParseInputContent()
	size := gridMeta.Size()
	R, C := size.Row(), size.Col()
	isSafe := func(row, col int) bool {
		return row >= 0 && row < R && col >= 0 && col < C
	}
	grid := gridMeta.Grid()
	start := gridMeta.Start()
	currLoc := start
	moves := gridMeta.Moves()
	var shift func(row, col, dr, dc int) bool
	shift = func(row, col, dr, dc int) bool {
		if !isSafe(row, col) || grid[row][col] == '#' {
			return false
		}
		if grid[row][col] == '.' {
			return true
		}
		if shift(row+dr, col+dc, dr, dc) {
			grid[row+dr][col+dc] = grid[row][col]
			grid[row][col] = '.'
			return true
		} else {
			return false
		}
	}
	for _, move := range moves {
		if move == '^' { // Up
			//fmt.Println("Up")
			if shift(currLoc.Row(), currLoc.Col(), -1, 0) {
				currLoc = Point{currLoc.Row() - 1, currLoc.Col()}
			}
		} else if move == 'v' { // Down
			//fmt.Println("Down")
			if shift(currLoc.Row(), currLoc.Col(), 1, 0) {
				currLoc = Point{currLoc.Row() + 1, currLoc.Col()}
			}
		} else if move == '<' { // Left
			//fmt.Println("Left")
			if shift(currLoc.Row(), currLoc.Col(), 0, -1) {
				currLoc = Point{currLoc.Row(), currLoc.Col() - 1}
			}
		} else if move == '>' { // Right
			//fmt.Println("Right")
			if shift(currLoc.Row(), currLoc.Col(), 0, 1) {
				currLoc = Point{currLoc.Row(), currLoc.Col() + 1}
			}
		}
		//RenderState(grid)
	}
	for row := 0; row < R; row++ {
		for col := 0; col < C; col++ {
			if grid[row][col] == 'O' {
				ans += 100*int64(row) + int64(col)
			}
		}
	}
	return ans
}
