package sliding_window_maximum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSlidingWindowMaximum_1(t *testing.T) {
	require.ElementsMatch(t, []int{3, 3, 5, 5, 6, 7}, maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func TestSlidingWindowMaximum_2(t *testing.T) {
	require.ElementsMatch(t, []int{1}, maxSlidingWindow([]int{1}, 1))
}
