package subtree_of_another_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if same(root, subRoot) {
		return true
	}

	if root == nil {
		return false
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func same(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil {
		return node2 == nil
	}

	if node2 == nil {
		return false
	}

	if node2.Val != node1.Val {
		return false
	}

	return same(node1.Left, node2.Left) && same(node1.Right, node2.Right)
}
