package longest

import "fmt"

func findCenter(s string, first int) int {
	for last := first; last < len(s); last++ {
		if s[first] != s[last] {
			return last - 1
		}
	}
	return first
}

func constructPalindrome(s string, first int) (int, int) {
	fmt.Printf("%s at %d\n", s, first)
	last := findCenter(s, first)
	prefix := ""
	for i := 0; i < first; i++ {
		prefix += " "
	}
	fmt.Printf("%s%s   %d,%d\n", prefix, string(s[first:last+1]), first, last)
	if first == 0 || last == len(s)-1 {
		return first, last
	}
	j := 1
	prev := first - j
	next := last + j
	for prev >= 0 && next < len(s) {
		if s[prev] != s[next] {
			prefix = ""
			for i := 0; i < prev+1; i++ {
				prefix += " "
			}
			fmt.Printf("%s%s\n", prefix, string(s[prev+1:next-1+1]))
			return prev + 1, next - 1
		} else if prev == 0 || next == len(s)-1 {
			return prev, next
		}
		fmt.Printf("\tj++\n")
		j++
		prev = first - j
		next = last + j
	}
	fmt.Println("\tfinal")
	return first - j, last + j
}

func longestPalindrome(s string) string {
	// locate and analyze?
	// maybe find
	return ""
}
