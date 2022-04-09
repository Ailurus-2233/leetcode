package main

import "fmt"

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 3, 5}
	fmt.Print(findRepeatNumber(arr))
}
