package main

import "fmt"

type Stack struct {
	data []int
	top  int
}

func (this *Stack) Push(value int) {
	this.data = append(this.data, value)
	this.top++
}

func (this *Stack) Pop() int {
	pop := this.data[this.top-1]
	if this.top == 0 {
		return -1
	}
	this.top--
	this.data = this.data[:this.top]
	return pop
}

func (this *Stack) Length() int {
	return this.top
}

func InitStack() *Stack {
	return &Stack{
		data: make([]int, 0),
		top:  0,
	}
}

type CQueue struct {
	stack1 Stack
	stack2 Stack
}

func Constructor() CQueue {
	return CQueue{
		stack1: *InitStack(),
		stack2: *InitStack(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.stack2.Length() != 0 {
		return this.stack2.Pop()
	} else {
		if this.stack1.Length() == 0 {
			return -1
		}
		for this.stack1.Length() != 0 {
			this.stack2.Push(this.stack1.Pop())
		}
		return this.stack2.Pop()
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
func main() {
	obj := Constructor()

	obj.AppendTail(3)
	fmt.Printf("%d\n", obj.DeleteHead())
	fmt.Printf("%d\n", obj.DeleteHead())
}
