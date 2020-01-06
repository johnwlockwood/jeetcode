package unionfind

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
func (u *UnionFind) Root(i int) int {
	for u.arr[i] != i {
		// path compression: point each to it's grandparent.
		// A root is it's own parent, so it's immediate children will find it as it's grandparent.
		u.arr[i] = u.arr[u.arr[i]]
		i = u.arr[i]
	}
	return i
}

// Find finds if the two nodes are connected
func (u *UnionFind) Find(a, b int) bool {
	if u.Root(a) == u.Root(b) {
		return true
	}
	return false
}

// Union connects two nodes in a way to keep the tree balanced
func (u *UnionFind) Union(a, b int) {
	rootA := u.Root(a)
	rootB := u.Root(b)
	if u.size[rootA] < u.size[rootB] {
		u.arr[rootA] = u.arr[rootB]
		u.size[rootB] += u.size[rootA]
	} else {
		u.arr[rootB] = u.arr[rootA]
		u.size[rootA] += u.size[rootB]
	}
}
