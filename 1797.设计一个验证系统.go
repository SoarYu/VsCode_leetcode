/*
 * @lc app=leetcode.cn id=1797 lang=golang
 *
 * [1797] 设计一个验证系统
 */
package daily

import "math"

func pickGifts(gifts []int, k int) int64 {
	for i := 0; i < k; i++ {
		// 找出最大
		index := 0
		for j := range gifts {
			if gifts[j] > gifts[index] {
				index = j
			}
		}
		gifts[index] = int(math.Sqrt(float64(gifts[index])))
	}

	var ans int
	for i := range gifts {
		ans += gifts[i]
	}
	return int64(ans)
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
// @lc code=end
