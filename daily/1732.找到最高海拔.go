package daily
// @lc code=start
func largestAltitude(gain []int) int {
	max, cur := 0, 0
	for _, h := range gain {
		if cur += h; cur > max {
			max = cur
		}
	}
	return max
}
// @lc code=end

