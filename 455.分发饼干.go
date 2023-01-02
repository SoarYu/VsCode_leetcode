/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */
package daily

import "sort"

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var ans, i, j int
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
			j++
			ans++
		} else {
			j++
		}
	}
	return ans
}

// @lc code=end
