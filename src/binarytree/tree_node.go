package binarytree

//TreeNode   a binary tree
type TreeNode struct {
	Value  int
	Parent *TreeNode
	Left   *TreeNode
	Right  *TreeNode
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
				Parent: node,
				Value:  value,
			}
			return
		}
	} else {

		if node.Right != nil {
			node.Right.newNode(value)
		} else {
			node.Right = &TreeNode{
				Parent: node,
				Value:  value,
			}
			return
		}
	}
}
