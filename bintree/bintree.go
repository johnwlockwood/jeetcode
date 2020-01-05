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

// HeapifyMin min heapifies a tree
// TODO: add tests
// should result in the minimum value being at that root
func HeapifyMin(node *TreeNode) {
	if node == nil {
		return
	}
	HeapifyMin(node.Left)
	HeapifyMin(node.Right)
	smallest := node
	if node.Left != nil && node.Left.Val < node.Val {
		smallest = node.Left
	}
	if node.Right != nil && node.Right.Val < smallest.Val {
		smallest = node.Right
	}
	if smallest != node {
		Swap(node, smallest)
	}
}

// Swap swaps two nodes
func Swap(n1, n2 *TreeNode) {
	tmp := n1.Left
	n1.Left = n2.Left
	n2.Left = tmp

	tmp = n1.Right
	n1.Right = n2.Right
	n2.Right = tmp
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

func getAllElementsPreOrder(root1 *TreeNode, root2 *TreeNode) []int {
	vals := make([]int, 0)
	getValsPreOrder(root1, &vals)
	getValsPreOrder(root2, &vals)
	sort.Ints(vals)
	return vals
}

func getValsPreOrder(node *TreeNode, vals *[]int) {
	// postorder traversal
	if node == nil {
		return
	}
	*vals = append(*vals, node.Val)
	if node.Left != nil {
		getValsPreOrder(node.Left, vals)
	}
	if node.Right != nil {
		getValsPreOrder(node.Right, vals)
	}
}

// /problems/binary-tree-level-order-traversal
func levelOrderValues(root *TreeNode) [][]int {
	// BFS
	if root == nil {
		return [][]int{}
	}
	out := make([][]int, 0)

	levelQueue := make([]int, 0)
	nodeQueue := make([]*TreeNode, 0)

	nodeQueue = append(nodeQueue, root)
	levelQueue = append(levelQueue, 0)
	for len(nodeQueue) > 0 {
		l := levelQueue[0]
		node := nodeQueue[0]
		levelQueue = levelQueue[1:]
		nodeQueue = nodeQueue[1:]
		if len(out) < l+1 {
			out = append(out, []int{})
		}
		out[l] = append(out[l], node.Val)
		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
			levelQueue = append(levelQueue, l+1)
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
			levelQueue = append(levelQueue, l+1)
		}
	}
	return out
}

// based on the given iterative solution
// this avoids needing to keep a queue for the
// level value
func levelOrderValuesIterative(root *TreeNode) [][]int {
	// BFS
	if root == nil {
		return [][]int{}
	}
	out := make([][]int, 0)

	nodeQueue := make([]*TreeNode, 0)
	nodeQueue = append(nodeQueue, root)

	level := 0

	for len(nodeQueue) > 0 {
		levelSize := len(nodeQueue)
		// start the current level
		if len(out) < level+1 {
			out = append(out, []int{})
		}
		for i := 0; i < levelSize; i++ {
			node := nodeQueue[0]
			nodeQueue = nodeQueue[1:]

			// fufill the current level
			out[level] = append(out[level], node.Val)

			// ad child nodes of the current level
			// in the queue for the next level
			if node.Left != nil {
				nodeQueue = append(nodeQueue, node.Left)
			}
			if node.Right != nil {
				nodeQueue = append(nodeQueue, node.Right)
			}
		}
		// go to the next level
		level++
	}
	return out
}
