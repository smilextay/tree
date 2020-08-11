package binarytree

//isBalanced to
func (node *NodeTree) isBalanced() (int, bool) {

	if node == nil {
		return 0, true
	}
	l, ok := node.Left.isBalanced()
	if !ok {
		return -1, ok
	}
	r, ok := node.Right.isBalanced()
	if !ok {
		return -1, ok
	}
	diff := r - l
	if diff > 1 || diff < -1 {
		return -1
	}
	if r > l {
		return r + 1, true
	}
	return l + 1, true
}
