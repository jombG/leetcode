package convert_binary_number_in_a_linked_list_to_integer

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	head = reverseList(head)

	result := 0
	pow := 2
	currentNode := head

	if currentNode == nil {
		return 0
	}

	if currentNode.Val == 1 {
		result = 1
	} else {
		result = 0
	}

	currentNode = currentNode.Next

	for currentNode != nil {
		if currentNode.Val == 1 {
			result += pow
		}
		pow *= 2
		currentNode = currentNode.Next
	}

	return result
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
