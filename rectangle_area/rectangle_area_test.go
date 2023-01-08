package rectangle_area

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestComputeArea(t *testing.T) {
	// 45
	require.Equal(t, 45, ComputeArea(-3, 0, 3, 4, 0, -1, 9, 2))

	// 16
	require.Equal(t, 16, ComputeArea(-2, -2, 2, 2, -2, -2, 2, 2))
}
