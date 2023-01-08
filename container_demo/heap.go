package container_demo

import (
	"container/heap"
	"sort"
)

type Interface interface {
	sort.Interface
	Push(x any) // add x as element Len()
	Pop() any   // remove and return element Len() - 1.
}

var _ heap.Interface = (*HeapInt)(nil)

type HeapInt []int

func (h HeapInt) Len() int {
	return len(h)
}

func (h HeapInt) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h HeapInt) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HeapInt) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *HeapInt) Pop() any {
	old := *h
	length := len(old)
	x := old[length-1]
	*h = old[:length-1]

	return x
}

type (
	Item struct {
		Key   string
		Count int
	}

	PriorityQueue []*Item
)

var _ heap.Interface = (*PriorityQueue)(nil)

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].Count > p[j].Count
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*p = append(*p, item)
}

func (p *PriorityQueue) Pop() any {
	old := *p
	length := len(old)
	item := old[length-1]
	*p = old[:length-1]

	return item
}
