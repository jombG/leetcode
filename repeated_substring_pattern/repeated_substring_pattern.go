package main

import "fmt"

func repeatedSubstringPattern(s string) bool {
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i == 0 {
			sub := s[:i]
			j := i
			for j < len(s) && s[j:j+i] == sub {
				j += i
			}

			if j == len(s) {
				return true
			}
		}
	}

	return false
}

func main() {
	fmt.Println(repeatedSubstringPattern("aba"))
}
