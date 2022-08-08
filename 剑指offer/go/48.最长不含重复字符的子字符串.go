package o48

func lengthOfLongestSubstring(s string) int {
	flag := make(map[rune]int)
	maxLen := 0
	curLen := 0
	start := 0
	for i, v := range s {
		if _, ok := flag[v]; ok {
			maxLen = max(maxLen, curLen)
			for j := start; j < i; j++ {
				if rune(s[j]) == v {
					start = j + 1
					curLen = i - start
					break
				}
			}
		}
		flag[v] = i
		curLen++
	}
	return max(maxLen, curLen)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
