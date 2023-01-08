package two_sum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, e := range nums {
		m[e] = i
	}

	for i, e := range nums {
		r := target - e

		if r == e && len(nums) == 2 {
			return []int{0, 1}
		}

		if _, ok := m[r]; ok && i != m[r] {
			return []int{i, m[r]}
		}
	}
	return []int{}
}
