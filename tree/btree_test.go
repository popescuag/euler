package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func buildBTree() []BTreeNode[int] {
	root := BTreeNode[int]{
		Data: 100,
	}
	left := BTreeNode[int]{
		Data: 50,
		Left: &BTreeNode[int]{
			Data: 30,
		}}
	root.appendLeftNode(&left)

	root.appendRightNode(&BTreeNode[int]{
		Data: 20,
		Right: &BTreeNode[int]{
			Data: 100,
		},
	})

	return []BTreeNode[int]{root, left}
}

func sum(node BTreeNode[int]) int {
	if node.Left == nil && node.Right == nil {
		return node.Data
	}
	s := node.Data
	if node.Left == nil {
		s += sum(*node.Right)
	} else if node.Right == nil {
		s += sum(*node.Left)
	} else {
		s += sum(*node.Left) + sum(*node.Right)
	}
	return s
}

func TestSum(t *testing.T) {
	nodes := buildBTree()
	assert.Equal(t, 300, sum(nodes[0]))
	assert.Equal(t, 80, sum(nodes[1]))
}
