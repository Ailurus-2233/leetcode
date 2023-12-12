/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

package main

import "fmt"

// @lc code=start
func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)
	dp := make([][]bool, m+1)

	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
		for j := 0; j < n+1; j++ {
			if j == 0 {
				dp[i][j] = i == 0
			} else {
				if p[j-1] != '*' {
					if i > 0 {
						if p[j-1] == s[i-1] || p[j-1] == '.' {
							dp[i][j] = dp[i][j] || dp[i-1][j-1]
						}
					}
				} else {
					if i > 0 {
						if p[j-2] == s[i-1] || p[j-2] == '.' {
							dp[i][j] = dp[i][j] || dp[i-1][j]
						}
						dp[i][j] = dp[i][j] || dp[i][j-2]
					} else {
						dp[i][j] = dp[i][j] || dp[i][j-2]
					}
				}
			}
		}
	}
	return dp[m][n]
}

// @lc code=end

func main() {
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
}
