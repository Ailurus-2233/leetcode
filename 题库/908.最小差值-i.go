/*
 * @lc app=leetcode.cn id=908 lang=golang
 *
 * [908] 最小差值 I
 */
package main

import "fmt"

// @lc code=start
func smallestRangeI(nums []int, k int) int {
	if len(nums) < 2 {
		return 0
	}
	tmp := nums[1]
	for i, v := range nums {
		if i == 1 {
			if nums[0] < tmp {
				nums[0] = tmp
			}
			if nums[1] > tmp {
				nums[1] = tmp
			}
		} else {
			if nums[0] < v {
				nums[0] = v
			}
			if nums[1] > v {
				nums[1] = v
			}
		}
	}

	if (nums[0]-nums[1])-2*k > 0 {
		return (nums[0] - nums[1]) - 2*k
	} else {
		return 0
	}
}

// @lc code=end

func main() {
	fmt.Println(smallestRangeI([]int{0, 10}, 2))
}
