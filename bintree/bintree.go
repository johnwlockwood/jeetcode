package bintree

import (
	"fmt"
	"math"
	"sort"

	"github.com/johnwlockwood/jeetcode/permutations"
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

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	vals := make([]int, 0)
	getVals(root1, &vals)
	getVals(root2, &vals)
	sort.Ints(vals)
	return vals
}

func getVals(node *TreeNode, vals *[]int) {
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

func helper(arr []int, visited []bool, start int) bool {
	if visited[start] {
		return false
	}
	visited[start] = true
	// for start, is if left is in range,
	if arr[start] == 0 {
		return true
	}
	l, r := start-arr[start], start+arr[start]
	n := len(arr)
	if l != start && l >= 0 && l < n {
		if helper(arr, visited, l) {
			return true
		}
	}
	if r != start && r >= 0 && r < n {
		if helper(arr, visited, r) {
			return true
		}
	}
	return false
}

// Jump Game III
func canReach(arr []int, start int) bool {
	visited := make([]bool, len(arr))
	return helper(arr, visited, start)
}

func decode(word string, charMap map[byte]int) int {
	total := 0
	for i := 0; i < len(word); i++ {
		total += charMap[word[i]] * int(math.Pow10(len(word)-1-i))
	}
	return total
}

// https://leetcode.com/problems/verbal-arithmetic-puzzle/
// Does not solve
// TODO: solve
// a solution is some form of backtracking, but instead of adding to the output, check
// if the permutation solves it
// even with checking all of the permutations of mappings, this doesn't
// always produce a result.
func isSolvable(words []string, result string) bool {
	// find unique set of characters
	fmt.Printf("%s\n", result)
	chars := make(map[byte]struct{}, 0)
	for _, word := range append(words, result) {
		for i := 0; i < len(word); i++ {
			chars[word[i]] = struct{}{}
		}
	}
	charList := make([]byte, 0, len(chars))
	charMap := make(map[byte]int, len(chars))
	for k := range chars {
		fmt.Printf("%s", string(k))
		charList = append(charList, k)
	}
	fmt.Printf("\n")
	indexes := make([]int, len(chars))
	for i := range indexes {
		indexes[i] = i
	}
	perms := permutations.PermuteBackTrackPractice(indexes)
	fmt.Printf("perms count %d\n", len(perms))
	for _, perm := range perms {
		for i, v := range perm {
			charMap[charList[i]] = v
		}
		// fmt.Printf("\n")
		// for _, word := range append(words, result) {
		// 	fmt.Printf("%s decoded to: %d\n", word, decode(word, charMap))
		// }

		total := 0
		for _, word := range words {
			total += decode(word, charMap)
		}
		resultTotal := decode(result, charMap)
		// fmt.Printf("? %d == %d\n", total, resultTotal)
		// fmt.Printf("\n")
		if total == resultTotal {
			fmt.Printf("%d == %d for charMap %v\n", total, resultTotal, charMap)
			fmt.Printf("\n")
			return true
		}
	}
	fmt.Printf("\n")
	return false
}
