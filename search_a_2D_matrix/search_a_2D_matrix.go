package search_a_2D_matrix

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if f := binarySearch(matrix[i], target); f != -1 {
			return true
		}
	}
	return false
}

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	var mid int
	for left <= right {
		mid = (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if left == right && nums[mid] == target {
		return mid
	}

	return -1
}
