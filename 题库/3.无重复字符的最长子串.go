/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
package main

import "fmt"

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	ans := 0
	start := 0
	for end := 0; end < len(s); end++ {
		if v, ok := m[s[end]]; ok && v >= start {
			start = max(start, m[s[end]]+1)

		} else {
			ans = max(ans, end-start+1)
		}
		m[s[end]] = end
		// fmt.Println(m, start, end)
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
func main() {
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}
