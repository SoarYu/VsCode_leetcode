/*
 * @lc app=leetcode.cn id=1417 lang=golang
 *
 * [1417] 重新格式化字符串
 */
package daily

// @lc code=start
func reformat(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	letter, num := []byte{}, []byte{}
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			letter = append(letter, s[i])
		} else {
			num = append(num, s[i])
		}
	}
	var max, min []byte
	if len(letter) > len(num) {
		max = letter
		min = num
	} else {
		max = num
		min = letter
	}

	if len(max)-len(min) > 1 {
		return ""
	}

	ans := make([]byte, n)
	m1, m2 := 0, 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			ans[i] = max[m1]
			m1++
		} else {
			ans[i] = min[m2]
			m2++
		}
	}

	return string(ans)
}

// @lc code=end
