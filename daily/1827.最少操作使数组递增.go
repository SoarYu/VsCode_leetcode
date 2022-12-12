/*
 * @lc app=leetcode.cn id=1827 lang=golang
 *
 * [1827] 最少操作使数组递增
 */
package daily

// @lc code=start
func minOperations(nums []int) int {
	var ans int
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			ans += nums[i-1] - nums[i]
			nums[i] += nums[i-1] - nums[i]
		}
	}
	return ans
}

// @lc code=end
