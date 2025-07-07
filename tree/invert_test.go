package tree

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if (root.Left != nil && root.Right != nil) || (root.Left != nil && root.Right == nil) || (root.Left == nil && root.Right != nil) {
		root.Left, root.Right = root.Right, root.Left
	}

	InvertTree(root.Left)
	InvertTree(root.Right)

	return root
}
