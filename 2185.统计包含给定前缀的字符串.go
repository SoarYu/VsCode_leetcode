/*
 * @lc app=leetcode.cn id=2185 lang=golang
 *
 * [2185] 统计包含给定前缀的字符串
 */
package daily
// @lc code=start
func prefixCount(words []string, pref string) int {
	var ans int
	length := len(pref)
	for _, word := range words {
		if len(word) >= length && word[:length] == pref {
			ans++
		}
	}
	return ans
}
// @lc code=end

