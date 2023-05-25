package kth_missing_positive_number

func findKthPositive(arr []int, k int) int {
	counter := 0

	includes := func(arr []int, index int) bool {
		for _, e := range arr {
			if e == index {
				return true
			}
		}
		return false
	}

	for i := 1; counter != k; i++ {
		if !includes(arr, i) {
			counter++
		}
		if counter == k {
			return i
		}
	}

	return 0
}
