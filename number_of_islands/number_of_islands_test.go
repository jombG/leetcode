package number_of_islands

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNumberOfIslands(t *testing.T) {
	require.Equal(
		t,
		1,
		numIslands([][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}),
	)
}
