/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */
package daily
// @lc code=start
func combine(n int, k int) [][]int {
    ans := [][]int{}
    var dfs func([]int, int)
    dfs = func(arr []int, index int) {
        if len(arr) >= k {
            t := make([]int, k)
            copy(t, arr)
            ans = append(ans, t)
            return
        }
        for i := index+1; i<n+1; i++ {
            arr = append(arr, i)
            dfs(arr, i)
            arr = arr[:len(arr)-1]
        }

    }
    dfs([]int{}, 0)
    return ans
}
// @lc code=end

