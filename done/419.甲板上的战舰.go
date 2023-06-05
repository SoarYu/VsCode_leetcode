/*
 * @lc app=leetcode.cn id=419 lang=golang
 *
 * [419] 甲板上的战舰
 */
package daily

// @lc code=start
func countBattleships(board [][]byte) int {
	var ans int
	m, n := len(board), len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				board[i][j] = 'O'
				ans++
				// 水平
				for t := j + 1; t < n; t++ {
					if board[i][t] == 'X' {
						board[i][t] = 'O'
					} else {
						break
					}
				}

				// 竖直
				for t := i + 1; t < m; t++ {
					if board[t][j] == 'X' {
						board[t][j] = 'O'
					} else {
						break
					}
				}
			}
		}
	}
	return ans
}

// @lc code=end
