package longest

import "fmt"

func findCenter(s string, first int) int {
	last := first
	for last < len(s) {
		if s[first] != s[last] {
			// the previous last is the end
			return last - 1
		} else if last == len(s)-1 {
			return last
		}
		last++
	}
	return last
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
	// break problem into parts
	// track the longest so far
	// find the center then expand
	longest := ""
	first := 0
	last := 0
	for first < len(s) {
		first, last = constructPalindrome(s, first)
		if last+1-first > len(longest) {
			longest = string(s[first : last+1])
		}
		first = last + 1
	}
	return longest
}
