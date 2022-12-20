/*
 * @lc app=leetcode.cn id=1760 lang=golang
 *
 * [1760] 袋子里最少数目的球
 */
package daily

import "sort"

// @lc code=start
// 二分查找
func minimumSize(nums []int, maxOperations int) int {
	var max int
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return sort.Search(max, func(maxSize int) bool {
		if maxSize <= 0 {
			return false
		}
		var ops int
		for _, size := range nums {
			if size > maxSize {
				ops += (size - 1) / maxSize
			}
		}
		return ops <= maxOperations
	})
}

// @lc code=end
