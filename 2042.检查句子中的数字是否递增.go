/*
 * @lc app=leetcode.cn id=2042 lang=golang
 *
 * [2042] 检查句子中的数字是否递增
 */
package daily

func areNumbersAscending(s string) bool {
	pre := 0
	for i := 0; i < len(s); {
		if s[i] < '0' || s[i] > '9' {
			i++
			continue
		}
		sum := 0
		for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
			sum = sum*10 + int(s[i]-'0')
		}
		if sum <= pre {
			return false
		}
		pre = sum
	}
	return true
}

// @lc code=end
