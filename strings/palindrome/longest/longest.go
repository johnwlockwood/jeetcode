package longest

import "fmt"

func findCenter(s string, first int) int {
	// find the center going forward only
	last := first + 1
	for last < len(s) && s[first] == s[last] {
		last++
	}
	return last - 1
}

func expandFromCenter(s string, first, last int) (int, int) {
	prev := first - 1
	next := last + 1
	for prev >= 0 && next < len(s) && s[prev] == s[next] {
		prev--
		next++
	}
	return prev + 1, next - 1
}

func constructPalindrome(s string, first int) (int, int) {
	// find the center then expand
	fmt.Printf("%s at %d\n", s, first)
	last := findCenter(s, first)
	first, last = expandFromCenter(s, first, last)
	return first, last
}

func longestPalindrome(s string) string {
	// break problem into parts
	// track the longest so far
	// for the first first, construct the palindrome by
	// 	finding the center then expanding to find the first and last
	// make the next first after the last
	longest := ""
	first := 0
	last := 0
	for first < len(s) {
		centerLast := findCenter(s, first)
		first, last = expandFromCenter(s, first, centerLast)
		if last+1-first > len(longest) {
			longest = s[first : last+1]
		}
		first = centerLast + 1
	}
	return longest
}
