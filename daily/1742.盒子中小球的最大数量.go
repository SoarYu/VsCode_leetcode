package daily

// @lc code=start
func countBalls(lowLimit int, highLimit int) int {
	// 模拟
	m := make(map[int]int)
	res := 0
	mindig, maxdig := 0, 0
	dp := []int{}
	// 123 12345	->  32100
	for x := lowLimit; x > 0; x /= 10 {
		mindig++
		dp = append(dp, x%10)
	}
	for x := highLimit; x > 0; x /= 10 {
		maxdig++
		if maxdig > mindig {
			dp = append(dp, 0)
		}
	}
	dp = append(dp, 0)

	for i := lowLimit; i <= highLimit; i++ {
		sum := 0
		for j := range dp {
			sum += dp[j]
		}
		if m[sum]++; m[sum] > res {
			res = m[sum]
		}
		dp[0]++
		for j := 0; dp[j] >= 10; j++ {
			dp[j+1]++
			dp[j] -= 10
		}
	}
	return res
}

// @lc code=end
