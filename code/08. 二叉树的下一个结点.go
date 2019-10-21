package code

// 中序遍历：先遍历左子树，然后根结点，最后右子树
func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
	if pNode == nil {
		return nil
	}
	// 如果pNode存在右子树，则将pNode作为父结点看待，返回其右子树中最左边的结点
	if pNode.Right != nil {
		node := pNode.Right
		for node.Left != nil {
			node = node.Left
		}
		return node
	}
	for pNode.Father != nil { // 如果pNode存在父结点
		// 当pNode是父结点的左子树时，说明pNode作为左子树已经遍历完了，返回其父结点
		if pNode == pNode.Father.Left {
			return pNode.Father
		}
		pNode = pNode.Father
	}
	return nil
}
