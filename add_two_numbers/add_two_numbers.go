package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r := 0
	head := &ListNode{}
	current := head

	for l1 != nil || l2 != nil || r != 0 {
		n1, n2 := 0, 0

		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		num := n1 + n2 + r
		r = num / 10
		current.Next = &ListNode{Val: num % 10}
		current = current.Next
	}

	return head.Next
}
