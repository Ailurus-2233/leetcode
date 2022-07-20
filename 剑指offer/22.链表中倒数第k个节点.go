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

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	for i := 0; i < k; i++ {
		p = p.Next
	}
	q := head
	for p != nil {
		p = p.Next
		q = q.Next
	}
	return q
}
