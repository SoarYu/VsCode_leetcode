/*
 * @lc app=leetcode.cn id=2032 lang=golang
 *
 * [2032] 至少在两个数组中出现的值
 */
package daily
// @lc code=start
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
    m := make(map[int]int)
    for _, num := range nums1 {
        m[num] = 1
    }
    for _, num := range nums2 {
        if v, hasV := m[num]; !hasV {
            m[num] = 2
        } else if v==1 {
            m[num] = 3
        }
    }
    for _, num := range nums3 {
        if v, hasV := m[num]; hasV {
            if v == 1 || v==2 || v==3 {
                m[num] += 4
            }
        }
    }
    ans := []int{}
    for k, v := range m {
        if v == 3 || v == 5 || v == 6 || v==7{
            ans = append(ans, k)
        }
    }
    return ans
}
// @lc code=end

