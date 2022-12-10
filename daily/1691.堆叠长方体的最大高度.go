/*
 * @lc app=leetcode.cn id=1691 lang=golang
 *
 * [1691] 堆叠长方体的最大高度
 */
package daily

import "sort"

// @lc code=start
func maxHeight(cuboids [][]int) int {
	for i := range cuboids {
		sort.Ints(cuboids[i])
	}

	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][0] == cuboids[j][0] && cuboids[i][1] == cuboids[j][1] {
			return cuboids[i][2] > cuboids[j][2]
		} else if cuboids[i][0] == cuboids[j][0] {
			return cuboids[i][1] > cuboids[j][1]
		}
		return cuboids[i][0] > cuboids[j][0]
	})

	n := len(cuboids)
	dp := make([]int, n)
	dp[0] = cuboids[0][2]

	for i := 1; i < n; i++ {
		dp[i] = cuboids[i][2]
		for j := 0; j < i; j++ {
			if cuboids[j][1] >= cuboids[i][1] && cuboids[j][2] >= cuboids[i][2] {
				dp[i] = max(dp[i], dp[j]+cuboids[i][2])
			}
		}
	}

	var ans int
	for _, v := range dp {
		ans = max(ans, v)
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
// @lc code=end
