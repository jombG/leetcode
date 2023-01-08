package first_unique_character_in_a_string

func firstUniqChar(s string) int {
	m := make(map[rune]int)
	mi := make(map[rune]int)
	for i, v := range s {
		m[v]++
		mi[v] = i
	}

	for _, v := range s {
		if count := m[v]; count == 1 {
			return mi[v]
		}
	}

	return -1
}
