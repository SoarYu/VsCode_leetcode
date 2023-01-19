/*
 * @lc app=leetcode.cn id=1814 lang=golang
 *
 * [1814] 统计一个数组中好对子的数目
 */
package daily
// @lc code=start
func countNicePairs(nums []int) int {
    rev := func(x int) int {
        res := 0
        for x != 0 {
            res = res * 10 + x % 10
            x /= 10
        }
        return res
    }
    countmap := make(map[int]int)
    var ans int
    // 设存在一个数 j 使 i , 有下式成立
    // nums[i] + rev(nums[j]) = nums[j] + rev(nums[i])
    // i, j 符合条件： nums[i] - rev(nums[i]) = nums[j] - rev(nums[j]) = targetNum 配对加一
    for _, num := range nums {
        targetNum := num - rev(num)
        ans += countmap[targetNum]
        countmap[targetNum]++
    }
    return ans % (1e9+7)
}
// @lc code=end

