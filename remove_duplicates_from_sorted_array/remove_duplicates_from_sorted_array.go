package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	counter := 1
	indexer := 1
	for i := 1; i < len(nums); {
		if nums[i-1] != nums[i] {
			nums[indexer] = nums[i]
			indexer++
			counter++
		}
	}

	return counter
}
