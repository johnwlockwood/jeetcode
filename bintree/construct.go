package bintree

import (
	"fmt"
	"strings"
)

// My solution to problems/construct-binary-tree-from-inorder-and-postorder-traversal
// Hint: look for lengths of sides!
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) <= 0 || len(postorder) <= 0 {
		return nil
	}
	inOrderIndexMap := make(map[int]int, len(inorder))
	for i, v := range inorder {
		inOrderIndexMap[v] = i
	}
	return helper(inorder, postorder, inOrderIndexMap, 0, len(inorder)-1, 0, len(postorder)-1, 0)
}

func helper(inorder, postorder []int, inOrderIndexMap map[int]int, inLeft, inRight, postLeft, postRight, indent int) *TreeNode {
	if inLeft > inRight {
		return nil
	}
	fmt.Printf("%sinorder:   %v\n", strings.Repeat("\t", indent), inorder[inLeft:inRight+1])
	fmt.Printf("%spostorder: %v\n", strings.Repeat("\t", indent), postorder[postLeft:postRight+1])

	rootVal := postorder[postRight]
	root := &TreeNode{
		Val: rootVal,
	}

	fmt.Printf("%sVal: %d\n", strings.Repeat("\t", indent), rootVal)

	// find the root in inorder
	rootInOrderIndex := inOrderIndexMap[rootVal]
	// Build Left Node
	fmt.Printf("%sBuild Left for root %d\n", strings.Repeat("\t", indent), rootVal)
	lenLeft := rootInOrderIndex - inLeft
	lenRight := inRight - rootInOrderIndex
	root.Left = helper(inorder, postorder, inOrderIndexMap, inLeft, inLeft+lenLeft-1, postLeft, postRight-1-lenRight, indent+1)
	fmt.Println("")
	// Build Right Node
	fmt.Printf("%sBuild Right for root %d\n", strings.Repeat("\t", indent), rootVal)
	root.Right = helper(inorder, postorder, inOrderIndexMap, rootInOrderIndex+1, inRight, postRight-lenRight, postRight-1, indent+1)
	fmt.Println("")

	return root
}
