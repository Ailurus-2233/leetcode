/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	var crl map[*Node]*Node = map[*Node]*Node{}
	var ans = &Node{
		Val:    head.Val,
		Next:   nil,
		Random: nil,
	}
	crl[head] = ans
	var tmp = ans
	var p = head.Next
	for p != nil {
		tmp.Next = &Node{
			Val:    p.Val,
			Next:   nil,
			Random: nil,
		}
		crl[p] = tmp
		tmp = tmp.Next
		p = p.Next
	}

	p = head
	tmp = ans
	for p != nil {
		if p.Random != nil {
			tmp.Random = crl[p.Random]
		}
		tmp = tmp.Next
		p = p.Next
	}

	return ans
}
