/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */
package daily

// @lc code=start

func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	var sum, max int
	for i := range nums {
		sum += nums[i]
		if nums[i] > max {
			max = nums[i]
		}
	}

	if sum%2 != 0 || max > sum/2 {
		return false
	}

	n, target := len(nums), sum/2

	// dp := make([][]bool, n)
	// for i := range dp {
	// 	dp[i] = make([]bool, target+1)
	// }

	// for i := range dp {
	// 	dp[i][0] = true
	// }
	// 元素0， 直接true 表示值存在
	values := make([]bool, target+1)
	values[0], values[nums[0]] = true, true
	for i := 1; i < n; i++ {
		// 元素 i 是否装入
		newValues := make([]bool, target+1)
		copy(newValues, values)
		for j := 0; j <= target; j++ {
			if values[j] {
				if j+nums[i] == target {
					return true
				} else if j+nums[i] < target {
					newValues[j+nums[i]] = true
				}
			}
		}
		values = newValues
	}

	return values[target]
}

// @lc code=end
