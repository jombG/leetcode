package min_remove_to_make_valid

import "strings"

type Pair struct {
	P     rune
	index int
}

type Stack struct {
	list []Pair
}

func NewStack() *Stack {
	return &Stack{
		list: make([]Pair, 0),
	}
}

func (s *Stack) Push(p Pair) {
	s.list = append(s.list, p)
}

func (s *Stack) Peek() (Pair, bool) {
	if len(s.list) == 0 {
		return Pair{}, false
	}

	return s.list[len(s.list)-1], true
}

func (s *Stack) Pop() (Pair, bool) {
	if len(s.list) == 0 {
		return Pair{}, false
	}
	p := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return p, true
}

func minRemoveToMakeValid(s string) string {
	stack := NewStack()
	for index, elem := range s {
		if elem == '(' || elem == ')' {
			if p, ok := stack.Peek(); !ok {
				stack.Push(Pair{elem, index})
			} else {
				if elem == ')' && p.P == '(' {
					_, _ = stack.Pop()
					continue
				}
				stack.Push(Pair{elem, index})
			}
		} else {
			continue
		}
	}
	toRemove := map[int]struct{}{}
	result := strings.Builder{}

	v, ok := stack.Pop()
	for ok {
		toRemove[v.index] = struct{}{}
		v, ok = stack.Pop()
	}

	for index, elem := range s {
		if _, ok := toRemove[index]; !ok {
			result.WriteRune(elem)
		}
	}

	return result.String()
}
