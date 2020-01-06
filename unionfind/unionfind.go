package unionfind

import "fmt"

// UnionFind is a weighted union find or disjoint set data structure.
// learned from https://www.hackerearth.com/practice/notes/disjoint-set-union-union-find/
type UnionFind struct {
	arr  []int
	size []int
	n    int
}

// NewUnionFind makes a new UnionFind
func NewUnionFind(n int) *UnionFind {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	size := make([]int, n)
	for i := 0; i < n; i++ {
		size[i] = 1
	}
	return &UnionFind{arr: arr, n: n, size: size}
}

// Root finds the root a node and compresses the path to the root
func (u *UnionFind) Root(i int) (int, error) {
	if i < 0 || i >= u.n {
		return -1, u.outOfBoundsError(i)
	}
	return u.root(i), nil
}

func (u *UnionFind) root(i int) int {
	for u.arr[i] != i {
		// path compression: point each to it's grandparent.
		// A root is it's own parent, so it's immediate children will find it as it's grandparent.
		u.arr[i] = u.arr[u.arr[i]]
		i = u.arr[i]
	}
	return i
}

func (u *UnionFind) outOfBoundsError(i int) error {
	return OutOfBoundsError{Upper: u.n, Lower: 0, Index: i}
}

// OutOfBoundsError represents when something is out of bound
type OutOfBoundsError struct {
	Upper int
	Lower int
	Index int
}

func (e OutOfBoundsError) Error() string {
	return fmt.Sprintf("%d is not within %d and %d", e.Index, e.Lower, e.Upper)
}

// Find finds if the two nodes are connected
func (u *UnionFind) Find(a, b int) (bool, error) {
	if a < 0 || a >= u.n {
		return false, u.outOfBoundsError(a)
	}
	if b < 0 || b >= u.n {
		return false, u.outOfBoundsError(b)
	}
	if u.root(a) == u.root(b) {
		return true, nil
	}
	return false, nil
}

// Union connects two nodes in a way to keep the tree balanced
func (u *UnionFind) Union(a, b int) error {
	if a < 0 || a >= u.n {
		return u.outOfBoundsError(a)
	}
	if b < 0 || b >= u.n {
		return u.outOfBoundsError(b)
	}
	rootA := u.root(a)
	rootB := u.root(b)
	if u.size[rootA] < u.size[rootB] {
		u.arr[rootA] = u.arr[rootB]
		u.size[rootB] += u.size[rootA]
	} else {
		u.arr[rootB] = u.arr[rootA]
		u.size[rootA] += u.size[rootB]
	}
	return nil
}
