package main

func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	root := postorder[len(postorder)-1]
	i := 0
	for i < len(postorder)-1 {
		if postorder[i] > root {
			i++
		} else {
			return false
		}
	}
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 0; i < len(postorder)-1; i++ {
		if postorder[i] > root {
			right = append(right, postorder[i])
		} else {
			left = append(left, postorder[i])
		}
	}
	return verifyPostorder(left) && verifyPostorder(right)
}
