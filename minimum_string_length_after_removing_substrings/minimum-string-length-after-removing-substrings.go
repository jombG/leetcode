package minimum_string_length_after_removing_substrings

import "strings"

func minLength(s string) int {
	for strings.Contains(s, "AB") || strings.Contains(s, "CD") {
		s = strings.Replace(s, "AB", "", 1)
		s = strings.Replace(s, "CD", "", 1)
	}

	return len(s)
}
