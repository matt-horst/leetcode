package p2435

import "fmt"

const MOD int = 1e9+7

func solve(i, j, rem int, dp [][]map[int]int, grid [][]int, k int) int {
	if i == 0 && j == 0 {
		if (grid[0][0] + rem) % k == 0 {
			return 1
		} else {
			return 0
		}
	}

	if ans, ok := dp[i][j][rem]; ok {
		return ans
	}

	total := 0
	if i > 0 {
		total += solve(i - 1, j, (grid[i][j] + rem) % k, dp, grid, k)
		total %= MOD
	}

	if j > 0 {
		total += solve(i, j - 1, (grid[i][j] + rem) % k, dp, grid, k)
		total %= MOD
	}

	dp[i][j][rem] = total

	return total
}

func numberOfPaths(grid [][]int, k int) int {
	n := len(grid)
	m := len(grid[0])


	dp := make([][]map[int]int, n)
	for i := range n {
		dp[i] = make([]map[int]int, m)
		for j := range m {
			dp[i][j] = make(map[int]int, k)
		}
	}

	solve(n-1, m-1, 0, dp, grid, k)

	fmt.Printf("dp: %v\n", dp)

	return dp[n-1][m-1][0]
}
