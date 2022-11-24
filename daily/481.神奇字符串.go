package daily

import (
	"bytes"
)

/*
 * @lc app=leetcode.cn id=481 lang=golang
 *
 * [481] 神奇字符串
 */

// @lc code=start
func magicalString(n int) int {
	s := make([]byte, 0, n+1)
	s = append(s, 1, 2, 2)
	c := []byte{2} // [2]

	for i := 2; len(s) < n; i++ {
		c[0] ^= 3 // 1^3=2, 2^3=1，这样就能在 1 和 2 之间转换 || 01^11 = 10  10^11 = 01
		s = append(s, bytes.Repeat(c, int(s[i]))...)
	}
	return bytes.Count(s[:n], []byte{1})
}

// @lc code=end
