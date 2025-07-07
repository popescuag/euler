package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func SumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 0

	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			result += root.Left.Val
		}
		result += SumOfLeftLeaves(root.Left)
	}
	result += SumOfLeftLeaves(root.Right)

	return result
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}
