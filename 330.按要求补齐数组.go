/*
 * @lc app=leetcode.cn id=330 lang=golang
 *
 * [330] 按要求补齐数组
 */
package daily

// @lc code=start
func minPatches(nums []int, n int) int {
	var r int
	var ans int
	for i := 0; i < len(nums); {
		if r >= n {
			break
		}

		if nums[i] <= r || nums[i] == r+1 {
			r += nums[i]
			i++
		} else {
			ans++
			r += (r + 1)
		}
	}

	for r < n {
		r += (r + 1)
		ans++
	}

	return ans
}

// @lc code=end
