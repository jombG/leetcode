package container_with_most_water

//	func maxArea(height []int) int {
//		indexLeft := 0
//		indexRight := 1
//
//		max := 0
//
//		minFn := func(a, b int) int {
//			if a > b {
//				return b
//			}
//
//			return a
//		}
//
//		calcArea := func(i1, i2 int) int {
//			return (i2 - i1) * (height[i2] - height[i1])
//		}
//
//		for indexRight < len(height) {
//			if calcArea(indexLeft, indexRight) > max {
//				max = calcArea(indexLeft, indexRight)
//			} else if {
//			}
//
//		}
//
//		return 0
//	}

func maxArea(height []int) int {
	max := 0
	indexLeft := 0
	indexRight := len(height) - 1

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	calcArea := func(i1, i2 int) int {
		return (i2 - i1) * min(height[i2], height[i1])
	}

	for indexLeft < indexRight {
		area := calcArea(indexLeft, indexRight)
		if area > max {
			max = area
		}

		if height[indexLeft] < height[indexRight] {
			indexLeft++
		} else {
			indexRight--
		}
	}

	return max
}
