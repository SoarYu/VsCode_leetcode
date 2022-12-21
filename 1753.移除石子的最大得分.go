/*
 * @lc app=leetcode.cn id=1753 lang=golang
 *
 * [1753] 移除石子的最大得分
 */
package daily

import "sort"

// @lc code=start
func maximumScore(a int, b int, c int) int {
	arr := []int{a, b, c}
	sort.Ints(arr)
	a, b, c = arr[0], arr[1], arr[2]
	if a + b <= c {
		return a + b
	}
	return (a + b + c) / 2
}

// @lc code=end
