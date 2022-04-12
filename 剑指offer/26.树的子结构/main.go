package main

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	if A == nil {
		return false
	}
	return isSame(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSame(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil && B != nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return isSame(A.Left, B.Left) && isSame(A.Right, B.Right)
}
