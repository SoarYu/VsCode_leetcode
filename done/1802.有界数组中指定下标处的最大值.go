/*
 * @lc app=leetcode.cn id=1802 lang=golang
 *
 * [1802] 有界数组中指定下标处的最大值
 */
package daily

// @lc code=start
func maxValue(n int, index int, maxSum int) int {
    left, right := 1, maxSum
    ans := 1
    var f func(int, int) int 
    f = func(begin, length int) int {
        if length == 0 {
            return 0
        }
        if begin > length {
            return length * (2*begin - 1 - length) / 2
        }
        // 前(begin - 1 ) 3 项 + length - (begin - 1)  1项
        return f(begin, begin-1) + length - begin + 1
    }

    l1, l2 := index, n - index - 1

    for left <= right{
        mid := left + (right - left) / 2
        // 检查mid是否符合
        if tmp := mid + f(mid, l1) + f(mid, l2); tmp <= maxSum {
            // 加大mid
            if mid > ans {
                ans = mid
            }
            left = mid + 1
        } else {    // tmp > maxSum
            // 减小mid
            right = mid - 1
        }
    }
    return ans
}

// @lc code=end
