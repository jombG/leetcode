package group_anagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	res := make([][]string, 0)
	for _, str := range strs {
		sort := sortString(str)

		if _, ok := m[sort]; !ok {
			m[sort] = append([]string{}, str)
		} else {
			m[sort] = append(m[sort], str)
		}
	}

	for _, v := range m {
		res = append(res, v)
	}
	return res
}

type sortRunes []rune

func (s sortRunes) Len() int {
	return len(s)
}

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sortString(str string) string {
	sr := sortRunes(str)
	sort.Sort(sr)

	return string(sr)
}
