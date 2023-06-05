/*
 * @lc app=leetcode.cn id=2351 lang=golang
 *
 * [2351] 第一个出现两次的字母
 */
package daily

// @lc code=start
func repeatedCharacter(s string) byte {
	var arr [26]int
	for i := range s {
		arr[s[i]-'a']++
		if arr[s[i]-'a'] >= 2 {
			return s[i]
		}
	}
	return 0
}

// @lc code=end
