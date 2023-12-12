/**
 * 821.字符的最短距离
 * 2022-04-19 14:50:13
 * author: Ailurus
**/
package main

import "fmt"

func shortestToChar(s string, c byte) []int {
	ans := make([]int, len(s))
	for i, v := range s {
		if byte(v) == c {
			ans[i] = 0
			continue
		}
		for j := 1; ; j++ {
			if i-j >= 0 && i+j < len(s) {
				if s[i-j] == c || s[i+j] == c {
					ans[i] = j
					break
				}
			}
			if i-j < 0 && s[i+j] == c {
				ans[i] = j
				break
			}
			if i+j >= len(s) && s[i-j] == c {
				ans[i] = j
				break
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
}
