/*
 * @lc app=leetcode.cn id=941 lang=golang
 *
 * [941] 有效的山脉数组
 */
package daily

// @lc code=start
func validMountainArray(arr []int) bool {
	// 严格递增
	i := 0
	for ; i < len(arr)-1 && arr[i+1] > arr[i]; i++ {
	}
	if i == 0 || i == len(arr)-1 {
		return false
	}
	// 严格递减
	for ; i < len(arr)-1 && arr[i] > arr[i+1]; i++ {
	}
	return i == len(arr)-1
}

// @lc code=end
