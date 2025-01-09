package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	m := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	res := []string{}
	strLen := len(digits)
	var dfs func(str string, index int)

	dfs = func(str string, index int) {
		if len(str) == strLen {
			res = append(res, str)
			return
		}
		for _, v := range m[digits[index]] {
			dfs(str+v, index+1)
		}
	}

	for _, v := range m[digits[0]] {
		dfs(v, 1)
	}

	return res
}
