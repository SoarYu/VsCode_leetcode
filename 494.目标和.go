/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */
package daily

import "sort"

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	if target < 0 {
		target = -target
	}

	sort.Ints(nums)
	var sum int
	for i := range nums {
		sum += nums[i]
	}

	var ans int
	var dfs func(int, int)
	dfs = func(index, sum int) {
		if sum == target {
			ans++
		}
		if sum < target {
			return
		}
		for i:=index+1; i<len(nums); i++ {
			dfs(i, sum)
			dfs(i, sum - nums[index])
		}
	}
	dfs(0, sum)
	return ans
}

// @lc code=end
