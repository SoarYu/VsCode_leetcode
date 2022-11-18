/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */
package main

// @lc code=start
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin > i {
				continue
			} else if coin == i {
				dp[i] = 1
			} else {
				if a, b := dp[i], dp[i-coin]; b != 0 && a != 0 && a < b+1 {

					dp[i] = b + 1
				}
			}
		}
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
