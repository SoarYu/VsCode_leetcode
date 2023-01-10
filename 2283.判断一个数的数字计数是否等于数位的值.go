/*
 * @lc app=leetcode.cn id=2283 lang=golang
 *
 * [2283] 判断一个数的数字计数是否等于数位的值
 */
package daily

// @lc code=start
func digitCount(num string) bool {
	var a1, a2 [10]int
	for i := range num {
		n := int(num[i] - '0')
		a1[i] = n // a1[i] 希望出现 n 次
		a2[n]++   // a2[i] 出现 a2[i] 次
	}
	for i := range a1 {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

// @lc code=end
