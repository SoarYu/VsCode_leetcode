/*
 * @lc app=leetcode.cn id=1824 lang=golang
 *
 * [1824] 最少侧跳次数
 */
package daily

// @lc code=start
func minSideJumps(obstacles []int) int {
	var ans int
	n, cur, index := len(obstacles), 2, 0
	for index < n-1 {
		// 检查前面是否障碍
		if obstacles[index+1] == cur {
			// 测跳加一, 有两条路 选择最长的一条
			// cur = 2 next=1, 3
			var n1, n2 int
			switch cur {
			case 1:
				n1, n2 = 2, 3
			case 2:
				n1, n2 = 1, 3
			case 3:
				n1, n2 = 1, 2
			}
			var lf1, lf2 int
			for j := index; j < n && obstacles[j] != n1; j++ {
				lf1++
			}
			for j := index; j < n && obstacles[j] != n2; j++ {
				lf2++
			}
			if lf1 >= lf2 {
				cur = n1
				index += lf1 - 1
			} else {
				cur = n2
				index += lf2 - 1
			}
			ans++
		} else {
			index++
		}
	}
	return ans
}

// @lc code=end
