package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	f := make(map[byte]byte)
	fI := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		eS, okS := f[s[i]]
		eT, okT := fI[t[i]]

		if !okS && !okT {
			f[s[i]] = t[i]
			fI[t[i]] = s[i]

		} else if okS && okT {
			if eS != t[i] || eT != s[i] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isIsomorphic("paper", "title"))
}
