package find_minimum_in_rotated_sorted_array

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	res := nums[0]

	min := func(a, b int) int {
		if a > b {
			return b
		}

		return a
	}

	for left <= right {
		if nums[left] < nums[right] {
			res = min(res, nums[left])
			break
		}

		mid := (left + right) / 2
		res = min(nums[mid], res)

		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return res
}
