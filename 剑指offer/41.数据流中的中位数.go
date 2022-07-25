package main

import "container/heap"

type IntHeap struct {
	heap []int
	flag bool
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

func (h IntHeap) Swap(i, j int) {
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

type MedianFinder struct {
	smallHeap IntHeap //记录较大的一半
	bigHeap   IntHeap //记录较小的一半
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	smallHeap := &IntHeap{
		heap: []int{},
		flag: true,
	}
	bigHeap := &IntHeap{
		heap: []int{},
		flag: false,
	}
	heap.Init(smallHeap)
	heap.Init(bigHeap)
	return MedianFinder{
		smallHeap: *smallHeap,
		bigHeap:   *bigHeap,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.smallHeap.Len() == 0 {
		heap.Push(&this.smallHeap, num)
	} else {
		if this.FindMedian() < float64(num) {
			heap.Push(&this.smallHeap, num)
			if this.smallHeap.Len() > this.bigHeap.Len()+1 {
				heap.Push(&this.bigHeap, heap.Pop(&this.smallHeap))
			}
		} else {
			heap.Push(&this.bigHeap, num)
			if this.bigHeap.Len() > this.smallHeap.Len() {
				heap.Push(&this.smallHeap, heap.Pop(&this.bigHeap))
			}
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.smallHeap.Len() == this.bigHeap.Len() {
		return float64(this.smallHeap.Top()+this.bigHeap.Top()) / 2.0
	} else {
		return float64(this.smallHeap.Top())
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
