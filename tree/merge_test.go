package tree

func MergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	newRoot := TreeNode{}
	if root1 != nil && root2 != nil {
		newRoot.Val = root1.Val + root2.Val
	}
	if root1 == nil && root2 != nil {
		newRoot.Val = (*root2).Val
	}
	if root2 == nil && root1 != nil {
		newRoot.Val = (*root1).Val
	}
	l1, r1 := root1, root1
	if root1 != nil {
		l1 = root1.Left
		r1 = root1.Right
	}
	l2, r2 := root2, root2
	if root2 != nil {
		l2 = root2.Left
		r2 = root2.Right
	}
	newRoot.Left = MergeTrees(l1, l2)
	newRoot.Right = MergeTrees(r1, r2)

	return &newRoot
}
