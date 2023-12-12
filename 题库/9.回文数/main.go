/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */
package main

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	y := 0
	for x > y {
		y = y*10 + x%10
		x /= 10
	}
	return x == y || x == y/10
}

// @lc code=end
