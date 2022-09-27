package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(node *TreeNode) []int {
	var result []int

	if node == nil {
		return nil
	}

	result = append(result, inorderTraversal(node.Left)...)
	result = append(result, node.Val)
	result = append(result, inorderTraversal(node.Right)...)

	return result
}

func preOrderTraversal(node *TreeNode) []int {
	var result []int

	if node == nil {
		return nil
	}

	result = append(result, node.Val)
	result = append(result, inorderTraversal(node.Left)...)
	result = append(result, inorderTraversal(node.Right)...)

	return result
}

func posOrderTraversal(node *TreeNode) []int {
	var result []int

	if node == nil {
		return nil
	}

	result = append(result, inorderTraversal(node.Left)...)
	result = append(result, inorderTraversal(node.Right)...)
	result = append(result, node.Val)

	return result
}
