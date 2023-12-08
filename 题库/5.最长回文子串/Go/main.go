/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
package main

import "fmt"

// @lc code=start
func longestPalindrome(s string) string {
	ans := ""
	for i := 0; i < len(s); i++ {
		tmp := getPalindrome(s, i, i)
		if len(tmp) > len(ans) {
			ans = tmp
		}
		tmp = getPalindrome(s, i, i+1)
		if len(tmp) > len(ans) {
			ans = tmp
		}
	}
	return ans
}

func getPalindrome(s string, left int, rigth int) string {
	for left >= 0 && rigth < len(s) && s[left] == s[rigth] {
		left--
		rigth++
	}
	return s[left+1 : rigth]
}

// @lc code=end

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}
