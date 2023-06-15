// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
// 示例 1:
//输入: [3,2,1,5,6,4], k = 2
//输出: 5
// 
//
// 示例 2: 
//输入: [3,2,3,1,2,4,5,5,6], k = 4
//输出: 4 
//
// 
//
// 提示： 
// 1 <= k <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴ 
// 
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 2197 👎 0
package main
import "fmt"
//leetcode submit region begin(Prohibit modification and deletion)

// 快排 nlogn  堆排 nlogn
func findKthLargest(nums []int, k int) int {
	var partitation func(int, int) int
	final := len(nums) - k
	partitation = func(left, right int) int {
		index := left
		for left <= right {
			for left <= right && nums[index] <= nums[right] {
				right--
			}
			// swap index, right
			if left <= right {
				nums[index], nums[right] = nums[right], nums[index]
				index = right
				right--
			}

			for left <= right && nums[index] >= nums[left] {
				left++
			}
			// swap index, left
			if left <= right {
				nums[index], nums[left] = nums[left], nums[index]
				index = left
				left++
			}
		}
		return index
	}
	left, right := 0, len(nums)-1
	// fmt.Println(nums, k)
	for left < right {
		// temp 确定了nums[temp]排序后最终位置
		temp := partitation(left, right)
		// fmt.Printf("%d\t%d\n", nums, temp)

		if temp == final {
			return nums[temp]
		} else if temp > final { // 往左边找k，temp无效,right有效 [left, temp-1]
			right = temp-1
		} else if temp < final { // 往右边找k，temp无效,left有效 [temp+1, right]
			left = temp+1
		}
	}
	return nums[left]
}
//leetcode submit region end(Prohibit modification and deletion)
