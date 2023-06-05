/*
 * @lc app=leetcode.cn id=938 lang=golang
 *
 * [938] 二叉搜索树的范围和
 */
package daily

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
	// 中序遍历 非递归
	queue := []*TreeNode{}
	node := root
	var ans int
	for len(queue) > 0 || node != nil {
		for node != nil {
			queue = append(queue, node)
			node = node.Left
		}

		node = queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if node.Val >= low && node.Val <= high {
			ans += node.Val
		} else if node.Val > high {
			break
		}

		node = node.Right
	}
	return ans
}

// @lc code=end
