/*
 * @lc app=leetcode.cn id=2303 lang=golang
 *
 * [2303] 计算应缴税款总额
 */
package daily

// @lc code=start
func calculateTax(brackets [][]int, income int) float64 {
	var ans float64
	pre := 0
	for _, b := range brackets {
		if b[0] > income {
			ans += float64((income-pre)*b[1]) * 0.01
		} else {
			ans += float64((b[0]-pre)*b[1]) * 0.01
		}
		pre = b[0]
		if pre > income {
			break
		}
	}
	return ans
}

// @lc code=end
