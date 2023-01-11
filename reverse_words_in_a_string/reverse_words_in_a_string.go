package reverse_words_in_a_string

import "strings"

func reverseWords(s string) string {
	list := strings.Split(s, " ")
	result := make([]string, 0)

	for i := len(list) - 1; i >= 0; i-- {
		if len(list[i]) != 0 {
			if len(result) != 0 {
				result = append(result, " ")
			}
			result = append(result, list[i])
		}
	}

	return strings.Join(result, "")
}
