package daily

// @lc code=start
func numSubarrayBoundedMax(nums []int, left int, right int) int {
	// 影响范围
	count , k1, k2 := 0, -1, -1
	for i, v := range nums {
		if v > right {
			k2 = i
		} else if v >= left {
			k1 = i
		}
		if k1 > k2 {
			count += k1-k2
		}
	}
	return count
}

// 2 2 1 4 3 3    
// @lc code=end

/*
2 1 4 3 2 1                    2 3
1*2 + 1*3 + 2*2 = 2 3 4 = 9
[2] [2 1] [3] [32] [321] [2] [21]
*/