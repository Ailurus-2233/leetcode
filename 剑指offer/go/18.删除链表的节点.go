package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	p := head
	for head.Next != nil {
		if head.Next.Val == val {
			head.Next = head.Next.Next
			break
		}
		head = head.Next
	}
	return p
}

func main() {

}
