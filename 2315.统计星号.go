/*
 * @lc app=leetcode.cn id=2315 lang=golang
 *
 * [2315] 统计星号
 */
package daily

// @lc code=start
func countAsterisks(s string) int {
	var ans int
	for i := 0; i < len(s); i++ {
		if s[i] == '|' {
			for i++; i < len(s) && s[i] != '|'; i++ {
			}
		} else if s[i] == '*' {
			ans++
		}
	}
	return ans
}

// @lc code=end
