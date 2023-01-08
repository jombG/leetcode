package valid_parentheses

type Stack struct {
	List []string
}

func (s *Stack) Push(elem string) {
	s.List = append(s.List, elem)
}

func (s *Stack) Pop() string {
	headElem := s.List[len(s.List)-1]
	if len(s.List) == 1 {
		s.List = make([]string, 0, 0)
		return headElem
	}
	s.List = s.List[:len(s.List)-1]
	return headElem
}

func (s *Stack) isEmpty() bool {
	return len(s.List) == 0
}

func (s *Stack) Top() string {
	return s.List[len(s.List)-1]
}

func IsValid(s string) bool {
	stack := Stack{}
	for index := range []rune(s) {
		if stack.isEmpty() {
			stack.Push(string(s[index]))
			continue
		}
		if string(s[index]) == ")" && stack.Top() == "(" || string(s[index]) == "}" && stack.Top() == "{" || string(s[index]) == "]" && stack.Top() == "[" {
			stack.Pop()
		} else {
			stack.Push(string(s[index]))
		}
	}

	return stack.isEmpty()
}
