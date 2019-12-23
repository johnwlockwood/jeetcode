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

// without print statements:
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Palindromic Substring.
// Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Longest Palindromic Substring.
// https://leetcode.com/submissions/detail/287942819/
func longestPalindrome(s string) string {
	// break problem into parts
	// track the longest so far
	// for the first first, construct the palindrome by
	// 	finding the center then expanding to find the first and last
	// make the next first after the last
	if len(s) < 2 {
		return s
	}
	longestLeft := 0
	longestRight := 0
	first := 0
	fmt.Printf("%s\n", s)
	for first < len(s) {
		centerRight := findCenter(s, first)
		left, right := expandFromCenter(s, first, centerRight)
		if right-left > longestRight-longestLeft {
			longestLeft = left
			longestRight = right
		}
		fmt.Printf("\tfirst %s \tf: %d +=%d\t%s\n", s[first:first+1], first, centerRight+1-first, s[longestLeft:longestRight+1])
		first = centerRight + 1
	}
	return s[longestLeft : longestRight+1]
}
