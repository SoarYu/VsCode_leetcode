/*
 * @lc app=leetcode.cn id=1658 lang=golang
 *
 * [1658] 将 x 减到 0 的最小操作数
 */
/*
 // 遇到大于等于x的，直接堵住
 // 二分， l <= x, r <= x  l+r <=x
*/
package daily

// @lc code=start
func minOperations(nums []int, x int) int {
	suml, sumr, l, r := 0, 0, -1, 0
	for _, v := range nums {
		sumr += v
	}
	if sumr < x {
		return -1
	}
	var min func(int, int) int
	min = func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	ans := len(nums) + 1
	for ; l <= len(nums)-1; l++ {
		if l != -1 {
			suml += nums[l]
		}
		for r < len(nums) && suml+sumr > x {
            sumr -= nums[r]
            r++
        }

		if suml + sumr == x {
			ans = min(ans, l + 1 + len(nums) - r)
		}
	}
	if ans > len(nums) {
		return -1
	}
	return ans
}
// @lc code=end
