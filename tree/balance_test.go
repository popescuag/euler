package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
	nums := inorder(root)
	return balance(nums, 0, len(nums)-1)
}

func balance(nums []*TreeNode, l int, r int) *TreeNode {
	if l > r {
		return nil
	}
	m := (l + r) / 2
	node := TreeNode{Val: nums[m].Val, Left: balance(nums, l, m-1), Right: balance(nums, m+1, r)}
	return &node
}

func inorder(node *TreeNode) []*TreeNode {
	if node == nil {
		return []*TreeNode{}
	}
	if node.Left == nil && node.Right == nil {
		return []*TreeNode{node}
	}
	arr := inorder(node.Left)
	arr = append(arr, node)
	arr2 := inorder(node.Right)
	arr = append(arr, arr2...)
	return arr
}

func TestBalance(t *testing.T) {
	root := TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}}
	newRoot := balanceBST(&root)
	assert.Equal(t, 2, newRoot.Val)
	assert.Equal(t, 3, newRoot.Right.Val)
	assert.Equal(t, 1, newRoot.Left.Val)
}
