/*
 * @lc app=leetcode.cn id=1124 lang=golang
 *
 * [1124] 表现良好的最长时间段
 */
package daily

// @lc code=start
func longestWPI(hours []int) int {
	preSum := make(map[int]int)
	var sum int
	var ans int
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i, hour := range hours {
		if hour > 8 {
			sum++
		} else {
			sum--
		}

		if sum > 0 {
			ans = max(ans, i+1)
		} else {
			// sum <= 0
			if j, hasVal := preSum[sum-1]; hasVal {
				ans = max(ans, i - j)
			}
		}
		if _, hasVal := preSum[sum]; !hasVal {
			preSum[sum] = i
		}
	}

	return ans
}

// @lc code=end

// 9 9 6 0 9 6
// 1 2 1 0 1 0

// 6   0  9  9 9 0  0  0  9
// -1 -2 -1  0 1 0 -1 -2 -1
// 0  1  0   3 4 3 0  1   0