/*
 * @lc app=leetcode.cn id=939 lang=golang
 *
 * [939] 最小面积矩形
 */
package daily

// @lc code=start
func minAreaRect(points [][]int) int {
	if len(points) < 4 {
		return 0
	}
	// 矩形的边平行于 x 轴和 y 轴
	// 找出 某个点 的 x,y 查找与它相同 x 或 y 的点， 三点确定后查找是否有对应的第四个点 计算面积
	getXbyY, getYbyX := make(map[int][]int), make(map[int][]int)
	pointMap := make(map[[2]int]bool)

	for _, p := range points {
		x, y := p[0], p[1]
		pointMap[[2]int{x, y}] = true
		if getYbyX[x] == nil {
			getYbyX[x] = make([]int, 0)
		}
		if getXbyY[y] == nil {
			getXbyY[y] = make([]int, 0)
		}
		getYbyX[x] = append(getYbyX[x], y)
		getXbyY[y] = append(getXbyY[y], x)
	}

	var ans int

	for _, p := range points {
		x1, y1 := p[0], p[1]
		// y 相同 getXbyY
		for _, x2 := range getXbyY[y1] {
			if x2 <= x1 {
				continue
			}
			xlength := x2 - x1
			// x 相同 getYbyX
			for _, y2 := range getYbyX[x1] {
				if y2 <= y1 {
					continue
				}
				ylength := y2 - y1
				// 是否存在 x2 y2 是矩阵成立
				if pointMap[[2]int{x2, y2}] {
					if ans == 0 || ans > xlength*ylength {
						ans = xlength * ylength
					}
				}
			}

		}
	}
	return ans
}

// @lc code=end
