package number_of_islands

func numIslands(grid [][]byte) int {
	weight := len(grid[0])
	height := len(grid)
	count := 0

	var dfs func(y, x int)
	dfs = func(y, x int) {
		grid[y][x] = '0'
		for _, elem := range [][]int{{y + 1, x}, {y, x + 1}, {y - 1, x}, {y, x - 1}} {
			if elem[0] < height && elem[0] >= 0 && elem[1] < weight && elem[1] >= 0 {
				if grid[elem[0]][elem[1]] == '1' {
					dfs(elem[0], elem[1])
				}
			}
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < weight; x++ {
			if grid[y][x] == '1' {
				dfs(y, x)
				count++
			}
		}
	}

	return count
}
