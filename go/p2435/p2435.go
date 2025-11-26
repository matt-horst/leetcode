package p2435

const MOD int = 1e9+7

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

	dp[0][0][grid[0][0] % k] = 1
	for i := range n {
		for j := range m {
			if i > 0 {
				for key, val := range dp[i-1][j] {
					rem := (key + grid[i][j]) % k
					dp[i][j][rem] += val
					dp[i][j][rem] %= MOD
				}
			}

			if j > 0 {
				for key, val := range dp[i][j-1] {
					rem := (key + grid[i][j]) % k
					dp[i][j][rem] += val
					dp[i][j][rem] %= MOD
				}
			}
		}
	}

	return dp[n-1][m-1][0]
}
