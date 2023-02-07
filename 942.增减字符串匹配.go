/*
 * @lc app=leetcode.cn id=942 lang=golang
 *
 * [942] 增减字符串匹配
 */
package daily

// @lc code=start
func diStringMatch(s string) []int {
	n := len(s)
	ans := make([]int, n+1)
	max, min := n, 0
	// 'I' 递增 'D' 递减
	for i := 0; i < n; i++ {
		// i : 0 - > n - 1
		if s[i] == 'I' {
			ans[i] = min
			min++
		} else {
			ans[i] = max
			max--
		}
	}
	ans[n] = min
	return ans
}

// @lc code=end
