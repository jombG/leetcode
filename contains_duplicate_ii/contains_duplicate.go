package contains_duplicate_ii

func abs(i, j int) int {
	if i > j {
		return i - j
	}

	return j - i
}

func containsNearbyDuplicate(nums []int, k int) bool {
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := dict[nums[i]]; !ok {
			dict[nums[i]] = i
		} else if abs(i, dict[nums[i]]) <= k {
			return true
		} else {
			dict[nums[i]] = i
		}
	}

	return false
}
