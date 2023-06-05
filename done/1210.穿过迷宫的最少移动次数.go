/*
 * @lc app=leetcode.cn id=1210 lang=golang
 *
 * [1210] 穿过迷宫的最少移动次数
 */
package daily

// @lc code=start
func minimumMoves1210(grid [][]int) int {
	n := len(grid)
	ans := -1
	type index struct {
		x, y, status int
	}
	isVisited := make(map[index]int)
	// dfs
	var dfs func(headx, heady, status, level int)
	dfs = func(headx, heady, status, level int) {
		if headx == n-1 && heady == n-1 && status == 0 {
			if ans < 0 || level < ans {
				ans = level
			}
			return
		}

		in := index{headx, heady, status}
		if val, hasVal := isVisited[in]; hasVal && level >= val {
			return
		}
		isVisited[in] = level

		if status == 0 {
			// right
			if headx+1 < n && grid[heady][headx+1] != 1 {
				dfs(headx+1, heady, status, level+1)
			}
			if heady+1 < n && grid[heady+1][headx] != 1 && grid[heady+1][headx-1] != 1 {
				// down
				dfs(headx, heady+1, status, level+1)
				// turndown
				dfs(headx-1, heady+1, 1, level+1)
			}

		} else {
			// down
			if heady+1 < n && grid[heady+1][headx] != 1 {
				dfs(headx, heady+1, status, level+1)
			}
			if headx+1 < n && grid[heady][headx+1] != 1 && grid[heady-1][headx+1] != 1 {
				// right
				dfs(headx+1, heady, status, level+1)
				// turnright
				dfs(headx+1, heady-1, 0, level+1)
			}
		}
	}

	dfs(1, 0, 0, 0)
	return ans
}

// @lc code=end
