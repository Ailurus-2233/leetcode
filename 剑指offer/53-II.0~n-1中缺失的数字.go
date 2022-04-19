package main

import "fmt"

// 时间复杂度为O(n)方法
// func missingNumber(nums []int) int {
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] != i {
// 			return i
// 		}
// 	}
// 	return -1
// }

// 时间复杂度为O(log2n)
// func missingNumber(nums []int) int {
// 	left := 0
// 	right := len(nums)
// 	middle := (left + right) / 2

// 	for left < right {
// 		if nums[middle] == middle {
// 			left = middle + 1
// 		} else {
// 			right = middle
// 		}
// 		middle = (left + right) / 2
// 	}

// 	return left
// }

// O(log2n) 递归写法
func missingNumber(nums []int) int {
	return innerFunc(nums, 0, len(nums))
}

func innerFunc(nums []int, left int, right int) int {
	if left == right {
		return left
	}

	middle := (left + right) / 2
	if nums[middle] == middle {
		return innerFunc(nums, middle+1, right)
	} else {
		return innerFunc(nums, left, middle)
	}
}

func main() {
	var a = []int{0, 1, 2, 3, 4, 6, 7, 8, 9}
	fmt.Println(missingNumber(a))
}
