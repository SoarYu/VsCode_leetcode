/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */
package hot

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	res := []int{}
	// 多余的数字存放在left
	for i := range nums {
		for k := nums[i]; k != nums[k-1]; k = nums[i] {
			nums[k-1], nums[i] = nums[i], nums[k-1]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}

// @lc code=end
