package code

// 当p和q都存在root中，返回其最近公共祖先
// 当p或q有且只有一个存在root中，则返回p或q
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil { // p和q分别存在root两侧子树
		return root
	}
	if left != nil {
		return left
	}
	return right
}
