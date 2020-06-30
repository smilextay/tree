package binarytree

//MaxDep get the max depth of a node
func (node *TreeNode) MaxDep() int {
	if node == nil {
		return 0
	}
	l := node.Left.MaxDep()
	r := node.Right.MaxDep()
	if r > l {
		return r + 1
	}
	return l + 1
}
