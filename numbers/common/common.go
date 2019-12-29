package common

// solution for /problems/maximum-length-of-repeated-subarray/
// Maximum Length of Repeated Subarray

func findLength(A []int, B []int) int {
	// Bottom up approach, Dynamic Programming
	// memoize, memoization
	// check every A against every B
	// if a pair match, set the value at that match to 1 + the value at the previous pair
	m := len(A)
	n := len(B)
	lcs := make([][]int, m+1)
	for i := range lcs {
		lcs[i] = make([]int, n+1)
	}
	max := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if A[i-1] == B[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
				if lcs[i][j] > max {
					max = lcs[i][j]
				}
			}
		}
	}
	return max
}
