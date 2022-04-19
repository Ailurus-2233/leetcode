package main

import "fmt"

//数论解法
//1. 任何大于1的数都可以由2和3相加获得
//2. 因为2*2 = 1*4,2*3>1*5, 所以将数拆分成最细力度的数可得乘积最大
//3. 因为2*2*2 < 3 * 3, 所以拆分结果中3的数量越多，乘积越大
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	ans := 1
	for n > 4 {
		ans *= 3
		n -= 3
	}
	return ans * n
}

//动态规划解法
func cuttingRopeDP(n int) int {
	if n <= 3 {
		return n - 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max2(dp[i], max2(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func max2(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Print(cuttingRopeDP(6))
}
