package main

import "fmt"

func possibleBipartition(n int, dislikes [][]int) bool {
	m := make(map[int][]int, n)
	for _, v := range dislikes {
		if _, ok := m[v[0]]; !ok {
			m[v[0]] = append([]int{}, v[1])
			m[v[1]] = append([]int{}, v[0])
		} else {
			m[v[0]] = append(m[v[0]], v[1])
			m[v[1]] = append(m[v[1]], v[0])
		}
	}

	color := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if color[i] == 0 {
			if !isBipartite(m, i, color) {
				return false
			}
		}
	}

	return true
}

func isBipartite(adj map[int][]int, node int, color []int) bool {
	var queue []int
	queue = append(queue, node)
	color[node] = 2

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, v := range adj[cur] {
			if color[v] == color[cur] {
				return false
			}

			if color[v] == 0 {
				color[v] = 3 - color[cur]
				queue = append(queue, v)
			}
		}
	}

	return false
}

func main() {
	fmt.Println(possibleBipartition(4, [][]int{{1, 2}, {1, 3}, {2, 4}}))
}
