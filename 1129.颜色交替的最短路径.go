/*
 * @lc app=leetcode.cn id=1129 lang=golang
 *
 * [1129] 颜色交替的最短路径
 */
package daily
// @lc code=start
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	// 路径类型0，1 表示0到x的存在最后为红或蓝色的路径
	type Pair struct {
		index, color int
	}
	// i 到 j 存在路径， 颜色为color
	pairs := make([][]Pair, n)
	
	for _, edge := range redEdges {
		from, to := edge[0], edge[1]
		pairs[from] = append(pairs[from], Pair{index: to, color: 0})
	}
	for _, edge := range blueEdges {
		from, to := edge[0], edge[1]
		pairs[from] = append(pairs[from], Pair{index:to, color: 1})
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	
	queue := []Pair{Pair{0,0}, Pair{0,1}}
	
	isVisited := make([][2]bool, n)
	
	for length:=0; len(queue)>0; length++{
		tmp := queue
		queue = []Pair{}
		
		for _, p := range tmp {
			// p[0]to, p[1]color
			to := p.index
			// 判断长度
			if ans[to] < 0 {
				ans[to] = length
			}
			for _, edge := range pairs[to] {
				if edge.color != p.color && !isVisited[edge.index][edge.color] {
					isVisited[edge.index][edge.color] = true
					queue = append(queue, edge)
				}
			}
		}
	}
	return ans
}
// @lc code=end

