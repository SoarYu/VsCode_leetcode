/*
 * @lc app=leetcode.cn id=2299 lang=golang
 *
 * [2299] 强密码检验器 II
 */
package daily

// @lc code=start
func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}
	var low, high, num, spec int
	specialString := "!@#$%^&*()-+"
	for i := range password {
		if i > 0 && password[i] == password[i-1] {
			return false
		}
		if password[i] >= 'a' && password[i] <= 'z' {
			low++
		} else if password[i] >= 'A' && password[i] <= 'Z' {
			high++
		} else if password[i] >= '0' && password[i] <= '9' {
			num++
		} else {
			for j := range specialString {
				if password[i] == specialString[j] {
					spec++
					break
				}
			}
		}
	}
	return low > 0 && high > 0 && num > 0 && spec > 0
}

// @lc code=end
