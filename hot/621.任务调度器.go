/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */
package hot

// @lc code=start
func leastInterval(tasks []byte, n int) int {
	var arr [26]int
	var lastFinish [26]int
	time, all := 0, len(tasks)
	for _, task := range tasks {
		arr[task-'A']++
		lastFinish[task-'A'] = -n - 1
	}
    var next func(int) int 
    next = func(time int) int {
        todo := -1
        for i := range arr {
            if arr[i] > 0 && time - lastFinish[i] > n {
                if todo < 0 {
                    todo = i
                } else if arr[i] > arr[todo] {
                    todo = i
                }
            }
        }
        return todo
    }
	for all > 0 {
        if todo := next(time); todo != -1 {
            arr[todo]--
            lastFinish[todo] = time
            all--
        }
        time++
	}
	return time
}

// @lc code=end
