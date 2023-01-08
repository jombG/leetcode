package __sum

import "sort"

type arr []int

func (a arr) Len() int {
	return len(a)
}

func (a arr) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a arr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)

	list := arr(nums)

	sort.Sort(list)

	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := len(nums) - 1

		for j < k {
			if nums[i]+nums[k]+nums[j] == 0 {
				res = append(res, []int{nums[i], nums[k], nums[j]})
				j++
				for (j < k) && nums[j] == nums[j-1] {
					j++
				}
			} else if nums[i]+nums[k]+nums[j] < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return res
}
