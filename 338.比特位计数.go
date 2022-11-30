/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */
package hot

// @lc code=start
func countBits(n int) []int {
	// 0 1  10 11 100 101 110 111 1000
	//
	res := make([]int, n+1)
	for i, pow := 1, 1; i <= n; i++ {
		for j := 1; j <= pow; j *= 2 {
			if i&j == j {
				res[i]++
			}
		}
		if i == pow {
			pow *= 2
		}
	}
	return res
}

func countBitsDP(n int) []int {
	bits := make([]int, n+1)
	if n == 0 {
		return bits
	}
	for i, high := 1, 1; i <= n; i++ {
		if i&(i-1) == 0 {
			bits[i] = 1
			high = i
		} else {
			bits[i] = bits[i-high] + 1
		}
	}

	return bits
}

// @lc code=end
