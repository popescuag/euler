package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func buildTree() MTreeNode[int] {
	root := MTreeNode[int]{
		Data: 100,
	}
	child1 := MTreeNode[int]{Data: 50}
	child1.appendChild(&MTreeNode[int]{Data: 150})
	child1.appendChild(&MTreeNode[int]{Data: 250})
	root.appendChild(&child1)
	root.appendChild(&MTreeNode[int]{Data: 150})
	root.appendChild(&MTreeNode[int]{Data: 250})

	return root
}

func sumTree(node MTreeNode[int]) int {
	if node.Children == nil {
		return node.Data
	}
	sum := node.Data
	for _, child := range node.Children {
		sum += sumTree(*child)
	}
	return sum
}

func (node MTreeNode[T]) BFS(nodeData *[]T) {
	if node.Children == nil {
		return
	}
	//(*nodeData) = append(*nodeData, node.Data)
	for _, child := range node.Children {
		(*nodeData) = append(*nodeData, child.Data)
		child.BFS(nodeData)
	}
}

func (node MTreeNode[T]) DFS(nodeData *[]T) {
	if node.Children == nil {
		return
	}
	//(*nodeData) = append(*nodeData, node.Data)
	for _, child := range node.Children {
		child.DFS(nodeData)
		(*nodeData) = append(*nodeData, child.Data)
	}
}

func TestSumTree(t *testing.T) {
	root := buildTree()
	assert.Equal(t, 950, sumTree(root))
}

func TestBFS(t *testing.T) {
	root := buildTree()
	nodeData := []int{}
	expectedData := []int{50, 150, 250, 150, 250}
	root.BFS(&nodeData)
	assert.Equal(t, expectedData, nodeData)
}

func TestDFS(t *testing.T) {
	root := buildTree()
	nodeData := []int{}
	expectedData := []int{150, 250, 50, 150, 250}
	root.DFS(&nodeData)
	assert.Equal(t, expectedData, nodeData)
}
