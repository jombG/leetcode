package top_k_frequent_words

import (
	"container/heap"
	"strings"
)

func topKFrequent(nums []string, k int) []string {
	m := make(map[string]int)
	for _, s := range nums {
		m[s]++
	}
	q := PriorityQueue{}
	for key, count := range m {
		heap.Push(&q, &Item{key: key, count: count})
	}
	var result []string
	for len(result) < k {
		item := heap.Pop(&q).(*Item)
		result = append(result, item.key)
	}
	return result
}

// Item define
type Item struct {
	key   string
	count int
}

// A PriorityQueue implements heap.Interface and holds Items.

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].count == pq[j].count {
		return strings.Compare(pq[i].key, pq[j].key) > 0
	}
	return pq[i].count > pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push define
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

// Pop define
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}
