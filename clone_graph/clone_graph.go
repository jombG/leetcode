package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[int]*Node)

	var _cloneGroup func(node *Node) *Node

	_cloneGroup = func(node *Node) *Node {
		if node == nil {
			return nil
		}

		dupe := *node
		dupe.Neighbors = make([]*Node, len(node.Neighbors))
		visited[dupe.Val] = &dupe
		for i, n := range node.Neighbors {
			if visited[n.Val] != nil {
				dupe.Neighbors[i] = visited[n.Val]
			} else {
				dupe.Neighbors[i] = _cloneGroup(n)
			}
		}

		return &dupe
	}

	return _cloneGroup(node)
}

func main() {
	cloneGraph(nil)
}
