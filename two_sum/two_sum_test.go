package two_sum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTwoSum_1(t *testing.T) {
	require.ElementsMatch(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
}

func TestTwoSum_2(t *testing.T) {
	require.ElementsMatch(t, []int{1, 2}, twoSum([]int{3, 2, 4}, 6))
}

func TestTwoSum_3(t *testing.T) {
	require.ElementsMatch(t, []int{0, 2}, twoSum([]int{3, 2, 3}, 6))
}
