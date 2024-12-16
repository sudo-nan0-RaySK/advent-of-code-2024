package day15

import "fmt"

func SolveTaskPart2() int64 {
	ans := int64(0)
	gridMeta := ParseInputContent()
	size := gridMeta.Size()
	R, C := size.Row(), 2*size.Col()
	isSafe := func(row, col int) bool {
		return row >= 0 && row < R && col >= 0 && col < C
	}
	grid := scaleGrid(gridMeta.Grid())
	start := gridMeta.Start()
	currLoc := Point{start.Row(), 2 * start.Col()}
	fmt.Printf("grid[%d][%d] = %s\n", currLoc.Row(), currLoc.Col(), string(grid[currLoc.Row()][currLoc.Col()]))
	moves := gridMeta.Moves()
	var shift func(grid [][]rune, row, col, dr, dc int) bool
	shift = func(grid [][]rune, row, col, dr, dc int) bool {
		if !isSafe(row, col) || grid[row][col] == '#' {
			return false
		}
		if grid[row][col] == '.' {
			return true
		}
		if dr == 0 || (dc == 0 && grid[row][col] != '[' && grid[row][col] != ']') {
			if shift(grid, row+dr, col+dc, dr, dc) {
				grid[row+dr][col+dc], grid[row][col] = grid[row][col], grid[row+dr][col+dc]
				return true
			} else {
				return false
			}
		} else {
			if grid[row][col] == '[' {
				if shift(grid, row+dr, col+dc, dr, dc) && shift(grid, row+dr, col+1+dc, dr, dc) {
					grid[row+dr][col+dc], grid[row][col] = grid[row][col], grid[row+dr][col+dc]
					grid[row+dr][col+1+dc], grid[row][col+1] = grid[row][col+1], grid[row+dr][col+1+dc]
					return true
				} else {
					return false
				}
			} else {
				if shift(grid, row+dr, col+dc, dr, dc) && shift(grid, row+dr, col-1+dc, dr, dc) {
					grid[row+dr][col+dc], grid[row][col] = grid[row][col], grid[row+dr][col+dc]
					grid[row+dr][col-1+dc], grid[row][col-1] = grid[row][col-1], grid[row+dr][col-1+dc]
					return true
				} else {
					return false
				}
			}
		}
	}
	//fmt.Println()
	for _, move := range moves {
		copyGrid := Copy2d(grid)
		if move == '^' { // Up
			//fmt.Println("Up")
			if shift(copyGrid, currLoc.Row(), currLoc.Col(), -1, 0) {
				currLoc = Point{currLoc.Row() - 1, currLoc.Col()}
				grid = copyGrid
			}
		} else if move == 'v' { // Down
			//fmt.Println("Down")
			if shift(copyGrid, currLoc.Row(), currLoc.Col(), 1, 0) {
				currLoc = Point{currLoc.Row() + 1, currLoc.Col()}
				grid = copyGrid
			}
		} else if move == '<' { // Left
			//fmt.Println("Left")
			if shift(copyGrid, currLoc.Row(), currLoc.Col(), 0, -1) {
				currLoc = Point{currLoc.Row(), currLoc.Col() - 1}
				grid = copyGrid
			}
		} else if move == '>' { // Right
			//fmt.Println("Right")
			if shift(copyGrid, currLoc.Row(), currLoc.Col(), 0, 1) {
				currLoc = Point{currLoc.Row(), currLoc.Col() + 1}
				grid = copyGrid
			}
		}
		//RenderState(grid)
		//time.Sleep(100 * time.Millisecond)
		//fmt.Print("\033[H\033[2J")
	}
	RenderState(grid)
	for row := 0; row < R; row++ {
		for col := 0; col < C; col++ {
			if grid[row][col] == '[' {
				ans += 100*int64(row) + int64(col)
			}
		}
	}
	return ans
}

func scaleGrid(grid [][]rune) [][]rune {
	R, C := len(grid), len(grid[0])
	newGrid := make([][]rune, R)
	for row := 0; row < R; row++ {
		c := 0
		newGrid[row] = make([]rune, 2*C)
		for col := 0; col < C; col++ {
			if grid[row][col] == '.' {
				newGrid[row][c] = '.'
				newGrid[row][c+1] = '.'
			} else if grid[row][col] == '#' {
				newGrid[row][c] = '#'
				newGrid[row][c+1] = '#'
			} else if grid[row][col] == '@' {
				newGrid[row][c] = '@'
				newGrid[row][c+1] = '.'
			} else if grid[row][col] == 'O' {
				newGrid[row][c] = '['
				newGrid[row][c+1] = ']'
			}
			c += 2
		}
	}
	return newGrid
}
