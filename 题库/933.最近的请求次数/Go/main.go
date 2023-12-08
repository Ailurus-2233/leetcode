/*
 * @lc app=leetcode.cn id=933 lang=golang
 *
 * [933] 最近的请求次数
 */
package main

// @lc code=start
type RecentCounter struct {
	requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	this.requests = append(this.requests, t)
	for len(this.requests) > 0 && this.requests[0] < t-3000 {
		this.requests = this.requests[1:]
	}
	return len(this.requests)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
// @lc code=end
