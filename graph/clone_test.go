package graph

import "fmt"

/**
 * Definition for a Node.
 */
type Node struct {
	Val       int
	Neighbors []*Node
}

func CloneGraph(node *Node) *Node {
	processedNodes := make(map[*Node]*Node)

	var dfs func(node *Node) *Node

	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if v, ok := processedNodes[node]; ok {
			return v
		}
		startNode := &Node{Val: node.Val, Neighbors: []*Node{}}
		processedNodes[node] = startNode
		fmt.Println(processedNodes)
		if node.Neighbors == nil {
			return startNode
		}
		for _, n := range node.Neighbors {
			startNode.Neighbors = append(startNode.Neighbors, dfs(n))
		}
		return startNode
	}

	return dfs(node)
}
