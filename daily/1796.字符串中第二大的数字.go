/*
 * @lc app=leetcode.cn id=1796 lang=golang
 *
 * [1796] 字符串中第二大的数字
 */
package daily

// @lc code=start
/*
 * @lc app=leetcode.cn id=1796 lang=golang
 *
 * [1796] 字符串中第二大的数字
 */
// @lc code=start
func secondHighest(s string) int {
	fir, sec := -1, -1
	for i := range s {
		if n := int(s[i] - '0'); n >= 0 && n <= 9 && n > sec {
			if n > fir {
				fir, sec = n, fir
			} else if n != fir {
				sec = n
			}
		}
	}
	return sec
}

// @lc code=end

// @lc code=end
