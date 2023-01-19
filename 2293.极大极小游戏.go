/*
 * @lc app=leetcode.cn id=2293 lang=golang
 *
 * [2293] 极大极小游戏
 */
package daily
// @lc code=start
func minMaxGame(nums []int) int {
    min := func(a, b int) int {
        if a <= b {
            return a
        }
        return b
    }
    max := func(a, b int) int {
        if a >= b {
            return a
        }
        return b
    }
    for n := len(nums) / 2; n >= 1; n = n / 2 {
        for i := 0; i< n; i++ {
            if i % 2 == 0 {
                nums[i] = min(nums[2*i], nums[2*i+1])
            } else {
                nums[i] = max(nums[2*i], nums[2*i+1])
            }
        }
    }
    return nums[0]
}
// @lc code=end

