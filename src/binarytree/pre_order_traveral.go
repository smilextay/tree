package binarytree

//PreOderTraversal visit each node
func (node *TreeNode) PreOderTraversal() []int {

	if node == nil {
		return []int{}
	}

	l := node.Left.PreOderTraversal()
	r := node.Right.PreOderTraversal()

	data := []int{node.Value}
	data = append(data, l...)
	data = append(data, r...)
	return data
}
