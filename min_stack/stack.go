package min_stack

import "math"

type MinStack struct {
	min      int
	minStack []int
	list     []int
}

func Constructor() MinStack {
	return MinStack{
		min:      math.MaxInt64,
		minStack: []int{},
		list:     []int{},
	}
}

func (s *MinStack) Push(val int) {
	if val <= s.min {
		s.min = val
		s.minStack = append(s.minStack, val)
	}
	s.list = append(s.list, val)
}

func (s *MinStack) Pop() {
	head := s.list[len(s.list)-1]
	if head > s.min {
		if len(s.list) == 1 {
			s.list = []int{}
			s.min = math.MaxInt64
			return
		}
		s.list = s.list[:len(s.list)-1]
		return
	}
	if len(s.minStack) == 1 {
		s.min = math.MaxInt64
		s.minStack = []int{}
	} else {
		s.min = s.minStack[len(s.minStack)-2]
		s.minStack = s.minStack[:len(s.minStack)-1]
	}
	if len(s.list) == 1 {
		s.list = []int{}
	} else {
		s.list = s.list[:len(s.list)-1]
	}
}

func (s *MinStack) Top() int {
	return s.list[len(s.list)-1]
}

func (s *MinStack) GetMin() int {
	return s.min
}
