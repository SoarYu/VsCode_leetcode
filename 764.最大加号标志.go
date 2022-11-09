/*
 * @lc app=leetcode.cn id=764 lang=golang
 *
 * [764] 最大加号标志
 */
package main


// @lc code=start
func orderOfLargestPlusSign(n int, mines [][]int) int {
	grid := make([][]int, n)
	for i:=0;i<n;i++ {
		grid[i] = make([]int, n)
	}
	for _, mine := range mines {
		grid[mine[0]][mine[1]] = 1
	}
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		width := 1
		for {
			flag := true
			for _, dir := range dirs {
				i, j := x+dir[0]*width, y+dir[1]*width
				if i >= 0 && j >= 0 && i < n && j < n && grid[i][j]!=1 {
					continue
				} else {
					flag = false
					break
				}
			}
			if flag {
				width++
			} else {
				break
			}
		}
		return width
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j]!=1 {
                // 检查是否有潜在最大值
                if i-res>=0 && i+res<n && j-res>=0 && j+res<n {
                    if tmp := dfs(i, j); tmp > res {
                        res = tmp
                    }
                }
			}
		}
	}
	return res
}

// @lc code=end
