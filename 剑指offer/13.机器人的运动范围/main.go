package main

func movingCount(m int, n int, k int) int {
	count := 0
	divide := k + 1
	if k >= 8 {
		divide = (k - 7) * 10
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n && i+j < divide; j++ {
			if isValid(i, j, k) {
				count++
			}
		}
	}
	return count
}

func isValid(x int, y int, k int) bool {
	sum := 0
	for x > 0 {
		sum += x % 10
		x /= 10
	}
	for y > 0 {
		sum += y % 10
		y /= 10
	}
	return sum <= k
}
