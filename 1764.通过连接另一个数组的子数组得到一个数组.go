/*
 * @lc app=leetcode.cn id=1764 lang=golang
 *
 * [1764] 通过连接另一个数组的子数组得到一个数组
 */
package daily

// @lc code=start
// !!! 这些子数组在 nums 中出现的顺序需要与 groups 顺序相同
func canChoose(groups [][]int, nums []int) bool {
	index, count := 0, 0
	for _, g := range groups {
		for index+len(g) <= len(nums) {
			if equal(nums[index:index+len(g)], g) {
				count++
				index += len(g)
				break
			} else {
				index++
			}
		}
	}
	return count == len(groups)
}

func equal(a, b []int) bool {
	for i, x := range a {
		if x != b[i] {
			return false
		}
	}
	return true
}

// @lc code=end
