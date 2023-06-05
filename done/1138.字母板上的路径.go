/*
 * @lc app=leetcode.cn id=1138 lang=golang
 *
 * [1138] 字母板上的路径
 */
package daily

// @lc code=start
func alphabetBoardPath(target string) string {
	board := make(map[byte][]int)
	m, n := 6, 5

	for i, cnt := 0, 0; i < m-1; i++ {
		for j := 0; j < n; j++ {
			board['a'+byte(cnt)] = []int{i, j}
			cnt++
		}
	}
	board['z'] = []int{5, 0}

	ans := ""
	pre := byte('a')
	for i := range target {
		if pre == target[i] {
			ans += "!"
			continue
		}

		tmp := target[i]
		if target[i] == 'z' {
			tmp = 'u'
		}

		tmpPre := pre
		if pre == 'z' {
			tmpPre = 'u'
			ans += "U"
		}

		y1, x1 := board[tmpPre][0], board[tmpPre][1]
		y2, x2 := board[tmp][0], board[tmp][1]
		ud, lr := y2-y1, x2-x1

		for lr != 0 {
			// lr > 0 x2 > x1 向右移动
			if lr > 0 {
				ans += "R"
				lr--
			} else { // lr < 0 x2 < x1 向左移动
				ans += "L"
				lr++
			}
		}

		for ud != 0 {
			// ud > 0 y2 > y1 向下移动
			if ud > 0 {
				ans += "D"
				ud--
			} else { // ud < 0 y2 < y1 向上移动
				ans += "U"
				ud++
			}
		}

		if target[i] == 'z' {
			ans += "D"
		}

		ans += "!"
		pre = target[i]
	}
	return ans
}

// @lc code=end
