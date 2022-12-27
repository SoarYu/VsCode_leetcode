/*
 * @lc app=leetcode.cn id=1759 lang=golang
 *
 * [1759] 统计同构子字符串的数目
 */
package daily

// @lc code=start
func countHomogenous(s string) int {
	var ans int
	ans += len(s)
	for i := 0; i < len(s); {
		j := i
		for j+1 < len(s) && s[j+1] == s[j] {
			j++
		}
		ans += (1+j-i+1) * (j-i+1) / 2
		i = j + 1
	}
	return ans
}

// @lc code=end

// aaaa 10  (1+4)*4 / 2
/*
a a a a
aa aa aa
aaa aaa
aaaa
*/

// aaaaa 15 = (1+5)*5 / 2
/*
a a a a a
aa aa aa aa
aaa aaa aaa
aaaa aaaa
aaaaa
*/
