package main

import "fmt"

func islandPerimeter(grid [][]int) int {
	perimeter := 0

	height, width := len(grid), len(grid[0])
	calcPerimeter := func(y, x int) {
		perimeter += 4
		for _, elem := range [][]int{{y + 1, x}, {y, x + 1}, {y - 1, x}, {y, x - 1}} {
			if elem[0] < height && elem[0] >= 0 && elem[1] < width && elem[1] >= 0 {
				if grid[elem[0]][elem[1]] == 1 || grid[elem[0]][elem[1]] == -1 {
					perimeter -= 1
				}
			}
		}
	}

	var dfs func(y, x int)
	dfs = func(y, x int) {
		calcPerimeter(y, x)
		grid[y][x] = -1
		for _, elem := range [][]int{{y + 1, x}, {y, x + 1}, {y - 1, x}, {y, x - 1}} {
			if elem[0] < height && elem[0] >= 0 && elem[1] < width && elem[1] >= 0 {
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
				return perimeter
			}
		}
	}

	return perimeter
}

func main() {
	fmt.Println(islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))
}
