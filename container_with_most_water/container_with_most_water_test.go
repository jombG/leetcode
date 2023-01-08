package container_with_most_water

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContainerWithMostWater_1(t *testing.T) {
	require.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func TestContainerWithMostWater_2(t *testing.T) {
	require.Equal(t, 1, maxArea([]int{1, 1}))
}
