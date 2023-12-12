/*
 * @lc app=leetcode.cn id=398 lang=golang
 *
 * [398] 随机数索引
 */

package main

import "math/rand"

// @lc code=start
type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (this *Solution) Pick(target int) int {
	count := 0
	ans := -1
	for i, v := range this.nums {
		if v == target {
			count++
			if rand.Intn(count) == 0 {
				ans = i
			}
		}
	}
	return ans
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
// @lc code=end

func main() {}
