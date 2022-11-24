package daily

// @lc code=start
func interpret(command string) string {
	res := ""
	for k, ch := range command {
		if ch == 'G' {
			res += "G"
		} else if ch == ')' {
			if command[k-1] == '(' {
				res += "o"
			} else {
				res += "al"
			}
		}
	}
	return res
}
// @lc code=end

