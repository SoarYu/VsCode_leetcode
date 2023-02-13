/*
 * @lc app=leetcode.cn id=1049 lang=golang
 *
 * [1049] 最后一块石头的重量 II
 */
package daily

// @lc code=start
func lastStoneWeightII(stones []int) int {
	var ans int

	return ans
}

/*

有 N 件物品和一个容量是 V 的背包。每件物品有且只有一件。

求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，且总价值最大。
*/
func np(values, weights []int, capacity int) int {
	n := len(values)
	dp := make([][]int, n+1) // 0 -> n
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, capacity+1) // 0 -> capacity
	}

	// dp[0][0] = 0  dp[1][0]=0 枚举第0件
	//

	// dp[i][j] 表示在 j 容量下，价值最大的背包
	// 「当前枚举到哪件物品」和「现在的剩余容量」

	// capacity=5 [4,2,3] [4,2,3]
	//  0 1 2 3 4 5
	//0 0 0 0 0 0 0
	//2 0 0 2 2 2 2
	//3 0 0 2 3 3 5
	//4 0 0 2 3 4 5
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i <= n; i++ {
		// 物品 i-1, 重量为 weights[i-1] 价值为values[i-1]
		w, v := weights[i-1], values[i-1]
		// 容量为 j 情况下， 是否要装下物品i-1
		for j := 1; j <= capacity; j++ {
			// 
			if j < w {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-w]+v)
			}
		}
	}

	return dp[n][capacity]
}

// @lc code=end
