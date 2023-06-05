/*
 * @lc app=leetcode.cn id=945 lang=golang
 *
 * [945] 使数组唯一的最小增量
 */
package daily

import "sort"

// @lc code=start
func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	var ans int

	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			t := nums[i-1] + 1
			ans += t - nums[i]
			nums[i] = t
		}
	}
	return ans
}

// @lc code=end
