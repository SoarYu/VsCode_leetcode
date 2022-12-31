/*
 * @lc app=leetcode.cn id=2037 lang=golang
 *
 * [2037] 使每位学生都有座位的最少移动次数
 */
package daily

import "sort"

// @lc code=start
func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	var ans int
	for i := range seats {
		ans += abs(seats[i] - students[i])
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end
