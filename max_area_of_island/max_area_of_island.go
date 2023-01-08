package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	maxIsland := 0
	currentIsland := 0

	height, width := len(grid), len(grid[0])

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var dfs func(y, x int)

	dfs = func(y, x int) {
		currentIsland++
		grid[y][x] = 0

		for _, elem := range [][]int{{y + 1, x}, {y, x + 1}, {y - 1, x}, {y, x - 1}} {
			if elem[0] >= 0 && elem[0] < height && elem[1] >= 0 && elem[1] < width {
				if grid[elem[0]][elem[1]] == 1 {
					dfs(elem[0], elem[1])
				}
			}
		}
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				maxIsland = max(maxIsland, currentIsland)
				currentIsland = 0
			}
		}
	}

	return maxIsland
}

func main() {
	fmt.Println(maxAreaOfIsland([][]int{
		{1},
		{1},
	}))
}
