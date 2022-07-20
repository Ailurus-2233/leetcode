package main

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	}
	stack := make([]int, 0)
	poppedIndex := 0
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[poppedIndex] {
			stack = stack[:len(stack)-1]
			poppedIndex += 1
		}
	}
	return len(stack) == 0
}
