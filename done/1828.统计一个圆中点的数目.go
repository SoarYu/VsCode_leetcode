/*
 * @lc app=leetcode.cn id=1828 lang=golang
 *
 * [1828] 统计一个圆中点的数目
 */
package daily
// @lc code=start
func countPoints(points [][]int, queries [][]int) []int {
    ans := make([]int, len(queries))
    for _, p := range points {
        px, py := p[0], p[1]
        for k, q := range queries {
            qx, qy, qr := q[0], q[1], q[2]
            // 这个 点p 是否存在 园k 内
            if (qx-px)*(qx-px) + (qy-py)*(qy-py) <= qr * qr {
                ans[k]++
            }
        }
    }
    return ans
}
// @lc code=end

