/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */
package main

import "math"

// @lc code=start
// func findDuplicates(nums []int) []int {
// 	ans := make([]int, 0, 50)

// 	for i := 0; i < len(nums); i++ {
// 		nums[i]--
// 	}

// 	for i := 0; i < len(nums); i++ {
// 		for nums[i] != i {
// 			if nums[i] == nums[nums[i]] {
// 				if !IsContain(ans, nums[i]+1) {
// 					ans = append(ans, nums[i]+1)
// 				}
// 				break
// 			}
// 			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
// 		}
// 	}
// 	return ans
// }

// func IsContain(items []int, item int) bool {
// 	for _, eachItem := range items {
// 		if eachItem == item {
// 			return true
// 		}
// 	}
// 	return false
// }
func findDuplicates(nums []int) []int {
	ans := make([]int, 0, 50)
	for i := 0; i < len(nums); i++ {
		num := int(math.Abs(float64(nums[i])))
		if nums[num-1] > 0 {
			nums[num-1] *= -1
		} else {
			ans = append(ans, num)
		}
	}
	return ans
}

// @lc code=end
