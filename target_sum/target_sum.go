package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum < target {
		return 0
	}

	if target < (-1 * sum) {
		return 0
	}

	if (sum+target)%2 != 0 {
		return 0
	}

	return subsetNum(nums, (target+sum)>>1)
}

func subsetNum(nums []int, sum int) int {
	dp := make([]int, sum+1)
	dp[0] = 1
	for _, n := range nums {
		for i := sum; i >= n; i-- {
			dp[i] = dp[i] + dp[i-n]
		}
	}

	return dp[sum]
}

func main() {
	fmt.Println(findTargetSumWays([]int{100}, -200))
}
