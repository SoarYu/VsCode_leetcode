/*
 * @lc app=leetcode.cn id=994 lang=golang
 *
 * [994] 腐烂的橘子
 */
package daily

// @lc code=startfunc orangesRotting(grid [][]int) int {
	// bfs queue
	m, n := len(grid), len(grid[0])
	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var fresh int
	queue := [][]int{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				fresh++
			} else if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	var time int
	for fresh > 0 && len(queue) > 0 {
		newQueue := [][]int{}
		// 腐烂队列
		for _, node := range queue {
			// x, y := node[0], node[1]
			for _, d := range dir {
				x, y := node[0]+d[0], node[1]+d[1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					grid[x][y] = 2
					newQueue = append(newQueue, []int{x, y})
					fresh--
				}
			}
		}
		queue = newQueue
		time++
	}
    if fresh > 0 {
        return -1
    }
	return time
}

// @lc code=end
