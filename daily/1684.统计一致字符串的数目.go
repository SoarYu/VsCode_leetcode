package daily

// @lc code=start
func countConsistentStrings(allowed string, words []string) (ans int) {
	m := make(map[rune]bool)
	for _, ch := range allowed {
		m[ch] = true
	}
	for _, word := range words {
		isVaild := true
		for _, ch := range word {
			if !m[ch] {
				isVaild = false
				break
			}
		}
		if isVaild {
			ans++
		}
	}
	return
}
// @lc code=end

