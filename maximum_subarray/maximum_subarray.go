package maximum_subarray

import "math"

func maxSubArray(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var dc func(l, r int) int

	dc = func(l, r int) int {
		if l == r {
			return nums[l]
		}

		mid := (l + r) / 2
		sum := 0

		leftMax := math.MinInt
		for i := mid; i >= l; i-- {
			sum += nums[i]
			if sum > leftMax {
				leftMax = sum
			}
		}

		rightMax := math.MinInt
		sum = 0
		for i := mid + 1; i <= r; i++ {
			sum += nums[i]
			if sum > rightMax {
				rightMax = sum
			}
		}

		maxLeftRight := max(dc(l, mid), dc(mid+1, r))
		return max(maxLeftRight, leftMax+rightMax)
	}

	if len(nums) == 0 {
		return 0
	}

	return dc(0, len(nums)-1)
}
