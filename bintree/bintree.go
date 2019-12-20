package bintree

import "fmt"

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
func PreOrder(root *TreeNode, left bool) string {
	if root == nil {
		if left {
			return "lnil"
		} else {
			return "rnil"
		}
	}
	s := fmt.Sprintf("#%d %s %s", root.Val, PreOrder(root.Left, true), PreOrder(root.Right, false))
	return s
}
