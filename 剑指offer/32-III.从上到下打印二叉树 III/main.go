/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main

import "fmt"

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
	queue = append(queue, root)
	level := -1
	for len(queue) != 0 {
		length := len(queue)
		level *= -1

		for i := 0; i < length; i++ {
			fmt.Printf("%d ", queue[i].Val)
		}
		fmt.Println()

		tmp := []int{}
		if level == -1 {
			for i := 0; i < length; i++ {
				tmp = append(tmp, queue[i].Val)
			}
		} else {
			for i := length - 1; i >= 0; i-- {
				tmp = append(tmp, queue[i].Val)
			}
		}

		for i := 0; i < length; i++ {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, tmp)
		queue = queue[length:]
	}
	return ans
}
