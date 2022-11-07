/*
 * @lc app=leetcode.cn id=901 lang=golang
 *
 * [901] 股票价格跨度
 */
package main

// @lc code=start
type StockSpanner struct {
	stack [][2]int
	idx   int
}

func Constructor() StockSpanner {
	return StockSpanner{[][2]int{}, -1}
}

func (s *StockSpanner) Next(price int) int {
	s.idx++
	if s.idx == 0 {
		s.stack = append(s.stack, [2]int{0, price})
		return 1
	}
	for len(s.stack) > 0 && price >= s.stack[len(s.stack)-1][1] {
		s.stack = s.stack[:len(s.stack)-1]
	}
	s.stack = append(s.stack, [2]int{s.idx, price})
	if len(s.stack) == 1 {
		return s.idx + 1
	}
	return s.idx - s.stack[len(s.stack)-2][0]
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
// @lc code=end
