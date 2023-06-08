package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	mSize := len(matrix)
	nSize := len(matrix[0])

	maxSlen := 0

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([]int, nSize+1)
	prev := 0

	for i := 1; i <= mSize; i++ {
		for j := 1; j <= nSize; j++ {
			temp := dp[j]
			if matrix[i-1][j-1] == '1' {
				dp[j] = min(dp[j-1], min(prev, dp[j])) + 1
				maxSlen = max(maxSlen, dp[j])
			} else {
				dp[j] = 0
			}
			prev = temp
		}
	}

	return maxSlen * maxSlen
}

func main() {
	fmt.Println(
		maximalSquare([][]byte{
			{'0', '1'}, {'1', '0'},
		}),
	)
}
