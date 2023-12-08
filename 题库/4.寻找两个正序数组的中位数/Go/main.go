/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */
package main

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	ans1, ans2 := 0, 0
	index1, index2 := 0, 0
	for i := 0; i <= length/2; i++ {
		if index1 < len(nums1) && index2 < len(nums2) {
			if nums1[index1] < nums2[index2] {
				index1++
				ans2 = ans1
				ans1 = nums1[index1-1]
			} else {
				index2++
				ans2 = ans1
				ans1 = nums2[index2-1]
			}
		} else if index1 < len(nums1) {
			index1++
			ans2 = ans1
			ans1 = nums1[index1-1]
		} else {
			index2++
			ans2 = ans1
			ans1 = nums2[index2-1]
		}
	}
	if length%2 == 0 {
		return (float64(ans1) + float64(ans2)) / 2
	} else {
		return float64(ans1)
	}
}

// @lc code=end
