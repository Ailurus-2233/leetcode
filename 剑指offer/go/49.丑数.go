package Offer49

import "container/heap"

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

func nthUglyNumber(n int) int {
	hp := &IntHeap{
		heap: []int{1},
		flag: true,
	}
	heap.Init(hp)
	ans := 1
	for i := 0; i <= n; i++ {
		heap.Push(hp, hp.Top()*2)
		heap.Push(hp, hp.Top()*3)
		heap.Push(hp, hp.Top()*5)
		ans = heap.Pop(hp).(int)
	}
	return ans
}
