/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package daily

// 输入头节点，翻转以头节点开始的k个节点，返回翻转后的头节点
func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	for i := 0; i < k; i++ {
		if cur == nil { // 节点数不够k个 
			//如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序
			return head
		}
		cur = cur.Next
	}
	// head当前组的头节点, cur下一组的头节点
	newHead := reverse(head, cur)     // 返回翻转后的头节点newHead
	head.Next = reverseKGroup(cur, k) // 翻转后head为尾节点,next指向下一组翻转后的头节点
	return newHead
}

// 输入头节点和下一组的头节点，返回翻转后的头节点
func reverse(start, end *ListNode) *ListNode {
	head := &ListNode{}
	cur := start
	for cur != end {
		tmp := cur.Next
		cur.Next = head.Next
		head.Next = cur
		cur = tmp
	}
	return head.Next
}

// @lc code=end
