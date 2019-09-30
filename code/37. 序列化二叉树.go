func Serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return fmt.Sprintf("%d!%s!%s",
		root.Val, Serialize(root.Left), Serialize(root.Right))
}

func Deserialize(str string) *TreeNode {
	values := strings.Split(str, "!")
	return deserialize(&values)
}

func deserialize(values *[]string) *TreeNode {
	if len(*values) == 0 {
		return nil
	}
	valStr := (*values)[0]
	*values = (*values)[1:]
	if valStr == "#" {
		return nil
	}
	val, _ := strconv.Atoi(valStr)
	root := &TreeNode{Val: val}
	root.Left = deserialize(values)
	root.Right = deserialize(values)
	return root
}