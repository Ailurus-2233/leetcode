package main

// func maxProfit(prices []int) int {
// 	/*
// 	 * 前日最大利润和今日利润比较
// 	 */

// 	if prices == nil || len(prices) == 0 {
// 		return 0
// 	}
// 	dp := make([]int, len(prices))
// 	dp[0] = 0

// 	for i := 1; i < len(prices); i++ {
// 		dp[i] = max(dp[i-1], prices[i]-min(prices[0:i]))
// 	}

// 	return dp[len(prices)-1]
// }

// func max(a int, b int) int {
// 	if a < b {
// 		return b
// 	}
// 	return a
// }

// func min(arr []int) int {
// 	ans := arr[0]
// 	for _, v := range arr {
// 		if v < ans {
// 			ans = v
// 		}
// 	}
// 	return ans
// }

//优化
func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}

	dp := 0
	min := prices[0]

	for i := 1; i < len(prices); i++ {
		dp = max(dp, prices[i]-min)
		if prices[i] < min {
			min = prices[i]
		}
	}
	return dp
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
