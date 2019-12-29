package bintree

import (
	"fmt"
	"sort"
)

// TreeNode is a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// GetVals gets the list of values of the tree in a depth first traversal
func GetVals(root *TreeNode) []int {
	if root == nil {
		return []int{-1}
	}
	vals := []int{root.Val}
	vals = append(vals, GetVals(root.Left)...)
	vals = append(vals, GetVals(root.Right)...)
	return vals
}

// PreOrder builds a string representation of the Tree
func PreOrder(root *TreeNode) string {
	s := ""
	if root == nil {
		return "nil "
	}
	s += PreOrder(root.Left)
	s += PreOrder(root.Right)
	s += fmt.Sprintf("%d ", root.Val)
	return s
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	vals := make([]int, 0)
	getVals(root1, &vals)
	getVals(root2, &vals)
	sort.Ints(vals)
	return vals
}

func getVals(node *TreeNode, vals *[]int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		getVals(node.Left, vals)
	}
	if node.Right != nil {
		getVals(node.Right, vals)
	}
	*vals = append(*vals, node.Val)
}
