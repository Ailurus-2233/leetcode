package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n == -1 {
		return 1.0 / x
	}
	double := myPow(x, n/2)
	half := myPow(x, n%2)
	return double * double * half
}

func main() {
	fmt.Println(myPow(2, 10))
}
