package daily

// @lc code=start
func numMatchingSubseq(s string, words []string) int {
	// 记录words的首字符
	res := 0
	dict := [26][]string{}
	for _, word := range words {
		dict[word[0]-'a'] = append(dict[word[0]-'a'], word)
	}
	for _, ch := range s {
		tmp := dict[ch-'a']
		dict[ch-'a'] = nil
		for _, word := range tmp {
			if len(word) == 1 {
				res++
			} else {
				dict[word[1]-'a'] = append(dict[word[1]-'a'], word[1:])
			}
		}
	}
	return res
}

// @lc code=end
