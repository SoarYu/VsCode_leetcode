/*
 * @lc app=leetcode.cn id=823 lang=golang
 *
 * [823] 带因子的二叉树
 */
package daily

import "sort"

// @lc code=start
func numFactoredBinaryTrees(arr []int) int {
	// 动态规划
	var ans int
	sort.Ints(arr)
	hasVal := make(map[int]int)
	// 确定根节点 arr[i] >= 2
	for i := range arr {
		hasVal[arr[i]]++
		for j := 0; j < i && arr[j] <= arr[i]/2; j++ {
			// 确定左子树 为 arr[j]， 右子树为 arr[i] / arr[j]
			if arr[i]%arr[j] == 0 && hasVal[arr[i]/arr[j]] != 0 {
				hasVal[arr[i]] += hasVal[arr[j]] * hasVal[arr[i]/arr[j]]
			}
		}
		ans += hasVal[arr[i]]
	}
	return ans % (1e9 + 7)
}

// @lc code=end
