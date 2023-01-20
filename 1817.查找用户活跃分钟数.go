/*
 * @lc app=leetcode.cn id=1817 lang=golang
 *
 * [1817] 查找用户活跃分钟数
 */
package daily

// @lc code=start
func findingUsersActiveMinutes(logs [][]int, k int) []int {
	ans := make([]int, k)
	cnt := make(map[int]map[int]bool)
	for _, user := range logs {
		id, opsTime := user[0], user[1]
		if cnt[id] == nil {
			cnt[id] = make(map[int]bool)
		}
		cnt[id][opsTime] = true
	}
	for _, m := range cnt {
		ans[len(m)-1]++
	}

	return ans
}

// @lc code=end
