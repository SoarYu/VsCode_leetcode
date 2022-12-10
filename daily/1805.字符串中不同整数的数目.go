/*
 * @lc app=leetcode.cn id=1805 lang=golang
 *
 * [1805] 字符串中不同整数的数目
 */
package daily

// @lc code=start
func numDifferentIntegers(word string) int {
	m := make(map[string]bool)
	ans := 0
	for i := 0; i < len(word); i++ {
		if word[i] < '0' || word[i] > '9' {
			continue
		}
		num := []byte{}
		for ; i < len(word) && word[i] >= '0' && word[i] <= '9'; i++ {
			if len(num) == 0 && word[i] == '0' {
				continue
			}
			num = append(num, word[i])
		}
		if !m[string(num)] {
			m[string(num)] = true
			ans++
		}
	}
	return ans
}

// @lc code=end
