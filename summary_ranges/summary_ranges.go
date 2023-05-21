package summary_ranges

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	out := make([]string, 0, len(nums))
	l, r := 0, 0
	length := 0
	for r = 1; r < len(nums); r++ {
		if nums[r]-nums[r-1] == 1 {
			length++
			continue
		}

		if length != 0 {
			out = append(out, fmt.Sprintf("%d->%d", nums[l], nums[r-1]))
			length = 0
		} else {
			out = append(out, fmt.Sprintf("%d", nums[l]))
		}
		l = r
	}

	if length != 0 {
		out = append(out, fmt.Sprintf("%d->%d", nums[l], nums[r-1]))
	} else {
		out = append(out, fmt.Sprintf("%d", nums[l]))
	}

	return out
}

// 0,2,3,4,6,8,9

// 0
// 2
