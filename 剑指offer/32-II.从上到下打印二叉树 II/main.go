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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	queue := []*TreeNode{}
	level := []int{}
	queue = append(queue, root)
	ans = append(ans, []int{root.Val})
	level = append(level, 0)
	for len(queue) != 0 {
		node := queue[0]
		nodeLevel := level[0] + 1
		queue = queue[1:]
		level = level[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
			level = append(level, nodeLevel)
			if nodeLevel == len(ans) {
				ans = append(ans, []int{node.Left.Val})
			} else {
				ans[nodeLevel] = append(ans[nodeLevel], node.Left.Val)
			}
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			level = append(level, nodeLevel)
			if nodeLevel == len(ans) {
				ans = append(ans, []int{node.Right.Val})
			} else {
				ans[nodeLevel] = append(ans[nodeLevel], node.Right.Val)
			}
		}
	}

	return ans
}
