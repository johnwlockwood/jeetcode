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
		arr[i] = 1
	}
	return &UnionFind{arr: arr, n: n, size: size}
}

func (u *UnionFind) root(i int) int {
	for u.arr[i] != i {
		i = u.arr[i]
	}
	return i
}

func (u *UnionFind) Find(a, b int) bool {
	if u.root(a) == u.root(b) {
		return true
	}
	return false
}

func (u *UnionFind) Union(a, b int) {
	rootA := u.root(a)
	rootB := u.root(b)
	u.arr[rootA] = rootB
}

func (u *UnionFind) WeightedUnion(a, b int) {
	rootA := u.root(a)
	rootB := u.root(b)
	if u.size[rootA] < u.size[rootB] {
		u.arr[rootA] = u.arr[rootB]
		u.size[rootB] += u.size[rootA]
	} else {
		u.arr[rootB] = u.arr[rootA]
		u.size[rootA] += u.size[rootB]
	}
}

func find(arr []int, a, b int) bool {
	if root(arr, a) == root(arr, b) {
		return true
	}
	return false
}

func root(arr []int, i int) int {
	for arr[i] != i {
		i = arr[i]
	}
	return i
}

func union(arr []int, a, b int) {
	rootA := root(arr, a)
	rootB := root(arr, b)
	arr[rootA] = rootB
}

func initialize(arr []int, n int) {
	for i := 0; i < n; i++ {
		arr[i] = i
	}
}
