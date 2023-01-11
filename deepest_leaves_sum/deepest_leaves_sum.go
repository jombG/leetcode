package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	m := make(map[int]int)

	var dfs func(node *TreeNode, lvl int)

	dfs = func(node *TreeNode, lvl int) {
		if _, ok := m[lvl]; !ok {
			m[lvl] = node.Val
		} else {
			m[lvl] += node.Val
		}

		if node.Left != nil {
			dfs(node.Left, lvl+1)
		}
		if node.Right != nil {
			dfs(node.Right, lvl+1)
		}
	}

	dfs(root, 0)

	return m[len(m)-1]
}

func main() {
	fmt.Println(
		deepestLeavesSum(&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		}),
	)
}
