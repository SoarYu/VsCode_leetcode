/*
 * @lc app=leetcode.cn id=1785 lang=golang
 *
 * [1785] 构成特定和需要添加的最少元素
 */
package daily
// @lc code=start
func minElements(nums []int, limit int, goal int) int {
    sum := 0
    for i := range nums {
        sum += nums[i]
    }
    diff := goal - sum
    if diff < 0 {
        diff = -diff
    }
    if diff % limit == 0 {
        return diff / limit
    }
    return diff / limit + 1
}
// @lc code=end

