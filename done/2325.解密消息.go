/*
 * @lc app=leetcode.cn id=2325 lang=golang
 *
 * [2325] 解密消息
 */
package daily

// @lc code=start
func decodeMessage(key string, message string) string {
	m := make(map[byte]byte)
	nums := 0
	for i := range key {
		if key[i] == ' ' {
			continue
		}
		if _, hasv := m[key[i]]; !hasv {
			m[key[i]] = byte('a' + nums)
			nums++
		}
	}
	ans := make([]byte, len(message))
	for i := range message {
		if message[i] == ' ' {
			ans[i] = ' '
		} else {
			ans[i] = m[message[i]]
		}
	}
	return string(ans)
}

// @lc code=end
