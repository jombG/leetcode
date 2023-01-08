package intersection_of_two_arrays_ii

func intersect(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	m := make(map[int]int)

	if len(nums1) > len(nums2) {
		for _, v := range nums2 {
			m[v]++
		}

		for _, v := range nums1 {
			if _, ok := m[v]; ok {
				m[v]--
				if m[v] == 0 {
					delete(m, v)
				}
				res = append(res, v)
			}
		}
	} else {
		for _, v := range nums1 {
			m[v]++
		}

		for _, v := range nums2 {
			if _, ok := m[v]; ok {
				m[v]--
				if m[v] == 0 {
					delete(m, v)
				}
				res = append(res, v)
			}
		}
	}

	return res
}
