package guess_number_higher_or_lower

func guess(num int) int {
	return 0
}

func guessNumber(n int) int {
	left := 1
	right := n
	middle := (left + right) / 2

	for left > right {
		r := guess(middle)

		if r == 0 {
			return middle
		}

		if r == 1 {
			left = middle + 1
		}

		if r == -1 {
			right = middle - 1
		}

		middle = (left + right) / 2
	}

	return 0
}
