package main

import "math"

func isPowerOfThree(n int) bool {
	power := 1

	absN := int(math.Abs(float64(n)))
	for absN > power {
		power *= 3
	}
	if power == n {
		return true
	}

	return false
}

func main() {
	isPowerOfThree(27)
}
