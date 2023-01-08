package valid_anagram

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidAnagram_1(t *testing.T) {
	require.True(t, validAnagram("anagram", "nagaram"))
}

func TestValidAnagram_2(t *testing.T) {
	require.False(t, validAnagram("rat", "car"))
}

func TestValidAnagram_3(t *testing.T) {
	require.False(t, validAnagram("ab", "a"))
}

func TestValidAnagram_4(t *testing.T) {
	require.False(t, validAnagram("aa", "a"))
}
