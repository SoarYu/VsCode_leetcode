/*
 * @lc app=leetcode.cn id=1945 lang=golang
 *
 * [1945] 字符串转化后的各位数字之和
 */
package daily

// @lc code=start
func getLucky(s string, k int) int {
	sum := 0
	for i := range s {
		n := int(s[i] - 'a' + 1)
		for n != 0 {
			sum += n % 10
			n /= 10
		}
	}
	for i := 1; i < k; i++ {
		tmp := 0
		for sum != 0 {
			tmp += sum % 10
			sum /= 10
		}
		sum = tmp
	}
	return sum
}

// @lc code=end
