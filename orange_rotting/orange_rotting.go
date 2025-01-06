package orangerotting

func orangesRotting(grid [][]int) int {
	widht := len(grid[0])
	height := len(grid)

	var dfs func(y, x, min int)
	dfs = func(y, x, min int) {
		if y < 0 || y >= height || x < 0 || x >= widht || (1 < grid[y][x] && grid[y][x] < min) || grid[y][x] == 0 {
			return
		}
		grid[y][x] = min
		for _, pair := range [][]int{{y + 1, x}, {y - 1, x}, {y, x + 1}, {y, x - 1}} {
			dfs(pair[0], pair[1], min+1)

		}
	}

	for i := 0; i < height; i++ {
		for j := 0; j < widht; j++ {
			if grid[i][j] == 2 {
				dfs(i, j, 2)
			}
		}
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	min := 2
	for i := 0; i < height; i++ {
		for j := 0; j < widht; j++ {
			if grid[i][j] == 1 {
				return -1
			}
			min = max(min, grid[i][j])
		}
	}

	return min - 2
}
