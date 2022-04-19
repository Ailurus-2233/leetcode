package main

import "fmt"

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	ans := 1
	for n > 4 {
		ans *= 3
		n -= 3
		ans %= 1000000007
	}
	return ans * n % 1000000007
}

func main() {
	fmt.Println(cuttingRope(2))
}
