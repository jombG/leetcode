package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	var next = head
	var currentNode = head

	for currentNode != nil {
		next = currentNode.Next
		currentNode.Next = prev
		prev = currentNode
		currentNode = next
	}

	return prev

}
