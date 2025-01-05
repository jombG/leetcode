package last_stone_weight

import "sort"

func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}
	sort.Ints(stones)
	l1, l2 := stones[len(stones)-1], stones[len(stones)-2]
	for l2 != 0 {
		if l1 == l2 {
			stones[len(stones)-1] = 0
			stones[len(stones)-2] = 0
		} else {
			stones[len(stones)-1] = l1 - l2
			stones[len(stones)-2] = 0
		}
		sort.Ints(stones)
		l1, l2 = stones[len(stones)-1], stones[len(stones)-2]
	}

	return l1
}
