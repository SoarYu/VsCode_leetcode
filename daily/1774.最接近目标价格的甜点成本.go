/*
 * @lc app=leetcode.cn id=1774 lang=golang
 *
 * [1774] 最接近目标价格的甜点成本
 */
package daily

// dfs
// @lc code=start
func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	m := len(toppingCosts)
	res := baseCosts[0]
	var dfs func(cur, index int)
	dfs = func(cur, index int) {
		// fmt.Println(cur)
		if t1, t2 := abs(cur-target), abs(res-target); t1 < t2 {
			res = cur
		} else if t1 == t2 && cur < res {
			res = cur
		} else if cur > target {
			return
		}

		if index == m {
			return
		}

		dfs(cur, index+1)
		dfs(cur+toppingCosts[index], index+1)
		dfs(cur+toppingCosts[index]*2, index+1)
	}
	// target = x + y1 + ... + yn
	for i := range baseCosts {
		dfs(baseCosts[i], 0)
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end
