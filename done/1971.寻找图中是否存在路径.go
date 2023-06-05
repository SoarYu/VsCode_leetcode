/*
 * @lc app=leetcode.cn id=1971 lang=golang
 *
 * [1971] 寻找图中是否存在路径
 */
package daily

// @lc code=start
func validPath(n int, edges [][]int, source int, destination int) bool {
	// 连通分量
	f := make([]int, n)
	for i := range f {
		f[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if f[x] != x {
			f[x] = find(f[x])
		}
		return f[x]
	}

	for _, edge := range edges {
		f[find(edge[0])] = find(edge[1])
	}
	return find(source) == find(destination)
}

// @lc code=end
