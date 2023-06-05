/*
 * @lc app=leetcode.cn id=1664 lang=golang
 *
 * [1664] 生成平衡数组的方案数
 */
package daily

// @lc code=start
func waysToMakeFair(nums []int) (res int) {
	var odd1, even1, odd2, even2 int
	// 奇数偶数后缀和
	for i, num := range nums {
		if i&1 > 0 {
			odd2 += num
		} else {
			even2 += num
		}
	}
	for i, num := range nums {
		// i 的后缀和减去nums[i]
		if i&1 > 0 {
			odd2 -= num
		} else {
			even2 -= num
		}
		// 	// 删除偶数坐标 i 前面不变， 前面偶数 + 后面奇数 = 前面奇数 + 后面偶数
		// 	// 删除奇数坐标 i 前面不变， 前面奇数 + 后面偶数 = 前面偶数 + 后面奇数
		if odd1+even2 == odd2+even1 {
			res++
		}
		// i 的前缀和加上nums[i]
		if i&1 > 0 {
			odd1 += num
		} else {
			even1 += num
		}
	}
	return
}

// @lc code=end
