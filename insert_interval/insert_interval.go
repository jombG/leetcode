package insert_interval

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0, len(intervals)+1)

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for _, e := range intervals {
		if e[1] < newInterval[0] {
			res = append(res, e)
		} else if e[0] > newInterval[1] {
			res = append(res, newInterval)
			newInterval = e
		} else if e[1] >= newInterval[0] || e[0] <= newInterval[1] {
			newInterval[0] = min(newInterval[0], e[0])
			newInterval[1] = max(newInterval[1], e[1])
		}
	}
	res = append(res, newInterval)

	return res
}

// [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
// [[1,2],[3,10],[12,16]]

//Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
//Output: [[1,5],[6,9]]
