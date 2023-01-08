package change_module_name

type (
	Queue[T any] struct {
		list []T
	}
)

func (q *Queue[T]) Push(elem T) {
	q.list = append(q.list, elem)
}

func (q *Queue[T]) Pop() T {
	var head T
	head, q.list = q.list[0], q.list[1:]
	return head
}

func (q *Queue[T]) Empty() bool {
	return len(q.list) == 0
}
