package daily

// 模拟过程
// @lc code=start
func champagneTower(poured int, query_row int, query_glass int) float64 {
	dp := make([][]float64, query_row+1)
	for i := 0; i < query_row+1; i++ {
		dp[i] = make([]float64, i+1)
	}
	dp[0][0] = float64(poured)
	for i := 0; i < query_row; i++ {
		for j, cup := range dp[i] {
			if cup > 1 {
				// 向两侧溢出
				cup -= 1
				dp[i+1][j] += cup / 2
				dp[i+1][j+1] += cup / 2
			}
		}
	}
	if val := dp[query_row][query_glass]; val > 1 {
		return 1
	} else {
		return val
	}

}

// @lc code=end
