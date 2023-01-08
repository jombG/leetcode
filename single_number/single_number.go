package single_number

func singleNumberHash(nums []int) int {
	m := make(map[int]int)

	for _, e := range nums {
		m[e] = m[e] + 1
	}

	for _, e := range nums {
		if v := m[e]; v == 1 {
			return e
		}
	}
	return 0
}

func singleNumberXor(nums []int) int {
	var r int

	for _, e := range nums {
		r = r ^ e
	}

	return r
}
