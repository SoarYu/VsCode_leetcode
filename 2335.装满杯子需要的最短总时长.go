/*
 * @lc app=leetcode.cn id=2335 lang=golang
 *
 * [2335] 装满杯子需要的最短总时长
 */
package daily

import "sort"

// @lc code=start
func fillCups(amount []int) int {
	// 选择最大的两个
	var ans int
	for {
		// mid min
		sort.Ints(amount)
		if amount[1] <= 0 && amount[0] <= 0 {
			ans += amount[2]
			break
		}
		amount[1] -= 1
		amount[2] -= 1
		ans++
	}
	return ans
}

// @lc code=end
