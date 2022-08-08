package main

import (
	"fmt"
)

func printNumbers(n int) []int {
	l := 1
	for i := 0; i < n; i++ {
		l *= 10
	}
	ans := make([]int, l-1)
	for i := 0; i < len(ans); i++ {
		ans[i] = i + 1
	}
	return ans
}

func main() {
	fmt.Print(printNumbers(3))
}
