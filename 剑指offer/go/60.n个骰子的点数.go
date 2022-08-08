package main

import "fmt"

func dicesProbability(n int) []float64 {
	nums := []float64{1.0 / 6.0, 1.0 / 6.0, 1.0 / 6.0, 1.0 / 6.0, 1.0 / 6.0, 1.0 / 6.0}
	for i := 1; i < n; i++ {
		tmp := make([]float64, 6*(i+1))
		for j := i - 1; j < len(nums); j++ {
			for k := 1; k <= 6; k++ {
				tmp[j+k] += nums[j] * (1.0 / 6.0)
			}
		}
		nums = tmp
	}
	return nums[n-1:]
}

func main() {
	fmt.Println(dicesProbability(3))
}
