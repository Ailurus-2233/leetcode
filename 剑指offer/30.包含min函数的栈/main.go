package main

import "fmt"

type MinStack struct {
	minStack []int
	stack    []int
	top      int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		minStack: make([]int, 0),
		stack:    make([]int, 0),
		top:      0,
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if this.top == 0 {
		this.minStack = append(this.minStack, x)
	} else {
		if x <= this.minStack[len(this.minStack)-1] {
			this.minStack = append(this.minStack, x)
		}
	}
	this.top++
}

func (this *MinStack) Pop() {
	this.top--
	tmp := this.stack[this.top]
	if tmp == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
	this.stack = this.stack[:this.top]
}

func (this *MinStack) Top() int {
	return this.stack[this.top-1]
}

func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

func main() {
	minStack := Constructor()
	minStack.Push(0)
	minStack.Push(1)
	minStack.Push(0)
	fmt.Println(minStack.Top())
	minStack.Pop()
	fmt.Println(minStack.Min())
}
