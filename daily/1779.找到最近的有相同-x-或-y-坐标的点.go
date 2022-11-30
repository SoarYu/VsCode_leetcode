/*
 * @lc app=leetcode.cn id=1779 lang=golang
 *
 * [1779] 找到最近的有相同 X 或 Y 坐标的点
 */
package daily

import "math"

// @lc code=start
func nearestValidPoint(x int, y int, points [][]int) int {
	ans, min := -1, math.MaxInt64
	var abs func(int) int
	abs = func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	for i, p := range points {
		if p[0] == x || p[1] == y {
			if tmp := abs(x-p[0]) + abs(y-p[0]); tmp < min {
				min = tmp
				ans = i
			}
		}
	}
	return ans
}

// @lc code=end
