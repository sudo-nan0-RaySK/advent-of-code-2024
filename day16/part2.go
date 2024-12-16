package day16

const maxScore = int64(102488)

type MemoKey struct {
	r, c  int
	d     Direction
	score int64
}

func SolveTaskPart2() int64 {
	ans := int64(0)
	gridInfo := ParseInputContent()
	R, C := gridInfo.R, gridInfo.C
	grid := gridInfo.Grid
	isSafe := func(r, c int) bool {
		return (r >= 0 && r < R && c >= 0 && c < C) && grid[r][c] != '#'
	}
	startX, startY := gridInfo.StartX, gridInfo.StartY
	endX, endY := gridInfo.EndX, gridInfo.EndY
	visited := make([][]bool, R)
	spcl := make([][]bool, R)
	for r := 0; r < R; r++ {
		visited[r] = make([]bool, C)
		spcl[r] = make([]bool, C)
		for c := 0; c < C; c++ {
			visited[r][c] = false
			spcl[r][c] = false
		}
	}
	memo := make(map[MemoKey]bool)
	var dfs func(r int, c int, d Direction, score int64) bool
	dfs = func(r int, c int, d Direction, score int64) bool {

		key := MemoKey{r, c, d, score}

		if val, ok := memo[key]; ok {
			return val
		}

		if score > maxScore {
			return false
		}

		if r == endX && c == endY {
			return true
		}

		canReach := false

		if d != UP && !visited[r+1][c] && isSafe(r+1, c) { // Going Down
			visited[r+1][c] = true
			if d == DOWN {
				canReach = dfs(r+1, c, DOWN, score+1) || canReach
			} else {
				canReach = dfs(r+1, c, DOWN, score+1001) || canReach
			}
			visited[r+1][c] = false
		}

		if d != RIGHT && !visited[r][c-1] && isSafe(r, c-1) { // Going Left
			visited[r][c-1] = true
			if d == LEFT {
				canReach = dfs(r, c-1, LEFT, score+1) || canReach
			} else {
				canReach = dfs(r, c-1, LEFT, score+1001) || canReach
			}
			visited[r][c-1] = false
		}

		if d != LEFT && !visited[r][c+1] && isSafe(r, c+1) { // Going Right
			visited[r][c+1] = true
			if d == RIGHT {
				canReach = dfs(r, c+1, RIGHT, score+1) || canReach
			} else {
				canReach = dfs(r, c+1, RIGHT, score+1001) || canReach
			}
			visited[r][c+1] = false
		}

		if d != DOWN && !visited[r-1][c] && isSafe(r-1, c) { // Going Up
			visited[r-1][c] = true
			if d == UP {
				canReach = dfs(r-1, c, UP, score+1) || canReach
			} else {
				canReach = dfs(r-1, c, UP, score+1001) || canReach
			}
			visited[r-1][c] = false
		}

		if canReach {
			spcl[r][c] = true
		}
		memo[key] = canReach
		return memo[key]
	}
	visited[startX][startY] = true
	dfs(startX, startY, RIGHT, 0)
	visited[startX][startY] = false
	spcl[endX][endY] = true
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if spcl[r][c] {
				ans += 1
			} else {
			}
		}
	}
	return ans
}
