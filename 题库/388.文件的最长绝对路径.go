/*
 * @lc app=leetcode.cn id=388 lang=golang
 *
 * [388] 文件的最长绝对路径
 */

package main

import "fmt"

// @lc code=start
func lengthLongestPath(input string) int {
	stack := make([]int, 0)
	level := 0
	maxLength := 0
	start := 0
	fileFlag := false
	length := 0
	for i, c := range input {
		if c == '\t' {
			level++
			start = i + 1
			fileFlag = false
		}
		if c == '.' {
			fileFlag = true
			continue
		}
		if c == '\n' {
			if fileFlag {
				if level == 0 {
					length = i - start
				} else {
					length = stack[level-1] - start + i
				}
				if length > maxLength {
					maxLength = length
				}
			} else {
				if level == len(stack) {
					if level == 0 {
						stack = append(stack, i-start+1)
					} else {
						stack = append(stack, i-start+stack[level-1]+1)
					}
				} else {
					if level == 0 {
						stack[level] = i - start + 1
					} else {
						stack[level] = i - start + stack[level-1] + 1

					}
				}
			}
			level = 0
			start = i + 1
			fileFlag = false
		}
		if i == len(input)-1 {
			if fileFlag {
				if level == 0 {
					length = i - start + 1
				} else {
					length = stack[level-1] - start + i + 1
				}
				if length > maxLength {
					maxLength = length
				}
			}
		}
	}

	return maxLength
}

// @lc code=end

func main() {
	fmt.Println(lengthLongestPath("a\n\tb.txt\na2\n\tb2.txt"))
	//a
	//->b.txt
	//a2
	//->b2.txt
}
