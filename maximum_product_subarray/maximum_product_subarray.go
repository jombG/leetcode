package maximum_product_subarray

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max, min, superMax := nums[0], nums[0], nums[0]

	maxFn := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	minFn := func(a, b int) int {
		if a > b {
			return b
		}

		return a
	}

	for i := 1; i < len(nums); i++ {
		temp := max

		max = maxFn(nums[i], maxFn(max*nums[i], min*nums[i]))
		min = minFn(nums[i], minFn(temp*nums[i], min*nums[i]))

		superMax = maxFn(maxFn(max, min), superMax)
	}

	return superMax
}
