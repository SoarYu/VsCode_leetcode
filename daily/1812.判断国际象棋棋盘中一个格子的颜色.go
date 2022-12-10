/*
 * @lc app=leetcode.cn id=1812 lang=golang
 *
 * [1812] 判断国际象棋棋盘中一个格子的颜色
 */
package daily

// @lc code=start
func squareIsWhite(coordinates string) bool {
	x, y := coordinates[0]-'a', coordinates[1]-'0'
	/*
		x为奇数：1. y为奇数	2. y为偶数  1.true  2.false
		x为偶数：1. y为奇数	2. y为偶数 	1.false 2.true
	*/
	return x%2 == 0 && y%2 == 0 || x%2 == 1 && y%2 == 1
}

// @lc code=end
