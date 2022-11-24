package daily

// @lc code=start
func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	// ruleKey = type \ color \ name
	var index int
	if ruleKey == "type" {
		index = 0
	} else if ruleKey == "color" {
		index = 1
	} else {
		index = 2
	}
	ans := 0
	for _, item := range items {
		if item[index] == ruleValue {
			ans++
		}
	}
	return ans
}

// @lc code=end
