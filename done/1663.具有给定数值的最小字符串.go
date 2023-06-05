/*
 * @lc app=leetcode.cn id=1663 lang=golang
 *
 * [1663] 具有给定数值的最小字符串
 */
package daily
// @lc code=start
func getSmallestString(n, k int) string {
    s := []byte{}
    for i := 1; i <= n; i++ {
        lower := max(1, k-(n-i)*26)
        k -= lower
        s = append(s, 'a'+byte(lower)-1)
    }
    return string(s)
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}
// @lc code=end

