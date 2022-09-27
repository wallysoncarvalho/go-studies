package main

import (
	"reflect"
	"testing"
)

func testBinaryTree() *TreeNode {
	leaf1 := TreeNode{Val: 15, Left: nil, Right: nil}
	leaf2 := TreeNode{Val: 28, Left: nil, Right: nil}
	leaf3 := TreeNode{Val: 55, Left: nil, Right: nil}
	leaf4 := TreeNode{Val: 70, Left: nil, Right: nil}

	lv31 := TreeNode{Val: 25, Left: &leaf1, Right: &leaf2}
	lv32 := TreeNode{Val: 35, Left: nil, Right: nil}
	lv33 := TreeNode{Val: 45, Left: nil, Right: nil}
	lv34 := TreeNode{Val: 60, Left: &leaf3, Right: &leaf4}

	lv21 := TreeNode{Val: 30, Left: &lv31, Right: &lv32}
	lv22 := TreeNode{Val: 50, Left: &lv33, Right: &lv34}

	return &TreeNode{Val: 40, Left: &lv21, Right: &lv22}
}

func TestTraverseBinaryTreeInOrder(t *testing.T) {
	tree := *testBinaryTree()

	expected := []int{15, 25, 28, 30, 35, 40, 45, 50, 55, 60, 70}

	if result := inorderTraversal(&tree); !reflect.DeepEqual(expected, result) {
		t.Errorf("got %v but was expecting %v", result, expected)
	}
}

func TestTraverseBinaryTreePreOrder(t *testing.T) {
	tree := *testBinaryTree()

	expected := []int{40, 15, 25, 28, 30, 35, 45, 50, 55, 60, 70}

	if result := preOrderTraversal(&tree); !reflect.DeepEqual(expected, result) {
		t.Errorf("got %v but was expecting %v", result, expected)
	}
}

func TestTraverseBinaryTreePosOrder(t *testing.T) {
	tree := *testBinaryTree()

	expected := []int{15, 25, 28, 30, 35, 45, 50, 55, 60, 70, 40}

	if result := posOrderTraversal(&tree); !reflect.DeepEqual(expected, result) {
		t.Errorf("got %v but was expecting %v", result, expected)
	}
}
