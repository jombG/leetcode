package palindromic_substrings

func countSubstrings(s string) int {
	counter := 0
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}

	for i := 0; i < len(s); i++ {
		dp[i][i] = true
		counter++
	}

	for i := 0; i < len(s)-1; i++ {
		dp[i][i+1] = s[i] == s[i+1]
		if dp[i][i+1] {
			counter++
		}
	}

	for l := 3; l <= len(s); l++ {
		for i := 0; ; i++ {
			j := l + i - 1
			if j >= len(s) {
				break
			}
			dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			if dp[i][j] {
				counter++
			}
		}
	}

	return counter
}
