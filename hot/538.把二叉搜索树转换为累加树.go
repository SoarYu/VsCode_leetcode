/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
 */
package hot

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//每个节点 node 的新值等于 原树中大于或等于 node.val 的值之和。
func convertBST(root *TreeNode) *TreeNode {
	// 0 1 2 3 4 5 6 7 8
	// 8 7 6 5 4 3 2 1 0	
	sum := 0
	var deInOrder func(node *TreeNode)
	deInOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		deInOrder(node.Right)
		sum += node.Val
		node.Val = sum
		deInOrder(node.Left)
	}

	return root
}


// @lc code=end
