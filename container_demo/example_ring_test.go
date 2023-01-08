package container_demo

import (
	"container/ring"
	"fmt"
)

func ExampleRing_Next() {
	r := ring.New(5)

	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}
	for j := 0; j < n; j++ {
		fmt.Printf("%d ", r.Value)
		r = r.Next()
	}

	// Output:
	// 0 1 2 3 4
}
