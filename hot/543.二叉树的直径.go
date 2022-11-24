package hot

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 路径长度 = 左子树高度 加 右子树高度
var res int

func diameterOfBinaryTree(root *TreeNode) int {
	depth(root)
	return res
}
func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := depth(root.Left), depth(root.Right)
	res = max(l+r, res)
	return max(l, r) + 1
}
func diameterOfBinaryTree1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// getH(r.L)
	l, r := getHeight(root.Left), getHeight(root.Right)
	// getH(r.R)
	return max(l+r+1, max(diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right)))
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
