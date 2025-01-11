package main

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	res := strings.Builder{}
	currentWord1, currentWord2 := 0, 0
	switcher := false
	for currentWord1 < len(word1) || currentWord2 < len(word2) {
		if !switcher && currentWord1 < len(word1) {
			res.WriteByte(word1[currentWord1])
			currentWord1++
			if currentWord2 < len(word2) {
				switcher = true
			}
		} else if switcher && currentWord2 < len(word2) {
			res.WriteByte(word2[currentWord2])
			currentWord2++
			if currentWord1 < len(word1) {
				switcher = false
			}
		}

	}
	return res.String()
}
