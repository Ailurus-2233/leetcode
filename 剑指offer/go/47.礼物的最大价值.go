package o47

func maxValue(grid [][]int) int {
	// dp := make([][]int, len(grid))
	// for i := range dp {
	// 	dp[i] = make([]int, len(grid[0]))
	// }
	// dp[0][0] = grid[0][0]
	// for i := 1; i < len(grid); i++ {
	// 	dp[i][0] = dp[i-1][0] + grid[i][0]
	// }
	// for i := 0; i < len(grid); i++ {
	// 	for j := 1; j < len(grid[0]); j++ {
	// 		if i == 0 {
	// 			dp[i][j] = dp[i][j-1] + grid[i][j]
	// 		} else {
	// 			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
	// 		}
	// 	}
	// }
	// return dp[len(grid)-1][len(grid[0])-1]
	dp := make([]int, len(grid[0]))
	dp[0] = grid[0][0]
	for i := 1; i < len(grid[0]); i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}
	for i := 1; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if j == 0 {
				dp[j] = dp[j] + grid[i][j]
			} else {
				dp[j] = max(dp[j], dp[j-1]) + grid[i][j]
			}
		}
	}
	return dp[len(grid[0])-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
