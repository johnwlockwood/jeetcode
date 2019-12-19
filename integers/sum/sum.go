package sum

// getSum uses Half adder logic
func getSum(a int, b int) int {
	for b != 0 {
		carry := a & b
		a = a ^ b
		b = carry << 1
	}
	return a
}
