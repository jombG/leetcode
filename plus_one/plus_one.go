package main

import "fmt"

func plusOne(digits []int) []int {
	res := make([]int, len(digits)+1)
	r := 1

	for i := len(res) - 1; i >= 1; i-- {
		if (digits[i-1] + r) > 9 {
			r = 1
			res[i] = 0
		} else {
			res[i] = digits[i-1] + r
			r = 0
		}
	}

	if r == 1 {
		res[0] = 1
	} else {
		return res[1:]
	}

	return res
}

func main() {
	fmt.Println(plusOne([]int{9, 9}))
}
