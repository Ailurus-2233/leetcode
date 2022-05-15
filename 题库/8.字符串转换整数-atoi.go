/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */
package main

import (
	"fmt"
	"math"
)

// @lc code=start

var digits = map[byte]int{
	0x30: 0,
	0x31: 1,
	0x32: 2,
	0x33: 3,
	0x34: 4,
	0x35: 5,
	0x36: 6,
	0x37: 7,
	0x38: 8,
	0x39: 9,
}

func myAtoi(s string) int {
	res, sign := 0, 1

	space := true
	signF := true
	for i := 0; i < len(s); i++ {
		if space && s[i] == ' ' {
			continue
		} else {
			space = false
			if signF && s[i] == '-' {
				sign = -1
				signF = false
			} else if signF && s[i] == '+' {
				sign = 1
				signF = false
			} else if s[i] >= 0x30 && s[i] <= 0x39 {
				signF = false
				res = res*10 + digits[s[i]]
				if sign*res > math.MaxInt32 {
					return math.MaxInt32
				}

				if sign*res < math.MinInt32 {
					return math.MinInt32
				}
			} else {
				break
			}
		}
	}

	return res * sign
}

// @lc code=end

func main() {
	fmt.Println(myAtoi("+-12"))
}
