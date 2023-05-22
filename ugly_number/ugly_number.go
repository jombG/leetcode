package ugly_number

func divByNumber(n, divide int) int {
	for n%divide == 0 {
		n = n / divide
	}

	return n
}

func isUgly(n int) bool {
	if n == 1 {
		return true
	}

	if n == 0 {
		return false
	}

	n = divByNumber(n, 2)
	n = divByNumber(n, 3)
	n = divByNumber(n, 5)

	if n == 1 {
		return true
	}

	return false
}
