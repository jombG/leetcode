package running_sum_of_1d_array

func runningSum(nums []int) []int {
	res := make([]int, 0, len(nums))
	sum := 0

	for _, v := range nums {
		sum += v
		res = append(res, sum)
	}

	return res
}
