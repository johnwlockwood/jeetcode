package fibonacci

var mem map[int]int

func init() {
	mem = make(map[int]int, 0)
}

func fibRecursiveMem(n int) int {
	// An example of Dynamic Programming
	// Using top down recursion with memoization
	// Results from previous calls are kept.
	if v, ok := mem[n]; ok {
		return v
	}
	if n == 0 {
		mem[n] = 0
	} else if n <= 2 {
		mem[n] = 1
	} else {
		mem[n] = fibRecursiveMem(n-1) + fibRecursiveMem(n-2)
	}
	return mem[n]
}

func fibBottomUp(n int) int {
	// An example of Dynamic Programming
	// using a bottom up approach with memoization
	mem := make([]int, n+1)

	for i := 0; i <= n; i++ {
		if i == 0 {
			mem[i] = 0
		} else if i <= 2 {
			mem[i] = 1
		} else {
			mem[i] = mem[i-1] + mem[i-2]
		}
	}
	return mem[n]
}
