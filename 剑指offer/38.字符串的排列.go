package main

func permutation(s string) []string {
	return exec(s)
}

func exec(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}
	ans := make([]string, 0)
	m := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; !ok {
			tmp := exec(s[:i] + s[i+1:])
			m[s[i]] = true
			for j := 0; j < len(tmp); j++ {
				ans = append(ans, string(s[i])+tmp[j])
			}
		}
	}
	return ans
}
