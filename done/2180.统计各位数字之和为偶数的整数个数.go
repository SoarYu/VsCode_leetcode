/*
 * @lc app=leetcode.cn id=2180 lang=golang
 *
 * [2180] 统计各位数字之和为偶数的整数个数
 */
package daily

// @lc code=start
func countEven(num int) int {
	var ans int
	for i := 1; i <= num; i++ {
		sum := 0
		for j := i; j > 0; j /= 10 {
			sum += j % 10
		}
		if sum%2 == 0 {
			ans++
		}
	}
	return ans
}

// @lc code=end
