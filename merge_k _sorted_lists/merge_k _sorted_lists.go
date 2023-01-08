package merge_k__sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	result := lists[0]
	for i := 1; i < len(lists); i++ {
		result = mergeList(result, lists[i])
	}

	return result
}

func mergeList(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head

	for l1 != nil && l2 != nil {
		v1, v2 := l1.Val, l2.Val
		v := 0
		if v1 > v2 {
			v = v2
			l2 = l2.Next
		} else {
			v = v1
			l1 = l1.Next
		}

		current.Next = &ListNode{Val: v}
		current = current.Next
	}

	if l1 != nil && l2 == nil {
		for l1 != nil {
			current.Next = &ListNode{Val: l1.Val}
			current = current.Next
			l1 = l1.Next
		}
	}

	if l2 != nil && l1 == nil {
		for l2 != nil {
			current.Next = &ListNode{Val: l2.Val}
			current = current.Next
			l2 = l2.Next
		}
	}

	return head.Next
}
