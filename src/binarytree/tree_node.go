package binarytree

//TreeNode   a binary tree
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

//Build build a binarytree sutrcut
//if data is nil ,the return value would be nil
func Build(data []int) *TreeNode {

	if len(data) < 1 {
		return nil
	}

	node := &TreeNode{
		Value: data[0],
	}
	for i := 1; i < len(data); i++ {
		node.newNode(data[i])
	}

	return node
}

//newNode add a new node on
func (node *TreeNode) newNode(value int) {

	if value < node.Value {
		if node.Left != nil {
			node.Left.newNode(value)
		} else {
			node.Left = &TreeNode{
				Value: value,
			}
			return
		}
	} else {

		if node.Right != nil {
			node.Right.newNode(value)
		} else {
			node.Right = &TreeNode{
				Value: value,
			}
			return
		}
	}
}

//PreOderTraversal visit each node
func (node *TreeNode) PreOderTraversal() []int {

	if node == nil {
		return []int{}
	}
	data := []int{node.Value}
	l := node.Left.PreOderTraversal()
	r := node.Right.PreOderTraversal()
	data = append(data, l...)
	data = append(data, r...)
	return data
}

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

//PostOderTraversal visit each node
func (node *TreeNode) PostOderTraversal() []int {

	if node == nil {
		return []int{}
	}
	r := node.Right.InOderTraversal()

	data := append(r, node.Value)
	l := node.Left.InOderTraversal()
	data = append(data, l...)
	return data
}
