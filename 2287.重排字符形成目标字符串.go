/*
 * @lc app=leetcode.cn id=2287 lang=golang
 *
 * [2287] 重排字符形成目标字符串
 */
package daily

// @lc code=start
func rearrangeCharacters(s string, target string) int {
	ans := len(s)
	m1, m2 := make(map[byte]int), make(map[byte]int)
	for i := range target {
		m1[target[i]]++
	}
	for i := range s {
		m2[s[i]]++
	}
	for k, v1 := range m1 {
		if v2, hasv := m2[k]; hasv && v2 >= v1 {
			if v2/v1 < ans {
				ans = v2 / v1
			}
		} else {
			return 0
		}
	}
	return ans
}

// @lc code=end
