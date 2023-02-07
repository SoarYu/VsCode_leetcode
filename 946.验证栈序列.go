/*
 * @lc app=leetcode.cn id=946 lang=golang
 *
 * [946] 验证栈序列
 */
package daily

// @lc code=start
func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	x := 0
	for i := range popped {
		// 存在栈顶 出栈
		if len(stack) > 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			continue
		}
		// 栈为空 或 栈顶元素不为popped[i]
		for x < len(pushed) && pushed[x] != popped[i] {
			stack = append(stack, pushed[x])
			x++
		}
		if x == len(pushed) {
			return false
		}
		x++
	}
	return true
}

// @lc code=end
