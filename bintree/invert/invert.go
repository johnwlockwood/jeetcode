package invert

import (
	"github.com/johnwlockwood/jeetcode/bintree"
)

func invertTree(root *bintree.TreeNode) *bintree.TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
