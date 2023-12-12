/*
 * @lc app=leetcode.cn id=824 lang=golang
 *
 * [824] 山羊拉丁文
 */

package main

import (
	"fmt"
	"strings"
)

// @lc code=start
func toGoatLatin(sentence string) string {
	build := strings.Builder{}
	count := 1
	first := true
	unvowel := ' '
	for _, char := range sentence {
		if first {
			if isVowel(byte(char)) {
				build.WriteByte(byte(char))
			} else {
				unvowel = char
			}
			first = false
		} else {
			if char == ' ' {
				first = true
				if unvowel != ' ' {
					build.WriteRune(unvowel)
				}
				build.WriteString("ma")
				for i := 0; i < count; i++ {
					build.WriteRune('a')
				}
				build.WriteRune(' ')
				count++
				unvowel = ' '
			} else {
				build.WriteRune(char)
			}
		}
	}
	if unvowel != ' ' {
		build.WriteRune(unvowel)
	}
	build.WriteString("ma")
	for i := 0; i < count; i++ {
		build.WriteRune('a')
	}
	return build.String()
}

func isVowel(char byte) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U'
}

// @lc code=end

func main() {
	fmt.Println(toGoatLatin("The quick brown fox jumped over the lazy dog") == "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa")
}
