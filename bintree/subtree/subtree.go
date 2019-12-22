package subtree

import (
	"strings"

	"github.com/johnwlockwood/jeetcode/bintree"
)

// easy problem

func isEqual(x *bintree.TreeNode, y *bintree.TreeNode) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return x.Val == y.Val && isEqual(x.Left, y.Left) && isEqual(x.Right, y.Right)
}

func traverse(s *bintree.TreeNode, t *bintree.TreeNode) bool {
	return s != nil && (isEqual(s, t) || traverse(s.Left, t) || traverse(s.Right, t))
}
func isSubtree(s *bintree.TreeNode, t *bintree.TreeNode) bool {
	return traverse(s, t)
}

func isSubtreeWithPreOrder(s *bintree.TreeNode, t *bintree.TreeNode) bool {
	sS := bintree.PreOrder(s)
	sT := bintree.PreOrder(t)
	return strings.Contains(sS, sT)
}
