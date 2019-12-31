package bintree

import (
	"fmt"
	"strings"
)

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) <= 0 || len(postorder) <= 0 {
		return nil
	}
	return helper(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1, 0)
}

func helper(inorder, postorder []int, inLeft, inRight, postLeft, postRight, indent int) *TreeNode {
	if inRight < inLeft {
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
	rootInOrderIndex := inLeft
	for rootInOrderIndex <= inRight {
		if inorder[rootInOrderIndex] == rootVal {
			break
		}
		rootInOrderIndex++
	}
	// Build Left Node
	// find the right boundary of the left side in postorder
	fmt.Printf("%sBuild Left for root %d\n", strings.Repeat("\t", indent), rootVal)
	leftPostOrderRightBoundary := postLeft
	if rootInOrderIndex > inLeft {
		for leftPostOrderRightBoundary <= postRight {
			if inorder[inLeft] == postorder[leftPostOrderRightBoundary] {
				break
			}
			leftPostOrderRightBoundary++
		}

	}
	// postorder right boundary of the left side is the postRight-1-(inRight-rootInOrderIndex)
	lenRight := inRight - rootInOrderIndex
	root.Left = helper(inorder, postorder, inLeft, rootInOrderIndex-1, postLeft, postRight-1-lenRight, indent+1)
	fmt.Println("")
	// Build Right Node
	fmt.Printf("%sBuild Right for root %d\n", strings.Repeat("\t", indent), rootVal)
	root.Right = helper(inorder, postorder, rootInOrderIndex+1, inRight, postRight-lenRight, postRight-1, indent+1)
	fmt.Println("")

	return root
}
