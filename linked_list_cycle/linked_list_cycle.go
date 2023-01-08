package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	runner := head
	walker := head

	for runner.Next != nil && runner.Next.Next != nil {
		runner = runner.Next.Next
		walker = walker.Next

		if runner == walker {
			return true
		}
	}

	return false
}
