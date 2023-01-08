package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	res := ""

	for i := 0; i < len(strs[0]); i++ {
		for _, v := range strs {
			if v[i] != strs[0][i] || len(v) == i {
				return res
			}
		}
		res = strs[0][:i+1]
	}

	return res
}
