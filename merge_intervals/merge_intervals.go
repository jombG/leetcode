package merge_intervals

import "sort"

type intervalsList [][]int

func (l intervalsList) Len() int {
	return len(l)
}

func (l intervalsList) Less(i, j int) bool {
	return l[i][0] < l[j][0]
}

func (l intervalsList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func merge(intervals [][]int) [][]int {
	list := intervalsList(intervals)
	sort.Sort(list)

	res := [][]int{list[0]}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for _, v := range list[1:] {
		last := res[len(res)-1][1]

		if last >= v[0] {
			res[len(res)-1][1] = max(v[1], last)
		} else {
			res = append(res, v)
		}
	}

	return res
}
