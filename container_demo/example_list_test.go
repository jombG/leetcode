package container_demo

import (
	"container/list"
	"fmt"
)

func Example_list() {
	l := list.New()

	l.PushBack("D")
	eA := l.PushFront("A")
	eB := l.InsertAfter("B", eA)
	l.InsertAfter("C", eB)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%s ", e.Value)
	}

	// Output:
	// A B C D

}
