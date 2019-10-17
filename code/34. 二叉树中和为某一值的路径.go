package code

func FindPath(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	var paths [][]int
	path := []int{root.Val}
	dfs2(root, target-root.Val, &paths, &path)
	return paths
}

func dfs2(root *TreeNode, sum int, paths *[][]int, path *[]int) {
	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			newPath := make([]int, len(*path))
			copy(newPath, *path)
			*paths = append(*paths, newPath)
		}
		return
	}
	for _, child := range []*TreeNode{root.Left, root.Right} {
		if child != nil {
			*path = append(*path, child.Val)
			dfs2(child, sum-child.Val, paths, path)
			*path = (*path)[:len(*path)-1]
		}
	}
}