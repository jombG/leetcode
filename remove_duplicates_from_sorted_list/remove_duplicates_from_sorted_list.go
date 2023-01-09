package remove_duplicates_from_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prevNode, currentNode := head, head.Next

	for currentNode != nil {
		if prevNode.Val == currentNode.Val {
			prevNode.Next = currentNode.Next
			currentNode = currentNode.Next
		} else {
			prevNode = currentNode
			currentNode = currentNode.Next
		}
	}

	return head
}
