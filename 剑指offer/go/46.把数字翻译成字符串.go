package main

func translateNum(num int) int {
	// dp := make([]int, 0)
	// dp = append(dp, 1)
	dp1, dp2 := 1, 1
	n := num % 10
	num /= 10
	for ; num > 0; num /= 10 {
		if num%10*10+n > 25 || num%10 == 0 {
			dp2 = dp1
		} else {
			dp1, dp2 = dp1+dp2, dp1
		}
		n = num % 10
	}
	return dp1
}
