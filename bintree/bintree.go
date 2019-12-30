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

	// uses a stack for the nodes with which we are not done with.
	// Prime the stack with the root and a state struct
	// while there are nodes on the stack
	// for a node and it's state
	// use the state to determine if the current node
	// should:
	// go back on the stack and add the left node;
	// the value from the node needs to be added to the vals;
	// or the right node needs to be added to the stack

	if root == nil {
		return []int{}
	}
	type inorderState struct {
		left bool
		node bool
	}
	states := make([]inorderState, 0)
	nodes := make([]*TreeNode, 0)
	vals := make([]int, 0)
	// traverse left until left is nil
	// append leaf val
	var curr *TreeNode
	nodes = append(nodes, root)
	states = append(states, inorderState{})

	for len(nodes) > 0 {
		curr = nodes[len(nodes)-1]
		s := states[len(states)-1]
		nodes = nodes[:len(nodes)-1]
		states = states[:len(states)-1]
		if !s.left {
			s.left = true
			nodes = append(nodes, curr)
			states = append(states, s)
			if curr.Left != nil {
				nodes = append(nodes, curr.Left)
				states = append(states, inorderState{})
			}
		} else if !s.node {
			vals = append(vals, curr.Val)
			s.node = true
			if curr.Right != nil {
				nodes = append(nodes, curr.Right)
				states = append(states, inorderState{})
			}
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
