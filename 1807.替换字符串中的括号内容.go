/*
 * @lc app=leetcode.cn id=1807 lang=golang
 *
 * [1807] 替换字符串中的括号内容
 */
package daily

// @lc code=start
func evaluate(s string, knowledge [][]string) string {
	kvMap := make(map[string]string)
	for _, kv := range knowledge {
		kvMap[kv[0]] = kv[1]
	}
	ans := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
            word := ""
			for i++; s[i] != ')'; i++ {
				word += string(s[i])
			}
			// s[i] = ')'
			if val, hasVal := kvMap[word]; hasVal {
    			ans += val
            } else {
                ans += "?"
            }
		} else {
			ans += string(s[i])
		}
	}
	return ans
}

// @lc code=end
