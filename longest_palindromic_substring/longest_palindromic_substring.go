package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
}

func longestPalindrome(s string) string {
	start, end := 0, 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < len(s); i++ {
		len1 := expandString(s, i, i)
		len2 := expandString(s, i, i+1)
		lenMax := max(len2, len1)

		if lenMax > (end - start) {
			end = i + (lenMax / 2)
			start = i - ((lenMax - 1) / 2)
		}
	}

	return s[start : end+1]
}

func expandString(s string, left, right int) int {
	if s == "" || left > right {
		return 0
	}

	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return right - left - 1
}
