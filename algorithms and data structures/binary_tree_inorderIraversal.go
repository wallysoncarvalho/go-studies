package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	leftVal := root.Left.Val
	//rightVal := root.Right.Val

	if leftVal < root.Val {

		

	} else {

	}

	inorderTraversal(root.Left)

	inorderTraversal(root.Right)

	return result
}
