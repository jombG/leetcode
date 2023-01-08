package binary_tree_inorder_traversal

import (
	"fmt"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	tree := TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(inorderTraversal(&tree))
}
