package balanced_binary_tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var dfs func(node *TreeNode) bool

	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		l, r := height(node.Left), height(node.Right)

		if math.Abs(float64(l-r)) > 1 {
			return false
		}

		return dfs(node.Left) && dfs(node.Right)
	}

	return dfs(root)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	return max(height(node.Left), height(node.Right)) + 1
}
