package binarytree

import (
	"testing"
)

//TestBuild a test function for build binarytree
func TestBuild(t *testing.T) {

	data := []int{5, 8, 6, 3, 4, 7, 1, 2}

	node := Build(data)
	if node == nil {
		t.Fatal("nil tree")
	}
	if node.Value != 5 {
		t.Fatal("tree error")
	}
	if node.Left == nil || node.Left.Value != 3 {
		t.Fatal("trre error on left node,parent node is ", node.Left.Value)
	}

	if node.Left.Left == nil || node.Left.Left.Value != 1 {
		t.Fatal("trre error on left node,parent node is ", node.Left.Value)
	}
	{

		if node.Left.Left.Left != nil {
			t.Fatal("trre error on left node,parent node is ", node.Left.Left.Value)
		}
		if node.Left.Left.Right == nil || node.Left.Left.Right.Value != 2 {
			t.Fatal("trre error on right node,parent node is ", node.Left.Left.Left.Value)
		}
		{
			if node.Left.Left.Right.Left != nil || node.Left.Left.Right.Right != nil {
				t.Fatal("trre error on left node,parent node is ", node.Left.Left.Right.Value)
			}
		}

		if node.Left.Right == nil || node.Left.Right.Value != 4 {
			t.Fatal("trre error on left node,parent node is ", node.Left)
		}
		{
			if node.Left.Right.Left != nil || node.Left.Right.Right != nil {
				t.Fatal("trre error on left node or right node,parent node is ", node.Left.Right.Value)
			}
		}
	}
	if node.Right == nil || node.Right.Value != 8 {
		t.Fatal("trre error on right node,patent node is ", node.Value)
	}
	{

		if node.Right.Left == nil || node.Right.Left.Value != 6 {
			t.Fatal("trre error on left node,patent node is ", node.Right.Value)
		}
		{
			if node.Right.Left.Left != nil {
				t.Fatal("trre error on left node,patent node is ", node.Right.Left.Value)
			}
			if node.Right.Left.Right == nil || node.Right.Left.Right.Value != 7 {
				t.Fatal("trre error on right node,patent node is ", node.Right.Left.Value)
			}
			{
				if node.Right.Left.Right.Right != nil || node.Right.Left.Right.Left != nil {
					t.Fatal("trre error on right node,patent node is ", node.Right.Left.Right)
				}
			}
		}
		if node.Right.Right != nil {
			t.Fatal("trre error on right node,patent node is ", node.Right.Value)
		}
	}

}
func TestPreOderTraversal(t *testing.T) {
	data := []int{5, 8, 6, 3, 4, 7, 1, 2}
	except := []int{5, 3, 1, 2, 4, 8, 6, 7}
	node := Build(data)
	result := node.PreOderTraversal()
	if len(result) != len(except) {
		t.Fatal("un except result in PreOderTraversal:", result)
	}
	for i := 0; i < len(result); i++ {
		if except[i] != result[i] {
			t.Fatal("un except result in PreOderTraversal:", result)
		}
	}
}

func TestInOderTraversal(t *testing.T) {
	data := []int{5, 8, 6, 3, 4, 7, 1, 2}
	except := []int{1, 2, 3, 4, 5, 6, 7, 8}
	node := Build(data)
	result := node.InOderTraversal()
	if len(result) != len(except) {
		t.Fatal("un except result in InOderTraversal:", result)
	}
	for i := 0; i < len(result); i++ {
		if except[i] != result[i] {
			t.Fatal("un except result in InOderTraversal:", result)
		}
	}
}
func TestPostOderTraversal(t *testing.T) {
	data := []int{5, 8, 6, 3, 4, 7, 1, 2}
	except := []int{2, 1, 4, 3, 7, 6, 8, 5}
	node := Build(data)
	result := node.PostOderTraversal()
	if len(result) != len(except) {
		t.Fatal("un except result in PostOderTraversal:", result)
	}
	for i := 0; i < len(result); i++ {
		if except[i] != result[i] {
			t.Fatal("un except result in PostOderTraversal:", result)
		}
	}
}
func TestBFS(t *testing.T) {
	data := []int{5, 8, 6, 3, 4, 7, 1, 2}
	except := []int{5, 3, 8, 1, 4, 6, 2, 7}
	node := Build(data)
	result := node.BFS()
	if len(result) != len(except) {
		t.Fatal("un except result in PostOderTraversal:", result)
	}
	for i := 0; i < len(result); i++ {
		if except[i] != result[i] {
			t.Fatal("un except result in PostOderTraversal:", result)
		}
	}
}

func TestMaxDepth(t *testing.T) {
	data := []int{5, 8, 6, 3, 4, 7, 1, 2}
	node := Build(data)
	except := 4
	result := node.MaxDep()
	if except != result {
		t.Fatal("un excpet result :", result)
	}

}
