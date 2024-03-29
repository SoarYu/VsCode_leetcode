package daily

// @lc code=start
func arraySign(nums []int) int {
	neg := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			neg++
		}
	}
	if neg%2 == 1 {
		return -1
	}
	return 1
}

// @lc code=end
