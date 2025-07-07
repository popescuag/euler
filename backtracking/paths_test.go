package backtracking

func AllPathsSourceTarget(graph [][]int) [][]int {
	lastNode := len(graph) - 1
	path := []int{0}
	result := [][]int{}

	var dfs func(currNode int)

	dfs = func(currNode int) {
		if currNode == lastNode {
			tmpPath := make([]int, len(path))
			copy(tmpPath, path)
			result = append(result, tmpPath)
			return
		}

		if len(graph[currNode]) == 0 {
			return
		}

		for _, newNode := range graph[currNode] {
			path = append(path, newNode)
			dfs(newNode)
			path = path[:len(path)-1]
		}
	}

	dfs(0)

	return result
}
