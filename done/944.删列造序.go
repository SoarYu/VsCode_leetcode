/*
 * @lc app=leetcode.cn id=944 lang=golang
 *
 * [944] 删列造序
 */
package daily

// @lc code=start
func minDeletionSize(strs []string) int {
	rows, cols := len(strs), len(strs[0])
	var ans int
	for col := 0; col < cols; col++ {
		pre := strs[0][col]
		for row := 1; row < rows; row++ {
			if strs[row][col] >= pre {
				pre = strs[row][col]
			} else {
				ans++
				break
			}
		}
	}
	return ans
}

// @lc code=end
