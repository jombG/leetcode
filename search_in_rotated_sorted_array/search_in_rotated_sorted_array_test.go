package search_in_rotated_sorted_array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFunc(t *testing.T) {
	require.Equal(t, 4, search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
}
