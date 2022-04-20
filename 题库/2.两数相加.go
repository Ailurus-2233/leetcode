/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	ans := l1
	for l1 != nil {
		if l2 != nil {
			l1.Val += l2.Val + carry
			l2 = l2.Next
		} else {
			l1.Val += carry
		}
		carry = l1.Val / 10
		l1.Val %= 10
		if l1.Next == nil {
			if l2 != nil {
				l1.Next = l2
				l2 = nil
			} else if carry != 0 {
				l1.Next = &ListNode{Val: carry, Next: nil}
				carry = 0
			}
		}
		l1 = l1.Next
	}
	return ans
}

// @lc code=end
