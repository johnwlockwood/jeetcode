package bintree

import (
	"fmt"
	"strings"
)

// My solution to problems/construct-binary-tree-from-inorder-and-postorder-traversal
// Hint: look for lengths of sides!

// 12ms!
// 27.6MB
//
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) <= 0 || len(postorder) <= 0 {
		return nil
	}
	inOrderIndexMap := make(map[int]int, len(inorder))
	for i, v := range inorder {
		inOrderIndexMap[v] = i
	}
	// pass slice of the arrays inorder and postorder to the helper.
	// this way they are alway pointing to the same underlying array
	// https://golang.org/ref/spec#Slice_types
	return helper(inorder[:], postorder[:], inOrderIndexMap, 0, len(inorder)-1, 0, len(postorder)-1, 0)
}

func helper(inorder, postorder []int, inOrderIndexMap map[int]int, inLeft, inRight, postLeft, postRight, indent int) *TreeNode {
	if inLeft > inRight {
		return nil
	}
	fmt.Printf("%sinorder:   %v\n", strings.Repeat("\t", indent), inorder[inLeft:inRight+1])
	fmt.Printf("%spostorder: %v\n", strings.Repeat("\t", indent), postorder[postLeft:postRight+1])
	rootVal := postorder[postRight]
	fmt.Printf("%sVal: %d\n", strings.Repeat("\t", indent), rootVal)
	rootInOrderIndex := inOrderIndexMap[rootVal]
	lenLeft := rootInOrderIndex - inLeft
	lenRight := inRight - rootInOrderIndex
	return &TreeNode{
		Val:   rootVal,
		Left:  helper(inorder, postorder, inOrderIndexMap, inLeft, inLeft+lenLeft-1, postLeft, postRight-1-lenRight, indent+1),
		Right: helper(inorder, postorder, inOrderIndexMap, rootInOrderIndex+1, inRight, postRight-lenRight, postRight-1, indent+1),
	}
}

// 20ms
// 20.5MB
func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(inorder) <= 0 || len(postorder) <= 0 {
		return nil
	}
	return helper2(inorder[:], postorder[:])
}

func helper2(inorder, postorder []int) *TreeNode {
	inRight := len(inorder) - 1
	postRight := len(postorder) - 1
	rootVal := postorder[postRight]
	root := &TreeNode{
		Val: rootVal,
	}
	// find the root in inorder
	rootInOrderIndex := 0
	for rootInOrderIndex < len(inorder) {
		if rootVal == inorder[rootInOrderIndex] {
			break
		}
		rootInOrderIndex++
	}
	// Build Left Node
	lenLeft := rootInOrderIndex
	lenRight := inRight - rootInOrderIndex
	if lenLeft > 0 {
		root.Left = helper2(inorder[:lenLeft], postorder[:postRight-lenRight])
	}
	// Build Right Node
	if lenRight > 0 {
		root.Right = helper2(inorder[rootInOrderIndex+1:inRight+1], postorder[postRight-lenRight:postRight])
	}
	return root
}
