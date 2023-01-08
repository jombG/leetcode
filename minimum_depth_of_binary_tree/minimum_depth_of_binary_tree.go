package minimum_depth_of_binary_tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	minLength := math.MaxInt64

	min := func(a, b int) int {
		if a > b {
			return b
		}

		return a
	}

	var dfs func(node *TreeNode, counter int)

	dfs = func(node *TreeNode, counter int) {
		if node.Right == nil && node.Left == nil {
			minLength = min(minLength, counter)
		}

		if node.Left == nil && node.Right != nil {
			dfs(node.Right, counter+1)
		}

		if node.Left != nil && node.Right == nil {
			dfs(node.Left, counter+1)
		}
		if node.Left != nil && node.Right != nil {
			dfs(node.Left, counter+1)
			dfs(node.Right, counter+1)
		}
	}

	if root == nil {
		return 0
	}

	dfs(root, 1)

	return minLength
}
