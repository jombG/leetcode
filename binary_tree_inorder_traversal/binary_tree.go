package binary_tree_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	list []int
}

func (s *Stack) Push(val int) {
	s.list = append(s.list, val)
}

func (s *Stack) Pop() {
	if len(s.list) == 1 {
		s.list = []int{}
	} else {
		s.list = s.list[:len(s.list)-1]
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.list) == 0
}

func (s *Stack) Top() int {
	return s.list[len(s.list)-1]
}

func inorderTraversal(root *TreeNode) []int {
	stack := &Stack{
		list: []int{},
	}
	stack.Push(root.Val)

	helper(root.Left, stack)
	helper(root.Right, stack)

	var results []int
	for !stack.IsEmpty() {
		results = append(results, stack.Top())
		stack.Pop()
	}
	return results
}

func helper(node *TreeNode, stack *Stack) {
	if node == nil {
		return
	}
	stack.Push(node.Val)
	helper(node.Left, stack)
	helper(node.Right, stack)
}
