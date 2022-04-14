/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 * Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
 * Output: [3,9,20,null,null,15,7]
 * [1,2] [1,2]
 *
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}
	inIndex := -1
	for i, v := range inorder {
		if v == preorder[0] {
			inIndex = i
			break
		}
	}

	root.Left = buildTree(preorder[1:1+inIndex], inorder[0:inIndex])
	root.Right = buildTree(preorder[1+inIndex:], inorder[inIndex+1:])

	return root
}
