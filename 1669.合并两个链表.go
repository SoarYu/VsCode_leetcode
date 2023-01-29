/*
 * @lc app=leetcode.cn id=1669 lang=golang
 *
 * [1669] 合并两个链表
 */
package daily
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    // 快慢指针
    // 1 <= a < b < n - 1
    count, fast, slow := 0, list1, list1 // 0, 0
    for count <= b {
        if count < a-1 {
            slow = slow.Next
        }
        fast = fast.Next
        count++
    }
    slow.Next = list2
    node := list2
    for node.Next!=nil{
        node = node.Next
    }
    node.Next = fast
    return list1
}
// @lc code=end

