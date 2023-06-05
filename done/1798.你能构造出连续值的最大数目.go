/*
 * @lc app=leetcode.cn id=1798 lang=golang
 *
 * [1798] 你能构造出连续值的最大数目
 */
package daily

import "sort"

// @lc code=start
func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)
	var r int
	for _, v := range coins {
		if v <= r || v == r+1 {
			r += v
		} else {
			break
		}
	}
	return r + 1
}

// @lc code=end
