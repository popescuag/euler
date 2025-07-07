package tree

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return dfs(root, 0, targetSum)
}

func dfs(root *TreeNode, sum int, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && sum+root.Val == targetSum {
		return true
	}

	return dfs(root.Left, sum+root.Val, targetSum) || dfs(root.Right, sum+root.Val, targetSum)
}
