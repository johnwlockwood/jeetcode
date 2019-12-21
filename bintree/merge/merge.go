package merge

import (
	"github.com/johnwlockwood/jeetcode/bintree"
)

// Example 1
// Input:
// 	Tree 1                     Tree 2
//           1                         2
//          / \                       / \
//         3   2                     1   3
//        /                           \   \
//       5                             4   7
// Output:
// Merged tree:
// 	     3
// 	    / \
// 	   4   5
// 	  / \   \
// 	 5   4   7
func mergeTrees(t1 *bintree.TreeNode, t2 *bintree.TreeNode) *bintree.TreeNode {
	outNode := &bintree.TreeNode{}
	if t1 != nil && t2 != nil {
		outNode.Val = t1.Val + t2.Val
		outNode.Left = mergeTrees(t1.Left, t2.Left)
		outNode.Right = mergeTrees(t1.Right, t2.Right)
		return outNode
	} else if t1 != nil {
		outNode.Val = t1.Val
		outNode.Left = mergeTrees(t1.Left, nil)
		outNode.Right = mergeTrees(t1.Right, nil)
		return outNode
	} else if t2 != nil {
		outNode.Val = t2.Val
		outNode.Left = mergeTrees(nil, t2.Left)
		outNode.Right = mergeTrees(nil, t2.Right)
		return outNode
	}
	return nil
}

// Merge Tree Leetcode Approach 1
func mergeTreesApproach1(t1 *bintree.TreeNode, t2 *bintree.TreeNode) *bintree.TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}
