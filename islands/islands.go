package islands

import "github.com/johnwlockwood/jeetcode/unionfind"

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	u := unionfind.NewUnionFind(m * n)
	// connect land
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				if i > 0 && grid[i-1][j] == '1' {
					_ = u.Union(gridToArray(n, i, j), gridToArray(n, i-1, j))
				}
				if j > 0 && grid[i][j-1] == '1' {
					_ = u.Union(gridToArray(n, i, j), gridToArray(n, i, j-1))
				}
			}
		}
	}

	// count common roots of land
	roots := make(map[int]struct{}, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				root, _ := u.Root(gridToArray(n, i, j))
				roots[root] = struct{}{}
			}
		}
	}

	// stub
	return len(roots)
}

func gridToArray(n, i, j int) int {
	return i*n + j
}

func arrayToGrid(x, n int) (int, int) {
	i := x / n
	j := x % n
	return i, j
}
