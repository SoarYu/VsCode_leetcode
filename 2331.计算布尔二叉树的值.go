/*
 * @lc app=leetcode.cn id=2331 lang=golang
 *
 * [2331] 计算布尔二叉树的值
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package daily

func evaluateTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == 1
	}
	if root.Val == 2 { // or
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	} else { // and
		return evaluateTree(root.Left) && evaluateTree(root.Right)
	}
}

// @lc code=end
