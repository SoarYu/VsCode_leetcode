/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */
package hot

import (
	"sort"
)

// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0]
	})

	res := make([][]int, len(people))
	for i := 0; i < len(people); i++ {
		res[i] = make([]int, 2)
	}

	return res
}

// @lc code=end
