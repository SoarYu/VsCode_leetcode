/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */
package daily

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
	zeroHash, oneHash := make(map[string]int), make(map[string]int)
	for i := range strs {
		for j := range strs[i] {
			if strs[i][j] == '0' {
				zeroHash[strs[i]]++
			} else {
				oneHash[strs[i]]++
			}
		}
	}

	length := len(strs)
	dp := make([][][]int, length+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}

	for i := range strs {
		zeros, ones := zeroHash[strs[i]], oneHash[strs[i]]
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[i+1][j][k] = dp[i][j][k]
				if j > zeros && k >= ones {
					dp[i+1][j][k] = max(dp[i+1][j][k], dp[i][j-zeros][k-ones]+1)
				}
			}
		}
	}

	return dp[length][m][n]
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// @lc code=end
