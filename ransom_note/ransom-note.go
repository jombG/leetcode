package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	dict := make(map[rune]int)

	for _, r := range magazine {
		dict[r] = dict[r] + 1
	}

	for _, r := range ransomNote {
		if counter, ok := dict[r]; !ok || (ok && counter == 0) {
			return false
		} else {
			dict[r] = counter - 1
		}
	}

	return true
}