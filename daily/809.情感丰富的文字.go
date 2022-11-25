/*
 * @lc app=leetcode.cn id=809 lang=golang
 *
 * [809] 情感丰富的文字
 */
package daily

// @lc code=start
func expressiveWords(s string, words []string) int {
	// 模拟
	res := 0
	for _, word := range words {
		// word与s比较
		i, j, ok := 0, 0, true
		for ; i < len(s) && j < len(word); i, j = i+1, j+1 {
			if s[i] != word[j] {
				ok = false
				break
			}
			// s[i] = word[j] 计算连续相同字符的长度
			lenS, lenW := 1, 1
			for ; i+1 < len(s) && s[i+1] == s[i]; i++ {
				lenS++
			}
			for ; j+1 < len(word) && word[j+1] == word[j]; j++ {
				lenW++
			}

			if lenS < lenW || (lenS > lenW && lenS < 3) {
				ok = false
				break
			}

		}
		// 是否完成遍历
		if i == len(s) && j == len(word) && ok {
			res++
		}
	}
	return res
}

// @lc code=end
