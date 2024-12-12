package day12

func SolveTaskPart2() int64 {
	ans := int64(0)
	grid := ParseInputContent()
	for ord := 65; ord < 91; ord++ {
		rnOrd := rune(ord)
		contribution := calcTotalDimensionsForAlphabet2(rnOrd, grid)
		ans += contribution
	}
	return ans
}

func calcTotalDimensionsForAlphabet2(ord rune, grid [][]rune) int64 {
	ans := int64(0)
	R, C := len(grid), len(grid[0])
	visited := make([][]int, R)
	for idx := 0; idx < R; idx++ {
		visited[idx] = make([]int, C)
	}

	isSafe := func(row, col int) bool {
		return row >= 0 && row < R && col >= 0 && col < C
	}

	horizontalSideCount := func(identifier int) int64 {
		sides := int64(0)
		// Upside
		for row := 0; row < R; row++ {
			lastState := false
			for col := 0; col < C; col++ {
				currState := visited[row][col] == identifier && (!isSafe(row-1, col) || visited[row-1][col] != identifier)
				if currState && !lastState {
					sides += 1
				}
				lastState = currState
			}
		}

		// Downside
		for row := 0; row < R; row++ {
			lastState := false
			for col := 0; col < C; col++ {
				currState := visited[row][col] == identifier && (!isSafe(row+1, col) || visited[row+1][col] != identifier)
				if currState && !lastState {
					sides += 1
				}
				lastState = currState
			}
		}
		return sides
	}

	verticalSideCount := func(identifier int) int64 {
		sides := int64(0)

		// Left
		for col := 0; col < C; col++ {
			lastState := false
			for row := 0; row < R; row++ {
				currState := visited[row][col] == identifier && (!isSafe(row, col-1) || visited[row][col-1] != identifier)
				if currState && !lastState {
					sides += 1
				}
				lastState = currState
			}
		}

		// Right
		for col := 0; col < C; col++ {
			lastState := false
			for row := 0; row < R; row++ {
				currState := visited[row][col] == identifier && (!isSafe(row, col+1) || visited[row][col+1] != identifier)
				if currState && !lastState {
					sides += 1
				}
				lastState = currState
			}
		}

		return sides
	}

	var traverseRegion func(int, int, rune, *int64, int)
	traverseRegion = func(r int, c int, ord rune, area *int64, itemIdentifier int) {
		*area += 1
		for _, d := range Directions {
			dr, dc := d[0], d[1]
			row, col := r+dr, c+dc
			if !isSafe(row, col) || grid[row][col] != ord {
				continue
			}
			if visited[row][col] == 0 {
				visited[row][col] = itemIdentifier
				traverseRegion(row, col, ord, area, itemIdentifier)
			}
		}
	}

	item := 1

	for row := 0; row < R; row++ {
		for col := 0; col < C; col++ {
			if grid[row][col] == ord && visited[row][col] == 0 {
				visited[row][col] = item
				area := int64(0)
				traverseRegion(row, col, ord, &area, item)
				perimeter := horizontalSideCount(item) + verticalSideCount(item)
				ans += area * perimeter
				item += 1
			}
		}
	}

	return ans
}
