package longest_substring_without_repeating_characters

type Set map[byte]struct{}

func (s Set) Put(elem byte) {
	s[elem] = struct{}{}
}

func (s Set) Check(elem byte) bool {
	_, ok := s[elem]
	return ok
}

func (s Set) Size() int {
	return len(s)
}

func (s Set) Delete(elem byte) {
	delete(s, elem)
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	left, right := 0, 0
	set := Set{}
	for right < len(s) {
		if !set.Check(s[right]) {
			set.Put(s[right])
			right++
			if set.Size() > max {
				max = set.Size()
			}
		} else {
			set.Delete(s[left])
			left++
		}
	}

	return max
}
