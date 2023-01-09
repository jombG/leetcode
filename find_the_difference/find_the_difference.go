package find_the_difference

func findTheDifference(s string, t string) byte {
	setT := make(map[rune]int)
	setS := make(map[rune]int)

	for _, v := range s {
		setS[v] += 1
	}

	for _, v := range t {
		setT[v] += 1
	}

	for _, v := range t {
		if count, ok := setS[v]; !ok {
			return byte(v)
		} else if count != setT[v] {
			return byte(v)
		}
	}

	return 0
}
