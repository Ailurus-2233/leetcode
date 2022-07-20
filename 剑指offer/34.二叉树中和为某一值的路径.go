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

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	path := []int{}
	dfs(root, target, path, &res)
	return res
}

func dfs(root *TreeNode, target int, path []int, res *[][]int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil && target == root.Val {
		*res = append(*res, path)
	}
	dfs(root.Left, target-root.Val, path, res)
	dfs(root.Right, target-root.Val, path, res)
	path = path[:len(path)-1]
}
