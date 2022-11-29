/*
 * @lc app=leetcode.cn id=1758 lang=golang
 *
 * [1758] 生成交替二进制字符串的最少操作数
 */
package daily

// @lc code=start
func minOperations(s string) int {
	step1, step2 := 0, 0
	// 01010101
	// 10101010
	for i := 0; i < len(s); i++ {
		if i%2 == 0 { // 偶数
			if s[i] == '0' {
				step2++
			} else {
				step1++
			}
		} else { //奇数
			if s[i] == '1' {
				step2++
			} else {
				step1++
			}
		}
	}

	if step2 < step1 {
		return step2
	}
	return step1
}

// 0000100 -> 1010
// 0100100

// 0100

// 0011
// @lc code=end
