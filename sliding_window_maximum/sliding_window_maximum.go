package sliding_window_maximum

type List struct {
	list []int
	head int
}

func (l *List) Push(e int) {
	l.list = append(l.list, e)
}

func (l *List) Pop() int {
	var e int
	e, l.list = l.list[len(l.list)-1], l.list[:len(l.list)-1]
	return e
}

func (l *List) Shift() {
	l.head++
}

func (l *List) Front() int {
	return l.list[l.head]
}

func (l *List) Back() int {
	return l.list[len(l.list)-1]
}

func (l *List) Len() int {
	return len(l.list) - l.head
}

func maxSlidingWindow(nums []int, k int) []int {
	var (
		queue  List
		result []int
	)

	for i := 0; i < len(nums); i++ {
		if queue.Len() > 0 && queue.Front() < i+1-k {
			queue.Shift()
		}

		for queue.Len() > 0 && nums[queue.Back()] < nums[i] {
			queue.Pop()
		}

		queue.Push(i)

		if i >= k-1 {
			result = append(result, nums[queue.Front()])
		}
	}

	return result
}
