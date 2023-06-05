package daily

func vowelStrings(words []string, queries [][]int) []int {
	// 前缀和
	isVowel := func(b byte) bool {
		if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' {
			return true
		}
		return false
	}

	isValid := make(map[int]int)
	// 前缀和
	preSum := make([]int, len(words))
	for i, s := range words {
		if isVowel(s[0]) && isVowel(s[len(s)-1]) {
			preSum[i]++
			isValid[i]++
		}
		if i+1 < len(words) {
			preSum[i+1] = preSum[i]
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = preSum[q[1]] - preSum[q[0]] + isValid[q[0]]
	}

	return ans
}
