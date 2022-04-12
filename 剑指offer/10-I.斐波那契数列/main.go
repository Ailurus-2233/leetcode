package main

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	a := 0
	b := 1

	for i := 1; i <= n; i++ {
		b = a + b
		a = b - a
		a %= 1e9 + 7
	}
	return a

}
