package islands

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	arr := make([]int, m*n)
	initialize(arr, m*n)
	// stub
	return 0
}

type UnionFind struct {
	arr  []int
	size []int
	n    int
}

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

func (u *UnionFind) Root(i int) int {
	for u.arr[i] != i {
		// path compression: point each to it's grandparent.
		// A root is it's own parent, so it's immediate children will find it as it's grandparent.
		u.arr[i] = u.arr[u.arr[i]]
		i = u.arr[i]
	}
	return i
}

func (u *UnionFind) Find(a, b int) bool {
	if u.Root(a) == u.Root(b) {
		return true
	}
	return false
}

func (u *UnionFind) Union(a, b int) {
	rootA := u.Root(a)
	rootB := u.Root(b)
	u.arr[rootA] = rootB
}

func (u *UnionFind) WeightedUnion(a, b int) {
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
