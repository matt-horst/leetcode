package p1262

import "math"

func maxSumDivThree(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, 3)
		dp[i] = []int{0, math.MinInt, math.MinInt}
	}

	dp[n-1][nums[n-1]%3] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		num := nums[i]
		mod := num % 3
		for j := range 3 {
			sum := num + dp[i+1][(3 - mod + j) % 3]
			dp[i][j] = max(sum, dp[i+1][j])
		}
	}

	return max(dp[0][0])
}
