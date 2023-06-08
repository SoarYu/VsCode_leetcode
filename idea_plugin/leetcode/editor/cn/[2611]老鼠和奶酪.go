// 有两只老鼠和 n 块不同类型的奶酪，每块奶酪都只能被其中一只老鼠吃掉。
// 下标为 i 处的奶酪被吃掉的得分为：
// 如果第一只老鼠吃掉，则得分为 reward1[i] 。
// 如果第二只老鼠吃掉，则得分为 reward2[i] 。
//
// 给你一个正整数数组 reward1 ，一个正整数数组 reward2 ，和 一个非负整数 k 。
// 请你返回第一只老鼠恰好吃掉 k 块奶酪的情况下，最大 得分为多少。
//
// 示例 1：
// 输入：reward1 = [1,1,3,4], reward2 = [4,4,1,1], k = 2
// 输出：15
// 解释：这个例子中，第一只老鼠吃掉第 2 和 3 块奶酪（下标从 0 开始），第二只老鼠吃掉第 0 和 1 块奶酪。
// 总得分为 4 + 4 + 3 + 4 = 15 。
// 15 是最高得分。
//
// 示例 2：
// 输入：reward1 = [1,1], reward2 = [1,1], k = 2
// 输出：2
// 解释：这个例子中，第一只老鼠吃掉第 0 和 1 块奶酪（下标从 0 开始），第二只老鼠不吃任何奶酪。
// 总得分为 1 + 1 = 2 。
// 2 是最高得分。
//
// 提示：
// 1 <= n == reward1.length == reward2.length <= 10⁵
// 1 <= reward1[i], reward2[i] <= 1000
// 0 <= k <= n
// Related Topics 贪心 数组 排序 堆（优先队列） 👍 78 👎 0
package main

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
func miceAndCheese(r1 []int, r2 []int, k int) int {
	type diff struct {
		index, diff int
	}
	// 下标为 i 处的奶酪被吃掉的得分为：r1, r2
	// 第一只老鼠恰好吃掉 k 块奶酪的情况下，最大 得分为多少。
	n := len(r1)
	// diff[i] = r1[i] - r2[i]  diff[i] > 0 r1[i] > r2[i]
	rMap := make(map[int]bool)
	dSlice := make([]diff, n)
	var res int
	for i := range r1 {
		if r1[i] > r2[i] {
			res += r1[i]
			rMap[i] = true
		} else {
			res += r2[i]
		}
		dSlice[i] = diff{index: i, diff: r1[i] - r2[i]}
	}
	// r1: x    r2: n-x  有x个 r1[i] > r2[i] 即diff[i] > 0
	x := len(rMap)
	// x = k return
	if x == k {
		return res
	}
	sort.Slice(dSlice, func(i, j int) bool {
		return abs(dSlice[i].diff) < abs(dSlice[j].diff)
	})
	var cnt int
	if x > k {
		// x > k  从r1中删除(x-k)个, 使总分最大，选择x范围内(x-k)个 r1[i]-r2[i]>0 最接近0的
		for i := range dSlice {
			if rMap[dSlice[i].index] && dSlice[i].diff >= 0 {
				cnt++
				res -= dSlice[i].diff
				if cnt >= x-k {
					return res
				}
			}
		}
	} else {
		// x < k  从r2中补充(k-x)个，使总分最大，选择x范围外 (k-x) r1[i]-r2[i]<0 选择最接近0的
		for i := range dSlice {
			if !rMap[dSlice[i].index] && dSlice[i].diff <= 0 {
				cnt++
				res += dSlice[i].diff
				if cnt >= k-x {
					return res
				}
			}
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//leetcode submit region end(Prohibit modification and deletion)
