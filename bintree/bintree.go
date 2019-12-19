package bintree

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
