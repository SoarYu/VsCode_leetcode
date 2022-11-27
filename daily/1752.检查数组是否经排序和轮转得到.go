/*
 * @lc app=leetcode.cn id=1752 lang=golang
 *
 * [1752] 检查数组是否经排序和轮转得到
 */
package daily

func check(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	// 递增区间 最多两个, 且第一个区间的最小值大于第二个区间的最大值 [3 4 5 1 4]
	i := 0
	for i+1 < len(nums) && nums[i+1] >= nums[i] {
		i++
	}
	// true : i=len(nums)-1 false: i=len(nums)
	// i=len(nums)-1
	if i == len(nums)-1 {
		return true
	}
	// nums[i] > nums[i+1]
	// left[0, i] right[i+1,]
	i++
	for i+1 < len(nums) && nums[i+1] >= nums[i] {
		i++
	}
	if i != len(nums)-1 || nums[i] > nums[0] {
		return false
	}
	return true
}

// @lc code=end
