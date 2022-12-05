/*
 * @lc app=leetcode.cn id=1687 lang=golang
 *
 * [1687] 从仓库到码头运输箱子
 */
package daily

// @lc code=start
func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	n := len(boxes)
	p := make([]int, n+1)
	w := make([]int, n+1)
	neg := make([]int, n+1)
	W := make([]int, n+1)
	for i := 1; i <= n; i++ {
		p[i] = boxes[i-1][0]
		w[i] = boxes[i-1][1]
		if i > 1 {
			neg[i] = neg[i-1]
			if p[i-1] != p[i] {
				neg[i]++
			}
		}
		W[i] = W[i-1] + w[i]
	}

	opt := []int{0}
	f := make([]int, n+1)
	g := make([]int, n+1)

	for i := 1; i <= n; i++ {
		for i-opt[0] > maxBoxes || W[i]-W[opt[0]] > maxWeight {
			opt = opt[1:]
		}

		f[i] = g[opt[0]] + neg[i] + 2

		if i != n {
			g[i] = f[i] - neg[i+1]
			for len(opt) > 0 && g[i] <= g[opt[len(opt)-1]] {
				opt = opt[:len(opt)-1]
			}
			opt = append(opt, i)
		}
	}

	return f[n]
}

// @lc code=end
