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

/*
 * 递归思路解决问题
 */
// func reversePrint(head *ListNode) []int {
// 	if head == nil {
// 		return []int{}
// 	}
// 	return append(reversePrint(head.Next), head.Val)
// }

/*
 * 非递归思路
 */
func reversePrint(head *ListNode) []int {
	tmp := []int{}
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}
	for i := 0; i < len(tmp)/2; i++ {
		tmp[i], tmp[len(tmp)-1-i] = tmp[len(tmp)-1-i], tmp[i]
	}
	return tmp
}

func main() {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}
	head.add(2)
	head.add(3)

	fmt.Print(reversePrint(head))

}
