/*
 * @lc app=leetcode.cn id=926 lang=golang
 *
 * [926] 将字符串翻转到单调递增
 */
package daily

// @lc code=start
func minFlipsMonoIncr(s string) int {

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	left, right := 0, len(s)-1
	for left < len(s) {
		if s[left] == '0' {
			left++
		} else {
			break
		}
	}
	for right >= 0 {
		if s[right] == '1' {
			right--
		} else {
			break
		}
	}
	if left >= right {
		return 0
	}
	// s[left]=1
	// s[right]=0
	var c1, c0 int
	for i := left; i <= right; i++ {
		if s[i] == '0' {
			c0++
		} else {
			c1++
		}
	}
	return min(c0, c1)
}

// "100111111100000000000010111011"
// @lc code=end
