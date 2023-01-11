package maximum_number_of_words_found_in_sentences

import "strings"

func mostWordsFound(sentences []string) int {
	maxWords := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for _, v := range sentences {
		maxWords = max(maxWords, len(strings.Split(v, " ")))
	}

	return maxWords
}
