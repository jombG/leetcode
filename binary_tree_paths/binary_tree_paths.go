package binary_tree_paths

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	var dfs func(root *TreeNode, str string)

	if root.Left == nil && root.Right == nil {
		return []string{fmt.Sprintf("%d", root.Val)}
	}

	dfs = func(node *TreeNode, str string) {
		if node.Right == nil && node.Left == nil {
			res = append(res, fmt.Sprintf("%s->%d", str, node.Val))
			return
		}

		if str == "" {
			if node.Right != nil {
				dfs(node.Right, fmt.Sprintf("%d", node.Val))
			}

			if node.Left != nil {
				dfs(node.Left, fmt.Sprintf("%d", node.Val))
			}

		} else {

			if node.Right != nil {
				dfs(node.Right, fmt.Sprintf("%s->%d", str, node.Val))
			}

			if node.Left != nil {
				dfs(node.Left, fmt.Sprintf("%s->%d", str, node.Val))
			}
		}

	}
	dfs(root, "")

	return res
}
