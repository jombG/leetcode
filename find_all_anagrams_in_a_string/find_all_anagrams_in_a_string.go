package find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	mp := make(map[byte]int)

	if len(p) > len(s) {
		return res
	}
	for i := 0; i < len(p); i++ {
		mp[p[i]]++
	}

	for i := 0; i <= len(s)-len(p); i++ {
		newMap := make(map[byte]int)
		ang := true

		for i, v := range mp {
			newMap[i] = v
		}

		for j := i; j < i+len(p); j++ {
			newMap[s[j]]--
		}

		for _, v := range newMap {
			if v != 0 {
				ang = false
				break
			}
		}

		if ang {
			res = append(res, i)
		}
	}

	return res
}
