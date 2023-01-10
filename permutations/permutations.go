package permutations

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	if len(nums) == 1 {
		return [][]int{nums}
	}

	result := make([][]int, 0)

	for i := range nums {
		s := make([]int, 0)
		s = append(s, nums[0:i]...)
		s = append(s, nums[i+1:]...)

		for _, v := range permute(s) {
			var tmp []int
			tmp = append(tmp, nums[i])
			tmp = append(tmp, v...)
			result = append(result, tmp)
		}
	}

	return result
}
