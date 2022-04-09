package main

import (
	"fmt"
)

func minArray(numbers []int) int {
	return numbers[binSearch(numbers, 0, len(numbers)-1)]
}

func binSearch(numbers []int, left int, right int) int {

	middle := (left + right) / 2
	if left == right {
		return middle
	}
	if numbers[middle] > numbers[right] {
		return binSearch(numbers, middle+1, right)
	} else if numbers[middle] < numbers[right] {
		return binSearch(numbers, left, middle)
	} else {
		return binSearch(numbers, left, right-1)
	}
}

func main() {
	arr := []int{3, 4, 5, 1, 2}
	fmt.Println(minArray(arr))
}
