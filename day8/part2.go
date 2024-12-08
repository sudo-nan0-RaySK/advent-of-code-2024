package day8

import (
	"fmt"
)

func SolveTaskPart2() int64 {
	ans := int64(0)
	grid := ParseInputContent()
	R, C := len(grid), len(grid[0])
	resonantPoints := make(map[Point]bool)
	isSafe := func(point *Point) bool {
		row, col := point.GetRow(), point.GetCol()
		return row >= 0 && row < R && col >= 0 && col < C
	}
	fqToAntenaLoc := make(map[rune][]Point)
	for row := 0; row < R; row++ {
		for col := 0; col < C; col++ {
			if grid[row][col] == '.' {
				continue
			}
			if points, ok := fqToAntenaLoc[grid[row][col]]; ok {
				points = append(points, Point{row, col})
				fqToAntenaLoc[grid[row][col]] = points
			} else {
				fqToAntenaLoc[grid[row][col]] = []Point{{row, col}}
			}
		}
	}
	for _, points := range fqToAntenaLoc {
		for p1Index := 0; p1Index < len(points); p1Index++ {
			for p2Index := p1Index + 1; p2Index < len(points); p2Index++ {
				point1, point2 := points[p1Index], points[p2Index]
				diffY := point2.GetRow() - point1.GetRow()
				diffX := point2.GetCol() - point1.GetCol()

				for scalar := 2; isSafe(&Point{point1.GetRow() + scalar*diffY, point1.GetCol() + scalar*diffX}); scalar++ {
					resonantPoints[Point{point1.GetRow() + scalar*diffY, point1.GetCol() + scalar*diffX}] = true
				}

				for scalar := 1; isSafe(&Point{point1.GetRow() - scalar*diffY, point1.GetCol() - scalar*diffX}); scalar++ {
					resonantPoints[Point{point1.GetRow() - scalar*diffY, point1.GetCol() - scalar*diffX}] = true
				}
			}
		}
	}
	for point := range resonantPoints {
		r, c := point.GetRow(), point.GetCol()
		if grid[r][c] == '.' {
			grid[r][c] = '#'
		}
	}

	for row := 0; row < R; row++ {
		for col := 0; col < C; col++ {
			if grid[row][col] != '.' {
				ans += 1
			}
			fmt.Printf("%c", grid[row][col])
		}
		fmt.Println()
	}
	return ans
}
