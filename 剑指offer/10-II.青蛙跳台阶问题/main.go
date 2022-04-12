package main

func numWays(n int) int {

	if n == 0 || n == 1 {
		return 1
	}

	/*
	 * 思想描述：
	 * n=0，不动，方法有1种 num_0 = 1
	 * n=1，跳一个台阶，方法有1种 num_1 = 1
	 * n=2，从台阶1跳1，从台阶0跳2，num_2 = num_0 + num_1
	 * n=3, 从台阶2跳1，从台阶1跳2，num_3 = num_1 + num_2
	 * ...
	 * n = k, num_k = num_k-2 + num_k-1
	 */

	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i < n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}
	return dp[n]
}
