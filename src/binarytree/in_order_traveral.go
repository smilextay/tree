package binarytree

//InOderTraversal visit each node
func (node *TreeNode) InOderTraversal() []int {

	if node == nil {
		return []int{}
	}
	l := node.Left.InOderTraversal()
	data := append(l, node.Value)

	r := node.Right.InOderTraversal()
	data = append(data, r...)
	return data
}
