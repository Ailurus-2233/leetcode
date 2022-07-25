package main

func maxSubArray(nums []int) int {
	// // dp[i]表示以nums[i]结尾的最大子数组的和
	// dp := make([]int, len(nums))
	// dp[0] = nums[0]
	// for i := 1; i < len(nums); i++ {
	// 	dp[i] = max(dp[i-1]+nums[i], nums[i])
	// }
	// m := dp[0]
	// for i := 0; i < len(nums); i++ {
	// 	m = max(m, dp[i])
	// }
	// return m
	// 用 tmp 来表示 dp[i] m 记录最大值
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i-1]+nums[i], nums[i])
		m = max(m, nums[i])
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
