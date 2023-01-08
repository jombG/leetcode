package search_in_rotated_sorted_array

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := 0

	for left <= right {
		mid = (left + right) >> 1

		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			if target >= nums[left] && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}

	return -1
}
