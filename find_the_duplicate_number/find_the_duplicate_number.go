package find_the_duplicate_number

import "fmt"

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]

	for slow != fast {
		fmt.Println(slow, fast)
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
