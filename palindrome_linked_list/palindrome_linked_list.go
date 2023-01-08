package palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

type stack []int

func (s *stack) Push(elem int) {
	*s = append(*s, elem)
}

func (s *stack) Pop() int {
	list := *s
	elem := list[len(list)-1]
	*s = list[:len(list)-1]
	return elem
}

func isPalindrome(head *ListNode) bool {
	current := head
	stack := stack{}
	for current != nil {
		stack.Push(current.Val)
		current = current.Next
	}

	current = head
	for current != nil {
		if current.Val != stack.Pop() {
			return false
		}

		current = current.Next
	}

	return true
}
