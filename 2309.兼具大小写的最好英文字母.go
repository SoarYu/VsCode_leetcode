/*
 * @lc app=leetcode.cn id=2309 lang=golang
 *
 * [2309] 兼具大小写的最好英文字母
 */
package daily

// @lc code=start
func greatestLetter(s string) string {
	m := make(map[byte]bool)
	for i := range s {
		m[s[i]] = true
	}
	for i := 25; i >= 0; i-- {
		if m[byte('a'+i)] && m[byte('A'+i)] {
			return string('A' + i)
		}
	}
	return ""
}

// @lc code=end
