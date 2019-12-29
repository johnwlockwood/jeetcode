package fibonacci

var mem map[int]int

func init() {
	mem = make(map[int]int, 0)
}

func fibRecursiveMem(n int) int {
	// An example of Dynamic Programming
	// Using top down recursion with memoization
	if v, ok := mem[n]; ok {
		return v
	}
	if n <= 2 {
		mem[n] = 1
	} else {
		mem[n] = fibRecursiveMem(n-1) + fibRecursiveMem(n-2)
	}
	return mem[n]
}
