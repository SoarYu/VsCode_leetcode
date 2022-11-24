package daily

// @lc code=start
func maxRepeating(sequence string, word string) int {
	m, n := len(sequence), len(word) // 5 - 2 = 3
	res := 0
	for i := 0; i <= m-n; i++ {
		tmp := 0
		for i <= m-n && sequence[i:i+n] == word {
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
