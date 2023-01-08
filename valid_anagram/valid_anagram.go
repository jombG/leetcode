package valid_anagram

func validAnagram(s string, t string) bool {
	mS := make(map[rune]int)
	mT := make(map[rune]int)

	for _, elem := range s {
		if _, ok := mS[elem]; !ok {
			mS[elem] = 1
		} else {
			mS[elem] = mS[elem] + 1
		}
	}

	for _, elem := range t {
		if _, ok := mT[elem]; !ok {
			mT[elem] = 1
		} else {
			mT[elem] = mT[elem] + 1
		}
	}

	if len(mS) != len(mT) {
		return false
	}

	for key, vS := range mS {
		vT, ok := mT[key]
		if !ok || vT != vS {
			return false
		}

	}

	return true
}
