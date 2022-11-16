/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
package main

// @lc code=start
func maxProfit(prices []int) int {
	res, min := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if tmp := prices[i] - min; tmp > res {
			res = tmp
		}
	}
	return res
}

// @lc code=end
