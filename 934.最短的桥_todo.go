/*
 * @lc app=leetcode.cn id=934 lang=golang
 *
 * [934] 最短的桥
 */

// @lc code=start
func shortestBridge(grid [][]int) int {
	type pair struct{ x, y int }
	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n := len(grid)
}
// @lc code=end

