package nearest_exit_from_entrance_in_maze

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNearestExit_1(t *testing.T) {
	// 1
	require.Equal(
		t,
		1,
		NearestExit([][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}}, []int{1, 2}),
	)
}

func TestNearestExit_2(t *testing.T) {
	// 2
	require.Equal(
		t,
		2,
		NearestExit([][]byte{{'+', '+', '+'}, {'.', '.', '.'}, {'+', '+', '+'}}, []int{1, 0}),
	)
}

func TestNearestExit_3(t *testing.T) {
	// 3
	out := NearestExit([][]byte{{'.', '+'}}, []int{0, 0})
	require.Equal(
		t,
		-1,
		out,
	)
}

func TestNearestExit_4(t *testing.T) {
	// 3
	out := NearestExit([][]byte{
		{'.', '+', '+', '+', '.', '.', '.', '+', '+'},
		{'.', '.', '+', '.', '+', '.', '+', '+', '.'},
		{'.', '.', '+', '.', '.', '.', '.', '.', '.'},
		{'.', '+', '.', '.', '+', '+', '.', '+', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '+', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '+', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '+'},
		{'+', '.', '.', '.', '+', '.', '.', '.', '.'},
	}, []int{5, 6})
	require.Equal(
		t,
		2,
		out,
	)
}
