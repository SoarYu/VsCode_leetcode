/*
 * @lc app=leetcode.cn id=1769 lang=golang
 *
 * [1769] 移动所有球到每个盒子所需的最小操作数
 */
package daily

// @lc code=start
func minOperations(boxes string) []int {
	res := make([]int, len(boxes))
	for i, b := range boxes {
		if b == '1' {
			for j := range res {
				res[j] += abs(i - j)
			}
		}
	}
	return res
}

func minOperations1(boxes string) []int {
	res := make([]int, len(boxes))
	left, right := int(boxes[0]-'0'), 0
	for i := 1; i < len(boxes); i++ {
		if boxes[i] == '1' {
			res[0] += i
			right++
		}
	}

	for i := 1; i < len(boxes); i++ {
		res[i] = res[i-1] + left - right
		if boxes[i] == '1' {
			left++
			right--
		}
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
