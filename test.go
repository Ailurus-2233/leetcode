package main

import (
	"container/heap"
	"fmt"
)

type IntHeap struct {
	heap []int
	flag bool //false 大顶堆 true 小顶堆
}

func (h IntHeap) Len() int {
	return len(h.heap)
}
func (h IntHeap) Less(i, j int) bool {
	if h.flag {
		return h.heap[i] < h.heap[j]
	} else {
		return h.heap[i] > h.heap[j]
	}
}

func (h *IntHeap) Swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *IntHeap) Push(x interface{}) {
	h.heap = append(h.heap, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	n := h.Len()
	x := h.heap[n-1]
	h.heap = h.heap[:n-1]
	return x
}

func (h IntHeap) Top() int {
	return h.heap[0]
}

func main() {
	h := &IntHeap{
		heap: []int{2, 1, 5, 100, 3, 6, 4, 5},
		flag: true,
	}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", h.heap[0])
	fmt.Println(h.heap)
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
		fmt.Println(h.heap)
	}

}
