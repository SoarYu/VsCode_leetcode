/*
 * @lc app=leetcode.cn id=1145 lang=golang
 *
 * [1145] 二叉树着色游戏
 */
package daily

// @lc code=start
//  * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	var outNum, xleftNum, xrightNum int
	var count func(root *TreeNode) int
	count = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := count(node.Left), count(node.Right)

		if node.Val == x {
			xleftNum, xrightNum = l, r
		}

		return 1 + l + r
	}
	outNum = count(root) - 1 - xleftNum - xrightNum

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 是否选择其子树

	// 1. 选择 out 不能再去leftorright out
	// 2. 选择left or right, 不能选择out和另外一个 left
	return max(outNum, max(xleftNum, xrightNum)) > n/2
}

/*
[1,2,3,4,5]
5

	1
   2 3
  4 5

  out left right 2 1 1
*/
// @lc code=end
