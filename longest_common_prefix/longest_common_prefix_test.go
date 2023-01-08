package longest_common_prefix

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLongest(t *testing.T) {
	require.Equal(t, "fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
