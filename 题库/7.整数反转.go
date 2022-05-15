/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

package main

// @lc code=start
func reverse(x int) int {
	ans := 0
	for x != 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	if ans > 2147483647 || ans < -2147483648 {
		return 0
	}
	return ans
}

// @lc code=end
