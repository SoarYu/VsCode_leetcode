/*
 * @lc app=leetcode.cn id=1773 lang=golang
 *
 * [1773] 统计匹配检索规则的物品数量
 */

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

