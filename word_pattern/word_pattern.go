package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	m := make(map[string]byte)
	patternMatch := make(map[byte]string)
	words := strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	for indexer, word := range words {
		_, okS := m[word]
		_, okP := patternMatch[pattern[indexer]]

		if !okS && !okP {
			m[word] = pattern[indexer]
			patternMatch[pattern[indexer]] = word
		} else if okS == okP {
			if m[word] != pattern[indexer] {
				return false
			}
		} else {
			return false
		}
	}
	
	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
}
