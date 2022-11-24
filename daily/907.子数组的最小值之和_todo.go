package daily

// @lc code=start
func sumSubarrayMins_todo(arr []int) int {
	n := len(arr)
	// 左边界数组 边界内都是小于等于i的值
	left := make([]int, n)
	//单调栈
	stack := []int{-1}
	// 从左到右	 left[i] 为左侧严格小于 arr[i] 的最近元素位置（不存在时为 -1）
	for i, x := range arr {
		// 栈顶大于x
		for len(stack) > 1 && arr[stack[len(stack)-1]] >= x {
			stack = stack[:len(stack)-1] //
		}
		left[i] = stack[len(stack)-1]
		stack = append(stack, i)
	}

	// 右边界 right[i] 为右侧小于等于 arr[i] 的最近元素位置（不存在时为 n）
	right := make([]int, n)
	stack = []int{n} // 方便赋值 right
	for i := n - 1; i >= 0; i-- {
		//  arr=[1,2,4,2,3,1]
		for len(stack) > 1 && arr[stack[len(stack)-1]] > arr[i] {
			stack = stack[:len(stack)-1]
		}
		right[i] = stack[len(stack)-1]
		stack = append(stack, i)
	}

	ans := 0
	for i, x := range arr {
		ans += x * (i - left[i]) * (right[i] - i)
	}
	return ans % (1e9 + 7)
}

// @lc code=end
