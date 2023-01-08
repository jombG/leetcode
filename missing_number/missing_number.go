package missing_number

func MissingNumber(nums []int) int {
	size := len(nums)
	sumNums := 0

	for _, elem := range nums {
		sumNums += elem
	}

	sum := size * (size + 1) / 2

	return sum - sumNums
}
