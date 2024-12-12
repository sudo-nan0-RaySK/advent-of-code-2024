package day12

import (
	"sync"
	"sync/atomic"
)

func SolveTaskPart1() int64 {
	var ans atomic.Int64
	ans.Store(0)
	grid := ParseInputContent()
	var wg sync.WaitGroup
	wg.Add(26)
	for ord := 65; ord < 91; ord++ {
		rnOrd := rune(ord)
		go func() {
			defer wg.Done()
			contribution := calcTotalDimensionsForAlphabet(rnOrd, grid)
			ans.Add(contribution)
		}()
	}
	wg.Wait()
	return ans.Load()
}

func calcTotalDimensionsForAlphabet(ord rune, grid [][]rune) int64 {
	ans := int64(0)
	R, C := len(grid), len(grid[0])
	visited := make([][]bool, R)
	for idx := 0; idx < R; idx++ {
		visited[idx] = make([]bool, C)
	}

	isSafe := func(row, col int) bool {
		return row >= 0 && row < R && col >= 0 && col < C
	}

	var traverseRegion func(int, int, rune, *int64, *int64)
	traverseRegion = func(r int, c int, ord rune, area *int64, perimeter *int64) {
		*area += 1
		for _, d := range Directions {
			dr, dc := d[0], d[1]
			row, col := r+dr, c+dc
			if !isSafe(row, col) || grid[row][col] != ord {
				*perimeter += 1
				continue
			}
			if !visited[row][col] {
				visited[row][col] = true
				traverseRegion(row, col, ord, area, perimeter)
			}
		}
	}

	for row := 0; row < R; row++ {
		for col := 0; col < C; col++ {
			if grid[row][col] == ord && !visited[row][col] {
				visited[row][col] = true
				area, perimeter := int64(0), int64(0)
				traverseRegion(row, col, ord, &area, &perimeter)
				ans += area * perimeter
			}
		}
	}

	return ans
}
