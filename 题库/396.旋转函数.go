/*
 * @lc app=leetcode.cn id=396 lang=golang
 *
 * [396] 旋转函数
 */

package main

// @lc code=start
func maxRotateFunction(nums []int) int {
	f := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		f += i * nums[i]
		sum += nums[i]
	}
	max := f
	for i := len(nums) - 1; i >= 0; i-- {
		f = f - len(nums)*nums[i] + sum
		if f > max {
			max = f
		}
	}
	return max
}

// @lc code=end

func main() {}
