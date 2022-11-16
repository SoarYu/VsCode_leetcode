/*
 * @lc app=leetcode.cn id=775 lang=golang
 *
 * [775] 全局倒置与局部倒置
 */
package main

// @lc code=start
func isIdealPermutation(nums []int) bool {
	n := len(nums)
	if n <= 2 {
		return true
	}
	// dp := make([]int, n)
	m := nums[n-1]
	for i := n - 3; i >= 0; i-- {
		if nums[i] > m { // 存在局部
			return false
		}
		if nums[i+1] < m {
			m = nums[i+1]
		}

	}
	return true
}

func isIdealPermutation1(nums []int) bool {
	n := len(nums)
	if n <= 2 {
		return true
	}
	for i := 0; i < n; i++ {
		if tmp := nums[i] - i; tmp > 1 || tmp < -1 {
			return false
		}
	}
	return true
}

// @lc code=end
