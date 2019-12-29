package generate

func sumZero(n int) []int {
	vals := make([]int, 0, n)
	mid := n / 2
	r := n % 2
	for i := 0 - mid; i <= mid; i++ {
		if i == 0 && r == 0 {
			continue
		}
		vals = append(vals, i)
	}
	return vals
}
