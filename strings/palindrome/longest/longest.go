package longest

import "fmt"

func findCenter(s string, first int) int {
	// find the center going forward only
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

func expandFromCenter(s string, first, last int) (int, int) {
	prefix := ""
	for i := 0; i < first; i++ {
		prefix += " "
	}
	fmt.Printf("%s%s   %d,%d\n", prefix, s[first:last+1], first, last)
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
			fmt.Printf("%s%s\n", prefix, s[prev+1:next-1+1])
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
		fmt.Printf("first %s\n", s[first:first+1])
		centerLast := findCenter(s, first)
		first, last = expandFromCenter(s, first, centerLast)
		if last+1-first > len(longest) {
			longest = s[first : last+1]
		}
		first = centerLast + 1
	}
	return longest
}
