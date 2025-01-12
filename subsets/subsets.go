package main

type Set[T comparable] struct {
	set map[T]struct{}
}

func (s *Set[T]) Add(e T) {
	s.set[e] = struct{}{}
}

func (s *Set[T]) Copy() Set[T] {
	newSet := make(map[T]struct{}, len(s.set))
	for e := range s.set {
		newSet[e] = struct{}{}
	}

	return Set[T]{
		set: newSet,
	}
}

func (s *Set[T]) ToList() []T {
	list := make([]T, 0, len(s.set))
	for e := range s.set {
		list = append(list, e)
	}

	return list
}

func subsets(nums []int) [][]int {
	res := []Set[int]{
		{},
	}

	for _, num := range nums {
		size := len(res)
		for i := 0; i < size; i++ {
			newSet := res[i].Copy()
			newSet.Add(num)
			res = append(res, newSet)
		}
	}

	list := make([][]int, 0, len(res))
	for _, s := range res {
		list = append(list, s.ToList())
	}

	return list
}
