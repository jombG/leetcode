package container_demo

import (
	"container/heap"
	"fmt"
)

func Example_heapInt() {
	h := &HeapInt{2, 3, 1, 0, 5}

	heap.Init(h)

	heap.Push(h, -1)

	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

	//	Output:
	//  -1 0 1 2 3 5
}

// Top K Frequent Words
func Example_priorityQueue() {
	list := []string{"i", "love", "leetcode", "i", "love", "coding"}

	m := make(map[string]int)
	for _, s := range list {
		m[s]++
	}

	q := &PriorityQueue{}

	for key, count := range m {
		heap.Push(q, &Item{Key: key, Count: count})
	}

	for q.Len() > 0 {
		item := heap.Pop(q).(*Item)
		fmt.Printf("%s - %d ", item.Key, item.Count)
	}

	//	Output:
	// 	i - 2 love - 2 leetcode - 1 coding - 1
}
