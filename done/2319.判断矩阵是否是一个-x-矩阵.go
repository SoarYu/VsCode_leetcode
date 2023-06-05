/*
 * @lc app=leetcode.cn id=2319 lang=golang
 *
 * [2319] 判断矩阵是否是一个 X 矩阵
 */
package daily

// @lc code=start
func checkXMatrix(grid [][]int) bool {
	n := len(grid)-1
	for i := 0; i < len(grid); i++ {
		row := grid[i]
		for j := range row {
			if j==i || j==n-i {
                if row[j] == 0 {
                    return false
                }
			} else if row[j] != 0 {
				return false
			}
		}
	}
	return true
}

// @lc code=end
