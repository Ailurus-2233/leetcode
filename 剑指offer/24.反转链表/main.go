/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (this *ListNode) add(x int) {
	var p = this
	for p.Next != nil {
		p = p.Next
	}
	p.Next = &ListNode{
		Val:  x,
		Next: nil,
	}
}

// 递归思路解决问题
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	p := reverseList(head.Next)
// 	head.Next.Next = head
// 	head.Next = nil
// 	return p
// }

// 非递归思路解决问题
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	return pre
}

func main() {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}

	head.add(2)
	head.add(3)
	head.add(4)

	head = reverseList(head)

	for head != nil {
		fmt.Print(head.Val, "->")
		head = head.Next
	}
}
