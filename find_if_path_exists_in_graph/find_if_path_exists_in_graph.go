package find_if_path_exists_in_graph

func validPath(n int, edges [][]int, source int, destination int) bool {
	var dfs func(node int) bool
	visited := make([]bool, n)

	mNode := make(map[int][]int)

	for _, e := range edges {
		if _, ok := mNode[e[0]]; !ok {
			mNode[e[0]] = append([]int{}, e[1])
		} else {
			mNode[e[0]] = append(mNode[e[0]], e[1])
		}

		if _, ok := mNode[e[1]]; !ok {
			mNode[e[1]] = append([]int{}, e[0])
		} else {
			mNode[e[1]] = append(mNode[e[1]], e[0])
		}
	}

	dfs = func(node int) bool {
		if node == destination {
			return true
		}

		if !visited[node] {
			visited[node] = true
			for _, v := range mNode[node] {
				if dfs(v) {
					return true
				}
			}
		}

		return false
	}

	return dfs(source)
}
