package main

func countDigitOne(n int) int {
	// 315 1, 11, 21, 31 ... 31次 31 * k + k 32
	// 		10， 11， 12... 19 10次 3 * 10 30次 5+1次 36
	// 		100 次 * 2 + 15 + 1 1 * 100 100

	ans := 0
	a, b, k := n%10, 0, 1
	for n > 0 {
		n /= 10
		if a > 1 {
			ans += n*k + k
		} else if a < 1 {
			ans += n * k
		} else {
			ans += n*k + b + 1
		}
		b += a * k
		k *= 10
		a = n % 10
	}
	return ans
}

func main() {
	println(countDigitOne(315))
}
