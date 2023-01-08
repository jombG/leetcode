package concatenation_of_array

func getConcatenation(nums []int) []int {
	length := len(nums)
	res := make([]int, 2*length)

	for i := 0; i < length; i++ {
		res[i], res[i+length] = nums[i], nums[i]
	}

	return res
}
