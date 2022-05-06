/*
 * @lc app=leetcode.cn id=713 lang=golang
 *
 * [713] 乘积小于 K 的子数组
 */

package main

import (
	"fmt"
)

// @lc code=start
func numSubarrayProductLessThanK(nums []int, k int) int {
	prod, left, ans := 1, 0, 0
	for right := 0; right < len(nums); right++ {
		prod *= nums[right]
		for ; left <= right && prod >= k; left++ {
			prod /= nums[left]
		}
		ans += right - left + 1
	}
	return ans
}

// @lc code=end

func main() {
	nums := []int{10, 5, 2, 6}
	k := 100
	fmt.Println(numSubarrayProductLessThanK(nums, k))
}
