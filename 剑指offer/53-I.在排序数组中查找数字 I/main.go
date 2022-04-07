package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	index := getIndex(nums, target)
	if index == -1 {
		return 0
	} else {
		count := 0
		for i := index; i < len(nums) && nums[i] == target; i++ {
			count++
		}
		for i := index - 1; i >= 0 && nums[i] == target; i-- {
			count++
		}
		return count
	}
}

func getIndex(nums []int, target int) int {
	if len(nums) == 1 && nums[0] != target {
		return -1
	}
	if len(nums) <= 0 {
		return -1
	}
	middle := int(len(nums) / 2)
	if nums[middle] == target {
		return middle
	}
	if nums[middle] > target {
		return getIndex(nums[:middle], target)
	} else {
		return getIndex(nums[middle+1:], target) + middle + 1
	}
}
