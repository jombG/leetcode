package validate_binary_search_tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isBST(root, math.MinInt64, math.MaxInt64)
}

func isBST(root *TreeNode, rangeMin, rangeMax int) bool {
	if root == nil {
		return true
	}

	if root.Val < rangeMin || root.Val > rangeMax {
		return false
	}

	return isBST(root.Left, rangeMin, root.Val-1) && isBST(root.Right, root.Val+1, rangeMax)
}
