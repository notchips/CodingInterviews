package code

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Build(pre []int, in []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	var pos int
	for i, v := range in {
		if v == pre[0] {
			pos = i
			break
		}
	}
	return &TreeNode{
		Val:   pre[0],
		Left:  Build(pre[1:pos+1], in[:pos]),
		Right: Build(pre[pos+1:], in[pos+1:]),
	}
}
