//给定 pushed 和 popped 两个序列，每个序列中的 值都不重复，只有当它们可能是在最初空栈上进行的推入 push 和弹出 pop 操作序列的结果时
//，返回 true；否则，返回 false 。 
//
// 示例 1： 
//输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
//输出：true
//解释：我们可以按以下顺序执行：
//push(1), push(2), push(3), push(4), pop() -> 4,
//push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
//
// 示例 2： 
//输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
//输出：false
//解释：1 不能在 2 之前弹出。
// 
// 提示：
// 1 <= pushed.length <= 1000
// 0 <= pushed[i] <= 1000 
// pushed 的所有元素 互不相同 
// popped.length == pushed.length 
// popped 是 pushed 的一个排列 
// Related Topics 栈 数组 模拟 👍 387 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func validateStackSequences(pushed []int, popped []int) bool {
	// 模拟
	stack := []int{}
	// pop
	index := 0
	for _, v := range popped {
		// 栈顶
		if len(stack) > 0 && stack[len(stack)-1] == v {
			stack = stack[:len(stack)-1]
			continue
		}
		// 入栈
		for index < len(pushed) && pushed[index] != v {
			stack = append(stack, pushed[index])
			index++
		}
		if index < len(pushed) && pushed[index] == v {
			index++
		} else {
			return false
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
