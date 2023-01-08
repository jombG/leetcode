package binary_search

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBinarySearch_1(t *testing.T) {
	require.Equal(t, -1, BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 2))
}

func TestBinarySearch_2(t *testing.T) {
	require.Equal(t, 4, BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 9))
}

func TestBinarySearch_3(t *testing.T) {
	require.Equal(t, -1, BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 13))
}
