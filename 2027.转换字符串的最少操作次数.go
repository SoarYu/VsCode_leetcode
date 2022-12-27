/*
 * @lc app=leetcode.cn id=2027 lang=golang
 *
 * [2027] 转换字符串的最少操作次数
 */
package daily

// @lc code=start
func minimumMoves(s string) int {
	var ans int
	for i := 0; i < len(s); {
		if s[i] == 'O' {
			i++
			continue
		}
		ans++
		i += 3
	}
	return ans
}

// @lc code=end
