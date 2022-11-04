/*
 * @lc app=leetcode.cn id=1668 lang=golang
 *
 * [1668] 最大重复子字符串
 */

// @lc code=start
func maxRepeating(sequence string, word string) int {
	m, n := len(sequence), len(word) // 5 - 2 = 3  
	res := 0
	for i:=0; i<=m-n; i++{
		tmp := 0
		for i<=m-n && sequence[i:i+n] == word {
			tmp++
			i += n
		}
		if tmp > 0 {
			if tmp > res {
				res = tmp
			}
			i -= n
		}
	}
	return res
}
// @lc code=end

/*
交汇点只有一个
"aaabaaaabaaaaba aabaaaabaaaabaaaabaaaaba"
"aaaba"
*/