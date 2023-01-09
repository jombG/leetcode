package remove_linked_list_elements

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	var prevNode *ListNode = nil
	currentNode := head

	for currentNode != nil {
		if currentNode.Val == val {
			if prevNode == nil {
				head = currentNode.Next
			} else {
				prevNode.Next = currentNode.Next
			}
			currentNode = currentNode.Next
		} else {
			prevNode = currentNode
			currentNode = currentNode.Next
		}
	}

	return head
}
