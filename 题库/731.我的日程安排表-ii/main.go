package main

/*
 * @lc app=leetcode.cn id=731 lang=golang
 *
 * [731] 我的日程安排表 II
 */

// @lc code=start

type Pair struct {
	start, end int
}

type MyCalendarTwo map[int]Pair

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (this MyCalendarTwo) update(start, end, add, l, r, idx int) {
	//边界条件
	if r < start || end < l {
		return
	}
	if start <= l && r <= end {
		p := this[idx]
		p.start += add
		p.end += add
		this[idx] = p
		return
	}
	mid := (l + r) / 2
	this.update(start, end, add, l, mid, 2*idx)
	this.update(start, end, add, mid+1, r, 2*idx+1)
	p := this[idx]
	p.start = p.end + max(this[2*idx].start, this[2*idx+1].start)
	this[idx] = p
}

func (this MyCalendarTwo) Book(start int, end int) bool {
	this.update(start, end-1, 1, 0, 10e9, 1)
	if this[1].start > 2 {
		this.update(start, end-1, -1, 0, 10e9, 1)
		return false
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
// @lc code=end
