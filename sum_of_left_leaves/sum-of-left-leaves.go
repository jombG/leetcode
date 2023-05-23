package sum_of_left_leaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0

	var r func(node *TreeNode)

	r = func(node *TreeNode) {
		if node == nil {
			return
		}
		left := node.Left
		if left != nil && (left.Left == nil && left.Right == nil) {
			sum = sum + left.Val
		}

		r(node.Left)
		r(node.Right)
	}

	r(root)

	return sum
}
