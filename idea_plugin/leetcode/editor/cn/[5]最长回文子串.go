//给你一个字符串 s，找到 s 中最长的回文子串。 
// 如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
//
// 示例 1：
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
// 示例 2： 
//输入：s = "cbbd"
//输出："bb"
//
// 提示： 
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成 
// Related Topics 字符串 动态规划 👍 6561 👎 0
package main
import "sort"
//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) (res string) {
	length := len(s)
	n := len(s)
	for length > 0 {
		for i := 0; i + length <= n; i++ {
			if isVaild(s[i:i+length]) {
				res = s[i:i+length]
				return
			}
		}
		length--
	}
	return
}

func isVaild(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
