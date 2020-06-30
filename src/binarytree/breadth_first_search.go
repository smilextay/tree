package binarytree

//BFS Breath-first search
func (node *TreeNode) BFS() []int {
	data := []int{}
	if node == nil {
		return data
	}
	data = append(data, node.Value)
	result := breathFirstSearch([]*TreeNode{node.Left, node.Right})
	data = append(data, result...)
	return data
}

func breathFirstSearch(nodes []*TreeNode) []int {
	result := []int{}
	if len(nodes) < 1 {
		return result
	}
	nextLvNode := []*TreeNode{}
	for _, node := range nodes {
		result = append(result, node.Value)
		if node.Left != nil {
			nextLvNode = append(nextLvNode, node.Left)
		}
		if node.Right != nil {
			nextLvNode = append(nextLvNode, node.Right)
		}
	}
	res := breathFirstSearch(nextLvNode)
	result = append(result, res...)
	return result
}
