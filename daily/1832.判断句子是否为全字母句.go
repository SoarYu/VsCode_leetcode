/*
 * @lc app=leetcode.cn id=1832 lang=golang
 *
 * [1832] 判断句子是否为全字母句
 */
package daily

// @lc code=start
func checkIfPangram(sentence string) bool {
	var arr [26]int
	for _, v := range sentence {
		arr[v-'a']++
	}

	for _, v := range arr {
		if v == 0 {
			return false
		}
	}

	return true
}

// @lc code=end
