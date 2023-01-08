package top_k_frequent_words

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHeap(t *testing.T) {
	require.ElementsMatch(t, []string{"i", "love"}, topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
}
