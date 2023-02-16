/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */
package daily

// @lc code=start
func generateParenthesis(n int) []string {
	var dfs func(str string, open, close int)
	ans := []string{}
	dfs = func(str string, open, close int) {
		if open == 0 && close == 0 {
			ans = append(ans, str)
			return
		}
		if open == close {
			dfs(str+"(", open-1, close)
		} else { // open < close
			// 可以是close ,可以是open
			// close 是为0
			dfs(str+")", open, close-1)
			if open > 0 {
				dfs(str+"(", open-1, close)
			}
		}
	}
	// 剩余 open 数 <= close数
	dfs("", n, n)
	return ans
}

// @lc code=end
