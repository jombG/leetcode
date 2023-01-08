package path_sum_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	results := make([][]int, 0)

	var dfs func(node *TreeNode, arr []int, sum int)

	dfs = func(node *TreeNode, arr []int, sum int) {
		if node == nil {
			return
		}

		arr = append(arr, node.Val)
		if node.Val == sum && node.Left == nil && node.Right == nil {
			results = append(results, arr)
			return
		}

		dfs(node.Left, append([]int{}, arr...), sum-node.Val)
		dfs(node.Right, append([]int{}, arr...), sum-node.Val)

	}

	dfs(root, []int{}, targetSum)

	return results
}
