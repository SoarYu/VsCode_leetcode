/*
 * @lc app=leetcode.cn id=1750 lang=golang
 *
 * [1750] 删除字符串两端相同字符后的最短长度
 */
package daily

// @lc code=start
func minimumLength(s string) int {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
			// left
			for left <= right && s[left] == s[left-1] {
				left++
			}

			// right
			for left <= right && s[right] == s[right+1] {
				right--
			}

		} else {
			break
		}
	}
	if left > right {
		return 0
	}

	return right - left + 1
}

// @lc code=end
