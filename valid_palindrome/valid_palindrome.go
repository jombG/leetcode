package valid_palindrome

import "strings"

func isPalindrome(s string) bool {
	str := strings.ToLower(strings.ReplaceAll(s, " ", ""))
	length := len(str)
	left, right := 0, length-1
	for left < right {
		if !ignore(str[left]) {
			left++
			continue
		}

		if !ignore(str[right]) {
			right--
			continue
		}

		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func ignore(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
}
