/*
 * @lc app=leetcode.cn id=813 lang=golang
 *
 * [813] 最大平均值和的分组
 */
package daily

import "math"

// @lc code=start
func largestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
	prefix := make([]float64, n+1)
	for i, x := range nums {
		prefix[i+1] = prefix[i] + float64(x)
	}
	dp := make([]float64, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = prefix[i] / float64(i)
	}
	for j := 2; j <= k; j++ { // 区间数量
		for i := n; i >= k; i-- { // 逆序找下一个区间的最大平均和
			for x := j - 1; x < i; x++ { //
				//
				dp[i] = math.Max(dp[i], dp[x]+(prefix[i]-prefix[x])/float64(i-x))
			}
		}
	}
	return dp[n]
}

func largestSumOfAveragesDfs(nums []int, k int) float64 {
	n := len(nums)
	prefix := make([]int, n+1)
	for i, v := range nums {
		// prefix[i+1] = sum(0,i)
		prefix[i+1] = prefix[i] + v
	}
	var dfsRes [101][101]float64

	var dfs func(i, k int) float64
	dfs = func(i, k int) float64 {
		if i == n {
			return 0
		}
		if k == 1 {
			// [i, n] 平均数
			// sum(0, n-1) - sum(0, i-1) = sum(i, n-1)
			return float64(prefix[n]-prefix[i]) / float64(n-i)
		}
		if dfsRes[i][k] > 0 {
			return dfsRes[i][k]
		}
		var res float64
		for j:=i+1; j<n; j++ {
			// sum(i,j)
			t := float64(prefix[j]-prefix[i]) / float64(j-i) + dfs(j, k-1)
			res = math.Max(res, t)
		}
		dfsRes[i][k] = res
		return res
	}
	return dfs(0, k)
}

// @lc code=end
