package path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	pathSumExists := false
	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}

		if node.Right == nil && node.Left == nil && sum+node.Val == targetSum {
			pathSumExists = true
		}

		dfs(node.Right, sum+node.Val)
		dfs(node.Left, sum+node.Val)
	}

	dfs(root, 0)

	return pathSumExists
}
