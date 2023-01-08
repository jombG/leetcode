package main

import (
	"fmt"
	"math"
)

type line [2]float64

func findLine(x0, x1, y0, y1 float64) line {
	if y0 == y1 {
		return line{0, y0}
	} else if x0 == x1 {
		return line{x0, math.MaxFloat64}
	}

	slope := (y1 - y0) / (x1 - x0)
	intercept := y0 - slope*x0

	return line{slope, intercept}
}

func maxPoints(points [][]int) int {
	if len(points) < 1 {
		return 1
	}

	lines := make(map[line]map[[2]int]struct{})

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			line := findLine(
				float64(points[i][0]),
				float64(points[j][0]),
				float64(points[i][1]),
				float64(points[j][1]),
			)
			if _, ok := lines[line]; !ok {
				lines[line] = make(map[[2]int]struct{})
				lines[line][[2]int{points[i][0], points[i][1]}] = struct{}{}
				lines[line][[2]int{points[j][0], points[j][1]}] = struct{}{}
			} else {
				lines[line][[2]int{points[i][0], points[i][1]}] = struct{}{}
				lines[line][[2]int{points[j][0], points[j][1]}] = struct{}{}
			}
		}
	}

	maxPoints := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	for _, v := range lines {
		maxPoints = max(maxPoints, len(v))
	}

	return maxPoints
}

func main() {
	fmt.Println(
		maxPoints([][]int{
			{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4},
		}),
	)
}
