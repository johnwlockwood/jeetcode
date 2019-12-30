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

// InorderTraversal traverses a binary tree inorder
func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	vals := make([]int, 0)
	lefts := InorderTraversal(root.Left)
	vals = append(vals, lefts...)
	vals = append(vals, root.Val)
	rights := InorderTraversal(root.Right)
	vals = append(vals, rights...)
	return vals
}

// InorderTraversalIterative traverses a binary tree inorder using an iterative approach
func InorderTraversalIterative(root *TreeNode) []int {
	// use a stack to hold nodes which haven't had
	// their value consumed or right added to the stack
	nodes := make([]*TreeNode, 0)
	vals := make([]int, 0)

	for root != nil || len(nodes) > 0 {
		if root != nil {
			nodes = append(nodes, root)
			root = root.Left
		} else {
			curr := nodes[len(nodes)-1]
			nodes = nodes[:len(nodes)-1]
			vals = append(vals, curr.Val)
			root = curr.Right
		}
	}
	return vals
}

// contest/weekly-contest-169/problems/all-elements-in-two-binary-search-trees/
// from contest 169
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	vals := make([]int, 0)
	getVals(root1, &vals)
	getVals(root2, &vals)
	sort.Ints(vals)
	return vals
}

func getVals(node *TreeNode, vals *[]int) {
	// postorder traversal
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
