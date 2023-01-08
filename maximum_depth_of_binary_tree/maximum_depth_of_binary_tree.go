package maximum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	mDepth := -1

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			mDepth = max(mDepth, depth)
			return
		}

		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}

	dfs(root, 0)

	return mDepth
}
