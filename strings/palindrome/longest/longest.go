package longest

func findCenter(s string, first int) int {
	for last := first; last < len(s); last++ {
		if s[first] != s[last] {
			return last - 1
		}
	}
	return first
}

func constructPalindrome(s string, first int) (int, int) {
	last := findCenter(s, first)
	for j := 1; first-j > 0 && last+j < len(s); j++ {
		prev := first - j
		next := last + j
		if s[prev] != s[next] {
			return prev + 1, next - 1
		}
	}
	return first, last
}

func longestPalindrome(s string) string {
	// locate and analyze?
	// maybe find
	return ""
}
