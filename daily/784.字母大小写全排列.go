package daily

// @lc code=start
func letterCasePermutation(s string) []string {
	//
	resMap := make(map[string]bool)
	n := len(s)
	bArr := []rune(s)
	var dfs func(index int)
	dfs = func(index int) {
		resMap[string(bArr)] = true
		if index > n {
			return
		}
		for i := index; i < n; i++ {
			if s[i] >= 'a' && s[i] <= 'z' {
				bArr[i] += 'A' - 'a'
				dfs(i + 1)
				bArr[i] -= 'A' - 'a'
			} else if s[i] >= 'A' && s[i] <= 'Z' {
				bArr[i] -= 'A' - 'a'
				dfs(i + 1)
				bArr[i] += 'A' - 'a'
			}
		}
	}
	dfs(0)
	res := []string{}
	for key := range resMap {
		res = append(res, key)
	}
	return res
}

// @lc code=end
