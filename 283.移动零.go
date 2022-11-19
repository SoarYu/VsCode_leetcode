/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */
package main

// @lc code=start
func moveZeroes(nums []int) {
	cur := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[cur] = nums[i]
			cur++
		}
	}
	for cur < len(nums) {
		nums[cur] = 0
		cur++
	}
}

func moveZeroes1(nums []int) {
	for i, j := 0, 1; i < len(nums); i++ {
		if nums[i] != 0 {
			continue
		} else {
			if j <= i {
				j = i + 1
			}
			for j < len(nums) && nums[j] == 0 {
				j++
			}
			if j >= len(nums) {
				return
			}
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

// @lc code=end
