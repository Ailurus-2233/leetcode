package main

import (
	"fmt"
)

func findNthDigit(n int) int {
	i, k, a, b := 2, 10, 0, 9
	for b < n {
		a, b = b, i*k*9+b
		i++
		k *= 10
	}

	k /= 10
	i -= 1

	number := (n-a)/i + k
	if (n-a)%i == 0 {
		number -= 1
	}
	digit := (n - a) % i
	if digit == 0 {
		digit = i - 1
	} else {
		digit--
	}
	for j := 0; j < digit; j++ {
		number -= number / k * k
		k /= 10
	}
	if k == 0 {
		return number
	}
	return number / k
}
func main() {
	fmt.Println(findNthDigit(2147483647))
}
