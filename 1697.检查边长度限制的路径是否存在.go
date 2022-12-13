/*
 * @lc app=leetcode.cn id=1697 lang=golang
 *
 * [1697] 检查边长度限制的路径是否存在
 */
package daily

import "sort"

// @lc code=start
func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	sort.Slice(edgeList, func(i, j int) bool { return edgeList[i][2] < edgeList[j][2] })

	// 并查集模板
	// 连通分量fa
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	// 寻找 x 的连通分量所属
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	// 将两个连通分量连接
	merge := func(from, to int) {
		fa[find(from)] = find(to)
	}

	// 记录返回结果的位置顺序
	for i := range queries {
		queries[i] = append(queries[i], i)
	}

	// 按照 limit 从小到大排序，方便离线
	sort.Slice(queries, func(i, j int) bool { return queries[i][2] < queries[j][2] })

	ans := make([]bool, len(queries))
	k := 0

	for _, q := range queries {
		// 边长度小于limit值， 连通from，to的连通分量
		for ; k < len(edgeList) && edgeList[k][2] < q[2]; k++ {
			merge(edgeList[k][0], edgeList[k][1])
		}

		// 检查是否属于同一个连通分量，即是否存在路径
		ans[q[3]] = find(q[0]) == find(q[1])
	}
	return ans
}

// @lc code=end
