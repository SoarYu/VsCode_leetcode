package daily

// @lc code=start
func parseBoolExpr(expression string) bool {
	opsStack := []rune{}
	t, f := 0, 0
	for _, ch := range expression {
		if ch == 't' {
			t++
		} else if ch == 'f' {
			f++
		} else if ch == '!' || ch == '&' || ch == '|' || ch == '(' {
			opsStack = append(opsStack, ch)
			f, t = 0, 0
		} else if ch == ')' {
			ops := opsStack[len(opsStack)-1]
			opsStack = opsStack[:len(opsStack)-1]
			for ops != '(' {
				if ops == 't' {
					t++
				} else if ops == 'f' {
					f++
				}
				ops = opsStack[len(opsStack)-1]
				opsStack = opsStack[:len(opsStack)-1]
			}
			ops = opsStack[len(opsStack)-1]
			opsStack = opsStack[:len(opsStack)-1]
			if ops == '!' {
				if t > 0 {
					opsStack = append(opsStack, 'f')
				} else if f > 0 {
					opsStack = append(opsStack, 't')
				}
			} else if ops == '&' {
				if f > 0 {
					opsStack = append(opsStack, 'f')
				} else {
					opsStack = append(opsStack, 't')
				}
			} else if ops == '|' {
				if t > 0 {
					opsStack = append(opsStack, 't')
				} else {
					opsStack = append(opsStack, 'f')
				}
			}
			f, t = 0, 0
		}
	}
	return opsStack[0] == 't'
}

// @lc code=end
