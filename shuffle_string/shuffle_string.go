package shuffle_string

func restoreString(s string, indices []int) string {
	res := make([]byte, len(s))

	for i, v := range s {
		res[indices[i]] = byte(v)
	}

	return string(res)
}
