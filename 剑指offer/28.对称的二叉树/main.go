package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSame(root.Left, root.Right)
}

func isSame(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil && right != nil {
		return left.Val == right.Val && isSame(left.Left, right.Right) && isSame(left.Right, right.Left)
	}
	return false
}
