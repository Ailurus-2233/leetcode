/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	ans := []int{}
	ans = append(ans, root.Val)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
			ans = append(ans, node.Left.Val)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			ans = append(ans, node.Right.Val)
		}
	}
	return ans
}
