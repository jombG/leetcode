package longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	maxCounter := 0
	set := make(map[int]bool)

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	for _, v := range nums {
		set[v] = true
	}

	for _, v := range nums {
		if !set[v-1] {
			counter := 0
			for set[v+counter] {
				counter++
			}
			maxCounter = max(counter, maxCounter)
		}
	}

	return maxCounter
}
