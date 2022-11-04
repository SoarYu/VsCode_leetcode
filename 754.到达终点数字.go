/*
 * @lc app=leetcode.cn id=754 lang=golang
 *
 * [754] 到达终点数字
 */

// @lc code=start
func reachNumber(target int) int {
	func reachNumber(target int) int {
		if target < 0 {
			target = -target
		}
		i, sum := 0, 0
		for sum<target{
			sum += i
			i++
		}
		if sum == target || (sum - target) % 2 == 0  {
			return i-1
		}
		if (sum + i - target) % 2 == 0 {
			return i
		}
		return i+1
	}
	/*	
	作者：xiaoY
	链接：https://leetcode.cn/problems/reach-a-number/solution/by-xiaoy-iura/
	来源：力扣（LeetCode）
	著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
	*/
}
// @lc code=end

