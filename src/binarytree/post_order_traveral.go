package binarytree

//PostOderTraversal visit each node
func (node *TreeNode) PostOderTraversal() []int {

	if node == nil {
		return []int{}
	}
	l := node.Left.PostOderTraversal()
	r := node.Right.PostOderTraversal()
	data := l
	data = append(data, r...)
	data = append(data, node.Value)
	return data
}
