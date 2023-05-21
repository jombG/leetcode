package is_subsequence

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	indexT := 0
	indexS := 0
	for indexT < len(t) {
		if s[indexS] == t[indexT] {
			indexS++
		}

		if indexS == len(s) {
			return true
		}

		indexT++
	}

	return false
}

// s =
// "acb"
// t =
// "ahbgdc"
