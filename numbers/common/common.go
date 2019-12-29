package common

// solution for /problems/maximum-length-of-repeated-subarray/
// Maximum Length of Repeated Subarray

func findLength(A []int, B []int) int {
	// Bottom up approach, Dynamic Programming
	// tabulation method
	// for every A that matches a B, we need to know the count at the previous pair.
	// if A[3] == B[2], we are going to set lcs[3][2] = 1 + lcs[2][1]
	// if A[2] == B[1] then that would be 2
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
