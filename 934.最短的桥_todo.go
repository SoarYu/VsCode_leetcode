package main

func shortestBridge(grid [][]int) (ans int) {
	type pair struct{ i, j int }
	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n := len(grid)
	pairs := []pair{} //记录第一个岛的所有陆地的坐标

	//  dfs
	var dfs func(int, int)
	dfs = func(x, y int) {
		grid[x][y] = 2
		pairs = append(pairs, pair{x, y})
		for _, dir := range dirs {
			a, b := x+dir.i, y+dir.j
			if a >= 0 && b >= 0 && a < n && b < n && grid[a][b] == 1 {
				dfs(a, b)
			}
		}
	}

	// dfs将第一岛的陆地 1 标记为 2
	flag := 1
	for i, row := range grid {
		if flag == 0 {
			break
		}
		for j, k := range row {
			if k == 1 {
				dfs(i, j)
				flag = 0
				break
			}
			// k = 1

		}
	}

	// bfs 扩展第二个岛
	for {
		// 从第一岛边缘坐标遍历做bfs前往第二个岛
		tmp := pairs
		pairs = nil
		for _, p := range tmp {
			for _, d := range dirs {
				x, y := p.i+d.i, p.j+d.j
				if x >= 0 && x < n && y >= 0 && y < n {
					if grid[x][y] == 1 {
						return
					}
					if grid[x][y] == 0 {
						grid[x][y] = 2
						pairs = append(pairs, pair{x, y})
					}
				}
			}
		}
		ans++
	}
}
