package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}
	result := [][]int{}

	arr := []TreeNode{}
	arr = append(arr, *root)

	for len(arr) > 0 {
		l := len(arr)
		level := []int{}
		for i := 0; i < l; i++ {
			currentNode := arr[0]
			arr = arr[1:]

			level = append(level, currentNode.Val)

			if currentNode.Left != nil {
				arr = append(arr, *currentNode.Left)
			}
			if currentNode.Right != nil {
				arr = append(arr, *currentNode.Right)
			}
		}
		result = append(result, level)
	}

	return result
}

func IslandPerimeter(grid [][]int) int {
	rc := len(grid)
	cc := len(grid[0])

	if rc == 1 && cc == 1 {
		return 4
	}

	perimeter := 0
	for i := range rc {
		for j := range cc {
			perimeter += edgesCount(grid, i, j)
		}
	}
	return perimeter
}

func edgesCount(grid [][]int, i int, j int) int {
	rc := len(grid)
	cc := len(grid[0])
	if i < 0 || j < 0 || i > rc-1 || j > cc-1 {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}
	count := 0
	if i < rc-1 && grid[i+1][j] == 1 {
		count++
	}
	if i > 0 && grid[i-1][j] == 1 {
		count++
	}

	if j < cc-1 && grid[i][j+1] == 1 {
		count++
	}
	if j > 0 && grid[i][j-1] == 1 {
		count++
	}

	return 4 - count
}
